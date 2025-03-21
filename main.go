package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"edconv/converter/aac"
	"edconv/converter/av1"
	"edconv/converter/eac3"
	"edconv/converter/h265"
	"edconv/info"
)

func main() {
	var ffmpeg string
	var ffprobe string

	showVersion := flag.Bool("version", false, "Show the version of the application")
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	format := flag.String("format", "", "File format: aac, eac3, av1 and h265")
	channels := flag.String("channels", "", "Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound and 62 for downmixing 5.1 to stereo")
	kbps := flag.String("kbps", "", "Bitrate in kbps (192 for 192 kbps)")
	vbr := flag.String("vbr", "", "Variable Bit Rate (1-5 for aac, 1 is lowest quality and 5 is highest quality)")
	sampleRate := flag.String("sampleRate", "", "Sample rate (44100 for 44100Hz)")
	preset := flag.String("preset", "", "Preset (0-13 for av1 and ultrafast, superfast, veryfast, faster, fast, medium, slow, slower and veryslow for h265)")
	crf := flag.String("crf", "", "Constant Rate Factor (0-63 for av1 and 0-51 for h265)")
	bit := flag.String("bit", "", "Pixel format (8 for 8bit and 10 for 10bit)")
	width := flag.String("width", "", "Width (1920 for 1080p, 1280 for 720p and 3840 for 2160p)")
	noAudio := flag.Bool("noAudio", false, "Video without audio")
	ffmpegPath := flag.String("ffmpeg", "", "FFmpeg path")
	ffprobePath := flag.String("ffprobe", "", "FFprobe path")
	flag.Parse()

	if *kbps == "" {
		*kbps = kbpsDefault(channels)
	}
	if *preset == "" {
		*preset = presetDefault(format)
	}
	if *crf == "" {
		*crf = crfDefault(format)
	}

	if *ffmpegPath != "" {
		ffmpeg = *ffmpegPath
	} else {
		ffmpeg = ffmpegDefault
	}
	if *ffprobePath != "" {
		ffprobe = *ffprobePath
	} else {
		ffprobe = ffprobeDefault
	}

	checkFlags(showVersion, format, inputFile, outputFile)
	info.FromMedia(ffprobe, *inputFile)
	formatHandler(ffmpeg, *format, *inputFile, *outputFile, *channels, *vbr, *kbps, *sampleRate, *preset, *crf, *bit, *width, *noAudio)
}

func kbpsDefault(channels* string) string {
	var kbps string

	switch *channels {
	case "2":
		kbps = kbps20Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               
	case "6":			
		kbps = kbps51Default
    default:
		kbps = kbps20Default
    }

	return kbps
}

func presetDefault(format* string) string {
	var preset string

	switch *format {
	case av1Format:
		preset = presetAv1Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               
	case h265Format:			
		preset = presetH265Default
    default:
		preset = ""
    }

	return preset
}

func crfDefault(format* string) string {
	var crf string

	switch *format {
	case av1Format:
		crf = crfAv1Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           
	case h265Format:			
		crf = crfH265Default
    default:
		crf = crfAv1Default
    }

	return crf
}

func showVersions(showVersion* bool) {
	if *showVersion {
		fmt.Println(appName + " v" + version) 
		fmt.Println("FFmpeg v" + ffmpegVersion)
		os.Exit(0)
	}
}

func checkFlags(showVersion* bool, format* string, inputFile* string, outputFile* string) {

	showVersions(showVersion)

	if *format == "" || *inputFile == "" || *outputFile == "" {
		log.Fatal("Error: You must specify the format, input file, and output file.")
	}
}

func formatHandler(ffmpeg, format, inputFile, outputFile, channels, vbr, kbps, sampleRate, preset, crf, bit, width string, noAudio bool) {
	var err error

	switch format {
    case aacFormat:
		err = aac.Convert(ffmpeg, inputFile, outputFile, channels, vbr, kbps, sampleRate)
	case eac3Format:
		err = eac3.Convert(ffmpeg, inputFile, outputFile, channels, kbps, sampleRate)	
	case av1Format:
		err = av1.Convert(ffmpeg, inputFile, outputFile, preset, crf, bit, width, noAudio)
	case h265Format:	
		err = h265.Convert(ffmpeg, inputFile, outputFile, preset, crf, bit, width, noAudio)
    default:
        log.Fatal("Error: Unsupported format")
    }

	if err != nil {
		log.Fatal("Error: ", err)
	}
}