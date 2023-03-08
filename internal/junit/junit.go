package junit

import (
	"encoding/xml"
	"os"
)

func Read(path string) (*Testsuites, error) {
	results := &Testsuites{}

	dat, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := xml.Unmarshal(dat, results); err != nil {
		return nil, err
	}

	return results, nil
}
