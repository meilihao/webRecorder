package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/meilihao/water"
	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar = logger.Sugar()
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
	id := hex.EncodeToString(tmpHash[:]) + ".wav"

	idPath := buildPathFn("sounds", id)
	err = ioutil.WriteFile(idPath, buf.Bytes(), 0666)
	if err != nil {
		ctx.Abort(500)

		return
	}

	ctx.DataJson(idPath)
}

var getPathFn = func(tag, id string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", tag, id[:2], id[2:4], id[4:6], id)
}

var buildPathFn = func(tag, id string) string {
	rootPath := fmt.Sprintf("%s/%s/%s/%s", tag, id[:2], id[2:4], id[4:6])
	os.MkdirAll(rootPath, 0777)

	return fmt.Sprintf("%s/%s", rootPath, id)
}
