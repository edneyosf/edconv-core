package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"edconv/converter/aac"
	"edconv/converter/av1"
	"edconv/converter/eac3"
	"edconv/converter/h265"
)

//go:embed bin/ffmpeg711
var ffmpeg []byte

func main() {
	showVersion := flag.Bool("version", false, "Show the version of the application")
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	format := flag.String("format", "", "File format: aac, eac3, av1 and h265")
	channels := flag.Int("channels", channelsDefault, "Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound and 62 for downmixing 5.1 to stereo")
	kbps := flag.Int("kbps", -1, "Bitrate in kbps (192 for 192 kbps)")
	sampleRate := flag.String("sampleRate", "", "Sample rate (44100 for 44100Hz)")
	preset := flag.String("preset", "", "Preset (0-13 for av1 and ultrafast, superfast, veryfast, faster, fast, medium, slow, slower and veryslow for h265)")
	crf := flag.Int("crf", -1, "Constant Rate Factor (0-63 for av1 and 0-51 for h265)")
	bit := flag.Int("bit", bitDefault, "Pixel format (8 for 8bit and 10 for 10bit)")
	width := flag.Int("width", widthDefault, "Width (1920 for 1080p, 1280 for 720p and 3840 for 2160p)")
	noAudio := flag.Bool("noAudio", false, "Video without audio")
	flag.Parse()

	if(*kbps == -1) {
		*kbps = kbpsDefault(channels)
	}
	if(*preset == ""){
		*preset = presetDefault(format)
	}
	if(*crf == -1) {
		*crf = crfDefault(format)
	}

	checkFlags(showVersion, format, inputFile, outputFile)
	formatHandler(format, retrieveFFmpeg(), inputFile, outputFile, channels, kbps, sampleRate, preset, crf, bit, width, noAudio)

	fmt.Println("Success!")
}

func kbpsDefault(channels* int) int {
	var kbps int

	switch *channels {
	case 2:
		kbps = kbps20Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               
	case 6:			
		kbps = kbps51Default
    default:
		kbps = kbps20Default
    }

	return kbps
}

func presetDefault(format* string) string {
	var preset string

	switch *format {
	case "av1":
		preset = presetAv1Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               
	case "h265":			
		preset = presetH265Default
    default:
		preset = ""
    }

	return preset
}

func crfDefault(format* string) int {
	var crf int

	switch *format {
	case "av1":
		crf = crfAv1Default                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           
	case "h265":			
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
		log.Fatal("You must specify the format, input file, and output file.")
	}
}

func retrieveFFmpeg() *os.File {
	tmpFile, err := os.CreateTemp("", "edconv-ffmpeg-")
	if err != nil {
		log.Fatalf("Error creating temporary file: %v\n", err)
	}

	_, err = tmpFile.Write(ffmpeg)
	if err != nil {
		log.Fatalf("Error writing binary to temporary file: %v\n", err)
	}

	err = os.Chmod(tmpFile.Name(), 0755)
	if err != nil {
		log.Fatalf("Error setting execution permissions: %v\n", err)
	}

	tmpFile.Close()

	return tmpFile
}

func formatHandler(
	format* string, ffmpegFile* os.File, inputFile* string, outputFile* string, channels* int, kbps* int, sampleRate* string,
	preset* string, crf* int, bit* int, width* int, noAudio* bool) {
	var err error

	switch *format {
    case "aac":
		err = aac.Convert(*ffmpegFile, *inputFile, *outputFile, *channels, *kbps, *sampleRate)
	case "eac3":
		err = eac3.Convert(*ffmpegFile, *inputFile, *outputFile, *channels, *kbps, *sampleRate)	
	case "av1":
		err = av1.Convert(*ffmpegFile, *inputFile, *outputFile, *preset, *crf, *bit, *width, *noAudio)
	case "h265":	
		err = h265.Convert(*ffmpegFile, *inputFile, *outputFile, *preset, *crf, *bit, *width, *noAudio)
    default:
        log.Fatal("Unsupported format")
    }

	defer os.Remove(ffmpegFile.Name())

	if err != nil {
		log.Fatal(err)
	}
}