package repository_disc

import (
	"errors"
	"io"
	"os"
)

type Disc struct {
	settings DiscSettings
}

type DiscSettings struct {
	SendAllwaysError bool
}

func NewDisc(settings DiscSettings) *Disc {
	return &Disc{settings: settings}
}

func (d *Disc) ReadJSONFromFile(path string) ([]byte, error) {
	if d.settings.SendAllwaysError {
		return nil, errors.New("forced error")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
