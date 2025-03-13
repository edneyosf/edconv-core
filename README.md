# Edconv

<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://shields.io/badge/FFmpeg-%23171717.svg?logo=ffmpeg&style=for-the-badge&labelColor=171717&logoColor=5cb85c"/>

[![Version](https://img.shields.io/badge/Version-1.2.2-blue)]()
[![Version](https://img.shields.io/badge/GoLang-v1.24.0-blue)]()
[![Version](https://img.shields.io/badge/FFmpeg-v7.1.1-blue)]()

## Supported formats

Audio: `AAC (libfdk_aac) and E-AC3`

## Usage

```
  -channels int
        Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound, 62 for downmixing 5.1 to stereo (default 2)
  -format string
        File format: AAC and E-AC3
  -input string
        Input file
  -kbps int
        Bitrate in kbps (192 for 192 kbps) (default 192)
  -output string
        Output file
  -sampleRate string
        Sample rate (44100 for 44100Hz)
  -version
        Show the version of the application
```