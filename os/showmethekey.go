package os

import (
	"bufio"
	"encoding/json"
	"os/exec"
)

func GetKeyEvents() <-chan KeyEvent {
	showMeTheKeyEvents := getShowMeTheKeyEvents()
	channel := make(chan KeyEvent)
	go func() {
		for {
			event := <-showMeTheKeyEvents
			state := RELEASED
			if event.StateName == "PRESSED" {
				state = PRESSED
			}
			channel <- KeyEvent{KeyCode: event.KeyCode, State: state}
		}
	}()
	return channel
}

func getShowMeTheKeyEvents() <-chan ShowMeTheKeyEvent {
	channel := make(chan ShowMeTheKeyEvent)
	command := exec.Command("showmethekey-cli")
	stdout, err := command.StdoutPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			text := scanner.Text()

			event := ShowMeTheKeyEvent{}
			err := json.Unmarshal([]byte(text), &event)
			if err != nil {
				panic(err)
			}
			channel <- event
		}
	}()

	go func() {
		if err := command.Run(); err != nil {
			panic(err)
		}
	}()
	return channel
}

type KeyEventState int

const (
	PRESSED KeyEventState = iota
	RELEASED
)

type KeyEvent struct {
	KeyCode int
	State   KeyEventState
}

type ShowMeTheKeyEvent struct {
	EventName string `json:"event_name"`
	EventType int    `json:"event_type"`
	TimeStamp uint   `json:"time_stamp"`
	KeyName   string `json:"key_name"`
	KeyCode   int    `json:"key_code"`
	StateName string `json:"state_name"`
	StateCode int    `json:"state_code"`
}
