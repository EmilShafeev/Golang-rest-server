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

func NewDiscRepository(disc *repository_disc.Disc) *Repository {
	return &Repository{
		DenchickData: repository_disc.NewDenchickDataDisk(disc),
	}
}
