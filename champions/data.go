package champions

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// Champion struct type
type Champion struct {
	Name    string   `json:"name"`
	Classes []string `json:"classes"`
	Origins []string `json:"origins"`
	Cost    int      `json:"cost"`
}

// Filter function
func Filter(list []string, f func(string) bool) []string {
	xs := []string{}

	for _, v := range list {
		if f(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

// HasAllClasses func
func (c Champion) HasAllClasses(classes ...string) bool {
	classCount := len(classes)
	foundClass := 0

	for _, class := range classes {
		if found := c.HasClass(class); found {
			foundClass++
		}
	}
	return foundClass == classCount
}

// HasSomeClasses func
func (c Champion) HasSomeClasses(classes ...string) bool {
	for _, class := range classes {
		if found := c.HasClass(class); found {
			return true
		}
	}
	return false
}

// HasClass method if there is any class
func (c Champion) HasClass(class string) bool {
	for _, champClass := range c.Classes {
		if champClass == class {
			return true
		}
	}
	return false
}

// HasOrigin check origin
func (c Champion) HasOrigin(origin string) bool {
	for _, ch := range c.Origins {
		if ch == origin {
			return true
		}
	}
	return false
}

// HasSomeOrigins check some
func (c Champion) HasSomeOrigins(origins ...string) bool {
	for _, origin := range origins {
		if found := c.HasOrigin(origin); found {
			return true
		}
	}
	return false
}

func (c Champion) HasAllOrigins(origins ...string) bool {
	originCount := len(origins)
	foundCount := 0
	for _, origin := range origins {
		if found := c.HasOrigin(origin); found {
			foundCount++
		}
	}

	return foundCount == originCount
}

// LoadChampions load data
func LoadChampions(fileName string) ([]Champion, error) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("could not read json data")
	}
	defer file.Close()

	var champions []Champion

	for {
		if err := json.NewDecoder(file).Decode(&champions); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
	}

	return champions, err
}


// CreateChampion func
func CreateChampion(c Champion) Champion{
	newChamp := Champion{
		Name: c.Name,
		Classes: c.Classes,
		Origins: c.Origins,
		Cost: c.Cost,		
	}
	return newChamp
}