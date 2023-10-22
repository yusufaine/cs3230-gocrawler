package filewriter

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func ToJSON(v any, filename string) error {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	// create folder if not exists
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}

	var f *os.File
	// if file exists, overwrite
	if _, err := os.Stat(filename); err == nil {
		f, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
	}

	f, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(d); err != nil {
		return err
	}
	if _, err := f.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}