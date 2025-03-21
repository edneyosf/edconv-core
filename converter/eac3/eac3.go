package eac3

import (
	"fmt"
	"log"

	"edconv/cmd"
	"edconv/converter"
)

func Convert(ffmpeg, inputFile, outputFile, channelsIn, kbpsIn, sampleRate string) error {
	values := []string{}
	kbps := fmt.Sprintf("%sk", kbpsIn)
	channels := channelsHandler(channelsIn)
	startValues := []string{
		"-i", inputFile, 
		"-loglevel", converter.LogLevel,
		"-c:a", codec,
	}
	endValues := []string{
		"-b:a", kbps,
		outputFile,
	}

	values = append(values, startValues[:]...)

	if sampleRate != "" {
		values = append(values, []string{"-ar", sampleRate}...)
	}
	if channels != "" {
		values = append(values, []string{"-ac", channels}...)
	}

	values = append(values, endValues[:]...)

	return cmd.Run(ffmpeg, values)
}

func channelsHandler(channelsIn string) (string) {
	var channel string

	switch channelsIn {
	case "6":
		channel = "6"
	case "":
		channel = ""	
    default:
        log.Fatal("Error: Unsupported number of channels")
    }

	return channel
}