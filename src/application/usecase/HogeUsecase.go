package usecase

import (
	"docker-go-api/src/domain/model"
	"docker-go-api/src/domain/repository"
)

type IHogeUsecase interface {
	Execute(*model.Hoge) (error error)
}

type HogeUsecase struct {
	hogeRepository repository.IHogeRepository
}

func NewHogeUsecase(hogeRepository repository.IHogeRepository) IHogeUsecase {
	hogeUsecase := HogeUsecase{hogeRepository: hogeRepository}
	return &hogeUsecase
}

func (hogeUsecase *HogeUsecase) Execute(hoge *model.Hoge) (error error) {
	error = hogeUsecase.hogeRepository.Persist(hoge)
	return error
}
