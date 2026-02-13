package os

import (
	"cmp"
	"os/exec"
	"regexp"
	"slices"
	"strconv"
)

func GetActiveWindowTitleChannel() <-chan string {
	channel := make(chan string)
	lastTitle := ""
	go func() {
		for {
			title := getActiveWindowTitle()
			if title == lastTitle {
				continue
			}
			lastTitle = title
			channel <- title
		}
	}()
	return channel
}

func getActiveWindowTitle() string {
	command := "hyprctl clients"
	output, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		panic(err)
	}
	outputText := string(output)
	expression := regexp.MustCompile(`Window (.*?) -> (.*?):[\S\s.]*?focusHistoryID: ([0-9]*)`)

	matches := expression.FindAllStringSubmatch(outputText, -1)

	slices.SortFunc(matches, func(a, b []string) int {
		aIndex, err := strconv.Atoi(a[3])
		if err != nil {
			panic(err)
		}
		bIndex, err := strconv.Atoi(b[3])
		if err != nil {
			panic(err)
		}
		return cmp.Compare(aIndex, bIndex)
	})
	return matches[0][2]
}
