package aac

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"edconv/converter"
)

const codec = "libfdk_aac"

// AAC com FDK-AAC
func Convert(ffmpegFile os.File, inputFileIn, outputFileIn string, channelsIn int, kbpsIn int, sampleRate string) error {
	values := []string{}
	kbps := fmt.Sprintf("%dk", kbpsIn)
	channels,af := channelsHandler(channelsIn)
	startValues := []string{
		"-i", inputFileIn, 
		"-loglevel", converter.LogLevel,
		"-c:a", codec,
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

func channelsHandler(channelsIn int) (string, []string) {
	var channel string
	var af []string

	switch channelsIn {
    case 2:
		channel = "2"
	case 62:
		channel = "2"
		af = append(af, "-af")
		//af = append(af, "pan=stereo| FL=0.5*FL + 0.5*FC + 0.5*SL + 0.5*LFE | FR=0.5*FR + 0.5*FC + 0.5*SR + 0.5*LFE")
		af = append(af, "pan=stereo|FL=0.8*FL+0.5*FC+0.3*SL+0.3*LFE|FR=0.8*FR+0.5*FC+0.3*SR+0.3*LFE")
	case 6:
		channel = "6"
    default:
        log.Fatal("Unsupported number of channels")
    }

	return channel,af
}