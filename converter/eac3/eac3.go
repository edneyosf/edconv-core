package eac3

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"edconv/converter"
)

const codec = "eac3"

func Convert(ffmpegFile os.File, inputFile, outputFile, channelsIn, kbpsIn, sampleRate string) error {
	values := []string{}
	kbps := fmt.Sprintf("%sk", kbpsIn)
	channels,af := channelsHandler(channelsIn)
	startValues := []string{
		"-i", inputFile, 
		"-loglevel", converter.LogLevel,
		"-c:a", codec,
	}
	endValues := []string{
		"-b:a", kbps, 
		"-ac", channels, 
		outputFile,
	}

	values = append(values, startValues[:]...)
	if len(af) > 0 {
		values = append(values, af[:]...)
	}
	if sampleRate != "" {
		values = append(values, []string{"-ar", sampleRate}...)
	}
	values = append(values, endValues[:]...)

	cmd := exec.Command(ffmpegFile.Name(), values[:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil { 
		return fmt.Errorf("error: %v", err) 
	}

	return nil
}

func channelsHandler(channelsIn string) (string, []string) {
	var channel string
	var af []string

	switch channelsIn {
	case "6":
		channel = "6"
    default:
        log.Fatal("Unsupported number of channels")
    }

	return channel,af
}