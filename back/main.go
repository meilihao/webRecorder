package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/labstack/echo"
	"github.com/meilihao/water"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger
var ffmpegLP string

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar = logger.Sugar()

	var err error
	ffmpegLP, err = exec.LookPath("ffmpeg")
	if err != nil {
		panic(err)
	}
}

func main() {
	router := water.NewRouter()

	router.Use(water.CORS(water.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE},
	}))
	router.Post("/sounds", _UploadSound)

	w := router.Handler()

	if err := water.ListenAndServe(":8081", w); err != nil {
		log.Fatalln(err)
	}
}

func _UploadSound(ctx *water.Context) {
	file, _, err := ctx.Request.FormFile("audio")
	if err != nil {
		sugar.Error(err)
		ctx.Abort(400)

		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		sugar.Error(err)
		ctx.Abort(400)

		return
	}

	tmpHash := md5.Sum(buf.Bytes())
	id := hex.EncodeToString(tmpHash[:])

	webmPath := buildPathFn("sounds", id+".webm")
	err = ioutil.WriteFile(webmPath, buf.Bytes(), 0666)
	if err != nil {
		sugar.Error(err)
		ctx.Abort(500)

		return
	}

	wavPah := buildPathFn("sounds", id+".wav")
	err = ConvertWAV(webmPath, wavPah)
	if err != nil {
		sugar.Error(err)
		ctx.Abort(500)

		return
	}

	ctx.DataJson(id + ".wav")
}

func ConvertWAV(webmPath, wavPah string) error {
	cmd := &exec.Cmd{
		Path: ffmpegLP,
	}

	cmd.Args = []string{ffmpegLP, "-i", webmPath, "-y",
		"-acodec", "pcm_s16le", "-ac", "1", "-ar", "8000", wavPah}

	output, err := ExecCmdWithError(cmd)
	if err != nil {
		sugar.Error(string(output))
		return err
	}

	return err
}

func ExecCmdWithError(c *exec.Cmd) ([]byte, error) {
	if c.Stderr != nil {
		return nil, errors.New("exec: Stderr already set")
	}
	var b bytes.Buffer
	c.Stderr = &b
	err := c.Run()
	return b.Bytes(), err
}

var getPathFn = func(tag, id string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", tag, id[:2], id[2:4], id[4:6], id)
}

var buildPathFn = func(tag, id string) string {
	rootPath := fmt.Sprintf("%s/%s/%s/%s", tag, id[:2], id[2:4], id[4:6])
	os.MkdirAll(rootPath, 0777)

	return fmt.Sprintf("%s/%s", rootPath, id)
}
