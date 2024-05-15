package service

import (
	den_task "Den_task1"
	"Den_task1/package/repository"
)

type DenData interface {
	GetDenData() (*den_task.DenData, error)
}

type Service struct {
	DenData
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		DenData: NewDenDataService(repos.DenchickData),
	}
}
