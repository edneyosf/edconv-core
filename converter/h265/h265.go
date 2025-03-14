package h265

import (
	"edconv/converter"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const codec = "libx265"

// H.265 com x265
func Convert(ffmpegFile os.File, inputFile, outputFile, preset, crf, bitIn, width string, noAudio bool) error {
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
	if(noAudio) {
		values = append(values, "-an")
	} else {
		values = append(values, []string{"-c:a", "copy"}...)
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

func bitHandler(bitIn string) string {
	var bit string

	switch bitIn {
	case "8":
		bit = "yuv420p"
	case "10":
		bit = "yuv420p10le"
    default:
        log.Fatal("Unsupported pixel format")
    }

	return bit
}