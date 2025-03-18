package info

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

func FromMedia(ffprobe, input string) {
	cmd := exec.Command(ffprobe, "-v", logLevel, "-show_entries", "format=" + formats, "-of", "default=noprint_wrappers=1", input)
	infoRegex := regexp.MustCompile(infoRegex)

    output, err := cmd.CombinedOutput()
    if err != nil {
		log.Fatal("Error: ", err)
    }

    matches := infoRegex.FindStringSubmatch(string(output))
    if len(matches) == matchLen {
		duration, err := strconv.ParseFloat(matches[1], 64)
		if err != nil {
			log.Fatal("Error: ", err)
		}

		size, err := strconv.ParseInt(matches[2], 10, 64)
		if err != nil {
			log.Fatal("Error: ", err)
		}
	
		mediaInfo := MediaInfo{
			Duration: formatDuration(duration),
			Size: size,
		}
	
		mediaInfoJSON, err := json.Marshal(mediaInfo)
		mediaInfoString := string(mediaInfoJSON)
		
		if err == nil {
			fmt.Println("mediaInfo=" + mediaInfoString)
		} else {
			log.Fatal("Error: ", err)
		}
    } else {
		log.Fatal("Error: ", err)
	}
}

func formatDuration(seconds float64) string {
	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	seconds = seconds - float64(hours * 3600) - float64(minutes * 60)

    return fmt.Sprintf("%02d:%02d:%05.2f", hours, minutes, seconds)
}