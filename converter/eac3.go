package converter

import (
	"fmt"
	"os"
	"os/exec"
)

func ToEAC3(ffmpegFile os.File, inputFileIn, outputFileIn string, channelsIn int, kbpsIn int) error {
	values := []string{}
	kbps := fmt.Sprintf("%dk", kbpsIn)
	startValues := []string{
		"-i", inputFileIn, 
		"-loglevel", "warning",
		"-c:a", "eac3",
	}
	endValues := []string{
		"-b:a", kbps, 
		"-ac", fmt.Sprintf("%d", channelsIn), 
		outputFileIn,
	}

	values = append(values, startValues[:]...)
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