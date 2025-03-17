package h265

import (
	"fmt"
	"log"

	"edconv/cmd"
	"edconv/converter"
)

func Convert(ffmpeg, inputFile, outputFile, preset, crf, bitIn, width string, noAudio bool) error {
	values := []string{}
	bit := bitHandler(bitIn)
	startValues := []string{
		"-i", inputFile, 
		"-loglevel", converter.LogLevel,
		"-c:v", codec,
		"-preset", preset,
		"-crf", crf,
		"-pix_fmt", bit,
		"-vf", fmt.Sprintf("scale=%s:-1", width),
		"-b:v", "0",
	}
	endValues := []string{outputFile}

	values = append(values, startValues[:]...)

	if noAudio {
		values = append(values, "-an")
	} else {
		values = append(values, []string{"-c:a", "copy"}...)
	}

	values = append(values, endValues[:]...)

	return cmd.Run(ffmpeg, values)
}

func bitHandler(bitIn string) string {
	var bit string

	switch bitIn {
	case "8":
		bit = bit8
	case "10":
		bit = bit10
    default:
        log.Fatal("Error: Unsupported pixel format")
    }

	return bit
}