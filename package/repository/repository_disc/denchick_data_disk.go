package repository_disc

import (
	den_task "Den_task1"
	"encoding/json"

	"github.com/spf13/viper"
)

type DenchickDataDisk struct {
	disc *Disc
}

func NewDenchickDataDisk(disc *Disc) *DenchickDataDisk {
	return &DenchickDataDisk{disc: disc}
}

func (ddd *DenchickDataDisk) GetMe() (dto *den_task.DenData, err error) {
	bytes, err := ddd.disc.ReadJSONFromFile(viper.GetString("json_path"))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &dto); err != nil {
		return nil, err
	}
	return dto, err
}
