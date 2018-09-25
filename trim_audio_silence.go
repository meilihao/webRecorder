// https://www.bookstack.cn/read/other-doc-cn-ffmpeg/ffmpeg-doc-cn-34.md#silencedetect
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	ffmpegPath = "/usr/local/bin/ffmpeg"
)

var (
	// 提取silencedetect的输出信息
	silenceReg = regexp.MustCompile(`(?m)^\[silencedetect\s@\s.*\]\s([^\n|\\|]+)`)
)

func main() {
	names, err := filepath.Glob("*")
	CheckErr(err)

	var tmp int
	for _, v := range names {
		if strings.Contains(v, ".wav") { // example 1.wav
			tmp, err = strconv.Atoi(strings.TrimSuffix(v, ".wav"))
			CheckErr(err)

			ps := Silencedetect(v)
			fmt.Println(ps)
			fmt.Println(ps[1], ps[len(ps)-2])
			ConvertWAV(v, fmt.Sprintf("%03d.1.wav", tmp), fmt.Sprintf("%f", ps[1]), fmt.Sprintf("%f", ps[len(ps)-2]))
		}
	}
}

// 按照指定时间截取音频
func ConvertWAV(originPath, wavPah, start, end string) error {
	cmd := &exec.Cmd{
		Path: ffmpegPath,
	}

	cmd.Args = []string{ffmpegPath, "-hide_banner", "-i", originPath, "-y",
		"-ss", start, "-to", end,
		wavPah}

	_, err := ExecCmdWithError(cmd)
	if err != nil {
		panic(err)
	}

	return nil
}

// 使用silencedetect参数识别空白
func Silencedetect(originPath string) []float32 {
	cmd := &exec.Cmd{
		Path: ffmpegPath,
	}

	// silencedetect参数不好调试, 看自己情况设置
	cmd.Args = []string{ffmpegPath, "-hide_banner", "-i", originPath,
		"-af", "silencedetect=noise=-48dB:d=0.04", "-f", "null", "-"}

	output, err := ExecCmdWithError(cmd)
	if err != nil {
		panic(err)
	}

	var points []float32
	var index int
	var tmpPoint float64

	ss := silenceReg.FindAllString(string(output), -1)
	for _, v := range ss {
		index = strings.Index(v, ":")
		if index == -1 {
			panic(err)
		}

		v = strings.TrimSpace(v[index+1:])
		tmpPoint, err = strconv.ParseFloat(v, 64)
		CheckErr(err)

		points = append(points, float32(tmpPoint))
	}

	if len(points)%2 != 0 {
		panic("invalid points")
	}

	return points
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

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
