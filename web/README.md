# README
web端录音并上传.

## 特点
1. 支持截取音频开头的空白

## index2.html
1. 在浏览器显示声音波形
1. 在波形支持下的录音截取

> webm转成AudioBuffer后无法转回去, 且WebAudio API没有encode AudioBuffer的方法.

## 浏览器支持
chrome

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
1. [网页音频接口的基本概念](https://developer.mozilla.org/zh-CN/docs/Web/API/Web_Audio_API/Basic_concepts_behind_Web_Audio_API)
1. [recorder.js](https://wangpengfei15974.github.io/recorder.js/)
1. [html5录音](https://www.jianshu.com/p/1b90743386b2)
1. [HTML5录音控件](https://www.cnblogs.com/xiaoqi/p/6993912.html)
1. [peaks.js:用于与音频波形交互的JavaScript UI组件, 音频可视化](https://github.com/bbc/peaks.js)
