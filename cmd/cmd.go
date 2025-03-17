package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"regexp"
	"strings"
)

func Run(name string, args []string) error {
	cmd := exec.Command(name, args[:]...)

	stdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	parseProgress(stdErr)

	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

func parseProgress(pipe io.ReadCloser) {
	scanner := bufio.NewScanner(pipe)
	scanner.Split(rawScan)
	
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "time=") {
			progressToJson(line)
		} else {
			fmt.Print(line)
		}
	}
}

func progressToJson(line string) {
	progressRegex := regexp.MustCompile(progressRegex)
	match := progressRegex.FindStringSubmatch(line)

	if len(match) == matchLen {
		progress := Progress{
			Size:    match[1],
			Time:    match[2],
			Bitrate: match[3],
			Speed:   match[4],
		}

		progressJSON, err := json.Marshal(progress)
		progressString := string(progressJSON)
		
		if err == nil {
			fmt.Println("progress=" + progressString)
		} else {
			fmt.Println("Error generating progress JSON: ", err)
		}
	}
}

func rawScan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	return len(data), dropCR(data), nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}

	return data
}