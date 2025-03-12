package main

import (
	"fmt"
	"os"
	"os/exec"
)

// AAC com FDK-AAC
func convertToAAC(ffmpegFile os.File, inputFileIn, outputFileIn string, channelsIn int, kbpsIn int) error {
	kbps := fmt.Sprintf("%dk", kbpsIn)
	channels := fmt.Sprintf("%d", channelsIn)
	outputFile := fmt.Sprintf("%s.aac", outputFileIn)

	cmd := exec.Command(
		ffmpegFile.Name(), 
		"-i", inputFileIn, 
		"-loglevel", "warning",
		"-c:a", "libfdk_aac", 
		"-b:a", kbps, 
		"-ac", channels, 
		outputFile)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil { 
		return fmt.Errorf("error: %v", err) 
	}

	return nil
}