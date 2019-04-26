# 网页录音工具
参考:
- [HTML5网页录音和上传到服务器，支持PC、Android，支持IOS微信](https://www.tuicool.com/articles/BvuA3aN)和[xiangyuecn/Recorder](https://github.com/xiangyuecn/Recorder)

## 工具
- [trim_audio_silence.go](trim_audio_silence.go): 使用silencedetect识别静默,再截去收尾(因为直接使用silenceremove会截取话语中的静默).