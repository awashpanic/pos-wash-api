package usecase

import (
	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"gorm.io/gorm"
)

type Usecase struct {
	Cnf  *config.Config
	Repo repository.IFaceRepository
	DB   *gorm.DB
}

func New(params *Usecase) IFaceUsecase {
	return &Usecase{
		Cnf:  params.Cnf,
		Repo: params.Repo,
		DB:   params.DB,
	}
}
