package av1

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"edconv/converter"
)

const codec = "libsvtav1"

// AV1 com SVT
func Convert(ffmpeg, inputFile, outputFile, preset, crf, bitIn, width string, noAudio bool) error {
	values := []string{}
	bit := bitHandler(bitIn)
	startValues := []string{
		"-i", inputFile, 
		"-loglevel", converter.LogLevel,
		"-c:v", codec,
		"-preset", preset,
		"-crf", crf,
		"-profile:v", "0",
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

	cmd := exec.Command(ffmpeg, values[:]...)
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