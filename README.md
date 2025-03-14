# Edconv

<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/> <img src="https://shields.io/badge/FFmpeg-%23171717.svg?logo=ffmpeg&style=for-the-badge&labelColor=171717&logoColor=5cb85c"/>

[![Version](https://img.shields.io/badge/Version-1.2.5-blue)]()
[![Version](https://img.shields.io/badge/GoLang-v1.24.0-blue)]()
[![Version](https://img.shields.io/badge/FFmpeg-v7.1.1-blue)]()

## Supported formats

Audio: `AAC (FDK), E-AC3, AV1 (SVT) and H.265 (x265)`

## Usage

```
  -bit string
        Pixel format (8 for 8bit and 10 for 10bit) (default "8")
  -channels string
        Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound and 62 for downmixing 5.1 to stereo (default "2")
  -crf string
        Constant Rate Factor (0-63 for av1 and 0-51 for h265)
  -format string
        File format: aac, eac3, av1 and h265
  -input string
        Input file
  -kbps string
        Bitrate in kbps (192 for 192 kbps)
  -noAudio
        Video without audio
  -output string
        Output file
  -preset string
        Preset (0-13 for av1 and ultrafast, superfast, veryfast, faster, fast, medium, slow, slower and veryslow for h265)
  -sampleRate string
        Sample rate (44100 for 44100Hz)
  -vbr string
        Variable Bit Rate (1-5 for aac, 1 is lowest quality and 5 is highest quality)
  -version
        Show the version of the application
  -width string
        Width (1920 for 1080p, 1280 for 720p and 3840 for 2160p) (default "1920")
```