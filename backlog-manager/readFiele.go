package backlogmanager

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Todo type
type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	About     string `json:"about"`
	Completed bool   `json:"completed"`
}

// ReadFile func
func ReadFile(fileName string) ([]Todo, error) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Could not parse current json file", err)
	}

	defer jsonFile.Close()

	var todos []Todo

	for {
		if err := json.NewDecoder(jsonFile).Decode(&todos); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}
	return todos, err
}
