package service

import (
	den_task "Den_task1"
	"Den_task1/package/repository"
)

type DenDataService struct {
	repo repository.DenchickData
}

func NewDenDataService(repo repository.DenchickData) *DenDataService {
	return &DenDataService{repo: repo}
}

func (s *DenDataService) GetDenData() (*den_task.DenData, error) {
	dto, err := s.repo.GetMe()

	if err != nil {
		return nil, err
	}
	return dto, nil
}
