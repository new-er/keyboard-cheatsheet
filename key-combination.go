package main

import (
	"encoding/json"
	"os"
)

type KeyCombination struct {
	Keys        []KeyCode
	Description string
	Application string
  Disabled  bool
}

func KeyCombinationsToJson(k []KeyCombination) (string, error) {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func KeyCombinationsFromJson(jsonData string) ([]KeyCombination, error) {
	var k []KeyCombination
	err := json.Unmarshal([]byte(jsonData), &k)
	if err != nil {
		return nil, err
	}
	return k, nil
}

func KeyCombinationsToFile(k []KeyCombination, filename string) error {
	jsonData, err := KeyCombinationsToJson(k)
	if err != nil {
		return err
	}
	return WriteToFile(jsonData, filename)
}

func KeyCombinationsFromFileOrPanic(filename string) []KeyCombination {
	k, err := KeyCombinationsFromFile(filename)
	if err != nil {
		panic(err)
	}
	return k
}
func KeyCombinationsFromFile(filename string) ([]KeyCombination, error) {
	jsonData, err := ReadFromFile(filename)
	if err != nil {
		return nil, err
	}
	return KeyCombinationsFromJson(jsonData)
}

func WriteToFile(data string, filename string) error {
	return os.WriteFile(filename, []byte(data), 0644)
}

func ReadFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


func FilterDisabledKeyCombinations(k []KeyCombination) []KeyCombination {
  return Filter(k, func(keyCombination KeyCombination) bool {
    return !keyCombination.Disabled
  })
}
