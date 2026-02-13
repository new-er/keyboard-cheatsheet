package data

import (
	"encoding/json"
	"io"
	"new-er/keyboard-cheatsheet/helpers"
	"os"
)

type Key struct {
	Names []string
	KeyCode int
}

func NewUnknownKey(keyCode int) Key {
	return Key{
		Names: []string{"Unknown"},
		KeyCode: keyCode,
	}
}

func ReadKeyMapFromFile() map[int]Key {
	keys := ReadKeyMapArrayFromFile()
	keyMap := make(map[int]Key)
	for _, key := range keys {
		keyMap[key.KeyCode] = key
	}
	return keyMap
}
func ReadKeyMapArrayFromFile() []Key {
	keyMapFile := "keymap.json"
	file, err := os.Open(keyMapFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}


	keys := []Key{}
	err = json.Unmarshal(byteValue, &keys)
	if err != nil {
		panic(err)
	}
	
	return keys
}

func (k Key) Equals(other helpers.Equatable) bool {
	otherKey, ok := other.(Key)
	if !ok {
		return false
	}
	return k.KeyCode == otherKey.KeyCode
}	
