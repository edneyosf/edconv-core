# Edconv

<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://shields.io/badge/FFmpeg-%23171717.svg?logo=ffmpeg&style=for-the-badge&labelColor=171717&logoColor=5cb85c"/>

[![Version](https://img.shields.io/badge/Version-1.2.3-blue)]()
[![Version](https://img.shields.io/badge/GoLang-v1.24.0-blue)]()
[![Version](https://img.shields.io/badge/FFmpeg-v7.1.1-blue)]()

## Supported formats

Audio: `AAC (libfdk_aac), E-AC3 and AV1`

## Usage

```
  -bit int
        Pixel format (8 for 8bit and 10 for 10bit) (default 8)
  -channels int
        Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound and 62 for downmixing 5.1 to stereo (default 2)
  -crf int
        Constant Rate Factor (default 25)
  -format string
        File format: AAC, E-AC3 and AV1
  -input string
        Input file
  -kbps int
        Bitrate in kbps (192 for 192 kbps) (default 192)
  -noAudio
        Video without audio
  -output string
        Output file
  -preset int
        Preset (default 4)
  -sampleRate string
        Sample rate (44100 for 44100Hz)
  -version
        Show the version of the application
  -width int
        Width (1920 for 1080p, 1280 for 720p and 3840 for 2160p) (default 1920)
```