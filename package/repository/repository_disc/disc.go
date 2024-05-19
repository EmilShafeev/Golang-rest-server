package repository_disc

import (
	"errors"
	"io"
	"os"
)

type Disc struct {
	SendAllwaysError bool
}

type Option func(*Disc)

func NewDisc(opts ...Option) *Disc {

	d := &Disc{
		SendAllwaysError: false,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithAllwaysError(flag bool) Option {
	return func(d *Disc) {
		d.SendAllwaysError = flag
	}
}

func (d *Disc) ReadJSONFromFile(path string) ([]byte, error) {
	if d.SendAllwaysError {
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
