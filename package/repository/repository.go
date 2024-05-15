package repository

import (
	den_task "Den_task1"
	"Den_task1/package/repository/repository_disc"
)

type DenchickData interface {
	GetMe() (*den_task.DenData, error)
}

type Repository struct {
	DenchickData
}

func NewDiscRepository(settings repository_disc.DiscSettings) *Repository {
	disc := repository_disc.NewDisc(settings)
	return &Repository{
		DenchickData: repository_disc.NewDenchickDataDisk(disc),
	}
}
