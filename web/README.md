# README

## 备忘
1. [navigator.getUserMedia已更名为MediaDevices.getUserMedia](https://developer.mozilla.org/zh-CN/docs/Web/API/Navigator/getUserMedia)
1. [HTML音频和视频的媒体格式](https://developer.mozilla.org/en-US/docs/Web/HTML/Supported_media_formats)

## MediaStreamRecorder
1. 采用`mimeType = 'audio/wav'`,`sampleRate=8000`没声音
1. 采用`mimeType = 'audio/wav'`, 保存的声音被截断(即只能播放n毫秒内的声音`MediaStreamRecorder.start(n)`)
1. bitsPerSecond 仅支持Firefox
```js
// currently supported only in Firefox
videoRecorder.bitsPerSecond = 12800;
```

## 参考
1. [https://github.com/mdn/web-dictaphone](https://github.com/mdn/web-dictaphone)
1. [Web Audio API 初识](https://github.com/o2team/H5Skills/issues/64)