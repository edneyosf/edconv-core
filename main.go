package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"log"
)

const ffmpeg = "ffmpeg-ed"

func main() {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	channels := flag.Int("channels", 2, "Number of channels (1 for mono, 2 for stereo, 6 for 5.1, 8 for 7.1)")
	kbps := flag.Int("kbps", 192, "Bitrate in kbps (192 for 192 kbps)")
	flag.Parse()

	if *inputFile == "" || *outputFile == "" {
		log.Fatal("You must provide both the input and output files.")
	}

	err := convertToAAC(*inputFile, *outputFile, *channels, *kbps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success!")
}

// AAC com FDK-AAC
func convertToAAC(inputFile, outputFile string, channels int, kbps int) error {
	cmd := exec.Command(
		ffmpeg, 
		"-i", inputFile, 
		"-loglevel", "warning",
		"-c:a", "libfdk_aac", 
		"-b:a", fmt.Sprintf("%dk", kbps), 
		"-ac", fmt.Sprintf("%d", channels), 
		fmt.Sprintf("%s.aac", outputFile))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}