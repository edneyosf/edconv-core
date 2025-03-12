package converter

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// AAC com FDK-AAC

func channelsHandler(channelsIn int) (string, []string) {
	var channel string
	var af []string

	switch channelsIn {
    case 2:
		channel = "2"
	case 62:
		channel = "2"
		af = append(af, "-af")
		af = append(af, "pan=stereo| FL=0.5*FL + 0.5*FC + 0.5*SL + 0.5*LFE | FR=0.5*FR + 0.5*FC + 0.5*SR + 0.5*LFE")
	case 6:
		channel = "6"
	case 8:
		channel = "8"
    default:
        log.Fatal("Unsupported channel")
    }

	return channel,af
}

func ToAAC(ffmpegFile os.File, inputFileIn, outputFileIn string, channelsIn int, kbpsIn int) error {
	values := []string{}
	kbps := fmt.Sprintf("%dk", kbpsIn)
	channels,af := channelsHandler(channelsIn)
	startValues := []string{
		"-i", inputFileIn, 
		"-loglevel", "warning",
		"-c:a", "libfdk_aac",
	}
	endValues := []string{
		"-b:a", kbps, 
		"-ac", channels, 
		outputFileIn,
	}

	values = append(values, startValues[:]...)
	if len(af) > 0 {
		values = append(values, af[:]...)
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