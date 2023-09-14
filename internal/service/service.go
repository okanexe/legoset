package service

import (
	"errors"

	"legoapi/internal/app"
	"legoapi/internal/repository"
)

type Service interface {
	GetAllLegoSets() ([]app.LegoSet, error)
	GetLegoSetByCode(code string) (app.LegoSet, error)
	CreateLegoSet(legoSet app.LegoSet) error
	UpdateLegoSet(code string, legoSet app.LegoSet) error
	DeleteLegoSet(code string) error
}

type LegoSetService struct {
	legoSetRepository repository.Repository
}

func NewLegoSetService(repository *repository.LegoSetRepository) *LegoSetService {
	return &LegoSetService{
		legoSetRepository: repository,
	}
}

func (s *LegoSetService) GetAllLegoSets() ([]app.LegoSet, error) {
	legoSets, err := s.legoSetRepository.GetAllLegoSets()
	if err != nil {
		return nil, err
	}
	return legoSets, nil
}

func (s *LegoSetService) GetLegoSetByCode(code string) (app.LegoSet, error) {
	legoSet, err := s.legoSetRepository.GetLegoSetByCode(code)
	if err != nil {
		return app.LegoSet{}, err
	}
	return legoSet, nil
}

func (s *LegoSetService) CreateLegoSet(legoSet app.LegoSet) error {
	// İş kurallarını burada doğrulayabilirsiniz, örneğin fiyat pozitif olmalıdır.
	if legoSet.Price <= 0 {
		return errors.New("fiyat pozitif olmalıdır")
	}

	err := s.legoSetRepository.CreateLegoSet(legoSet)
	if err != nil {
		return err
	}
	return nil
}

func (s *LegoSetService) UpdateLegoSet(code string, legoSet app.LegoSet) error {
	// İş kurallarını burada doğrulayabilirsiniz.

	err := s.legoSetRepository.UpdateLegoSet(code, legoSet)
	if err != nil {
		return err
	}
	return nil
}

func (s *LegoSetService) DeleteLegoSet(code string) error {
	err := s.legoSetRepository.DeleteLegoSet(code)
	if err != nil {
		return err
	}
	return nil
}
