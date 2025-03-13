package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"edconv/converter/aac"
	"edconv/converter/eac3"
)

const appName = "Edconv"
const version = "1.2.2"
const ffmpegVersion = "7.1.1"
const channelsDefault = 2
const kbpsDefault = 192

//go:embed bin/ffmpeg711
var ffmpeg []byte

func main() {
	showVersion := flag.Bool("version", false, "Show the version of the application")
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	format := flag.String("format", "", "File format: AAC and E-AC3")
	channels := flag.Int("channels", channelsDefault, "Number of channels: 2 for stereo, 6 for 5.1 surround sound, 8 for 7.1 surround sound, 62 for downmixing 5.1 to stereo")
	kbps := flag.Int("kbps", kbpsDefault, "Bitrate in kbps (192 for 192 kbps)")
	sampleRate := flag.String("sampleRate", "", "Sample rate (44100 for 44100Hz)")
	flag.Parse()

	checkFlags(showVersion, format, inputFile, outputFile)
	formatHandler(format, retrieveFFmpeg(), inputFile, outputFile, channels, kbps, sampleRate)

	fmt.Println("Success!")
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

func formatHandler(format* string, ffmpegFile* os.File, inputFile* string, outputFile* string, channels* int, kbps* int, sampleRate* string) {
	var err error

	switch *format {
    case "aac":
		err = aac.Convert(*ffmpegFile, *inputFile, *outputFile, *channels, *kbps, *sampleRate)
	case "eac3":
		err = eac3.Convert(*ffmpegFile, *inputFile, *outputFile, *channels, *kbps, *sampleRate)	
    default:
        log.Fatal("Unsupported format")
    }

	defer os.Remove(ffmpegFile.Name())

	if err != nil {
		log.Fatal(err)
	}
}