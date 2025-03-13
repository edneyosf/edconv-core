# Edconv

<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://shields.io/badge/FFmpeg-%23171717.svg?logo=ffmpeg&style=for-the-badge&labelColor=171717&logoColor=5cb85c"/>

[![Version](https://img.shields.io/badge/Version-1.2.1-blue)]()
[![Version](https://img.shields.io/badge/GoLang-v1.24.0-blue)]()
[![Version](https://img.shields.io/badge/FFmpeg-v7.1.1-blue)]()

## Supported formats

Audio: `AAC (libfdk_aac)`

## Usage

```
  -channels int
    	Number of channels (1 for mono, 2 for stereo, 6 for 5.1, 8 for 7.1) (default 2)
  -format string
    	File format (AAC, EAC3)
  -input string
    	Input file
  -kbps int
    	Bitrate in kbps (192 for 192 kbps) (default 192)
  -output string
    	Output file (without extension)
  -version
    	Show the version of the application
```