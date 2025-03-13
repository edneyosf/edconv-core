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
func Convert(ffmpegFile os.File, inputFileIn, outputFileIn string, presetIn int, crfIn int, bitIn int, widthIn int, noAudio bool) error {
	values := []string{}
	bit := bitHandler(bitIn)
	startValues := []string{
		"-i", inputFileIn, 
		"-loglevel", converter.LogLevel,
		"-c:v", codec,
		"-preset", fmt.Sprintf("%d", presetIn),
		"-crf", fmt.Sprintf("%d", crfIn),
		"-profile:v", "0",
		"-pix_fmt", bit,
		"-vf", fmt.Sprintf("scale=%d:-1", widthIn),
		"-b:v", "0",
	}
	endValues := []string{outputFileIn}

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

func bitHandler(bitIn int) string {
	var bit string

	switch bitIn {
	case 8:
		bit = "yuv420p"
	case 10:
		bit = "yuv420p10le"
    default:
        log.Fatal("Unsupported pixel format")
    }

	return bit
}