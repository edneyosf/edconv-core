package aac

import (
	"fmt"
	"log"

	"edconv/cmd"
	"edconv/converter"
)

func Convert(ffmpeg, inputFile, outputFile, channelsIn, vbr, kbpsIn, sampleRate string) error {
	values := []string{}
	channels,af := channelsHandler(channelsIn)
	startValues := []string{
		"-i", inputFile, 
		"-loglevel", converter.LogLevel,
		"-c:a", codec,
	}
	endValues := []string{outputFile}

	values = append(values, startValues[:]...)

	if len(af) > 0 {
		values = append(values, af[:]...)
	}
	if sampleRate != "" {
		values = append(values, []string{"-ar", sampleRate}...)
	}
	if vbr != "" {
		values = append(values, []string{"-vbr", vbr}...)
	} else {
		kbps := fmt.Sprintf("%sk", kbpsIn)
		values = append(values, []string{"-b:a", kbps}...)
	}
	if channels != "" {
		values = append(values, []string{"-ac", channels}...)
	}

	values = append(values, endValues[:]...)

	return cmd.Run(ffmpeg, values)
}

func channelsHandler(channelsIn string) (string, []string){
	var channel string
	var af []string

	switch channelsIn {
    case "2":
		channel = "2"
	case "62":
		channel = "2"
		af = append(af, "-af")
		af = append(af, filterChannels62)
	case "6":
		channel = "6"
	case "":
		channel = ""	
    default:
        log.Fatal("Error: Unsupported number of channels")
    }

	return channel, af
}