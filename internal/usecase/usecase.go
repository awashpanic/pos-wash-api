package usecase

import (
	"github.com/ffajarpratama/pos-wash-api/config"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/ffajarpratama/pos-wash-api/pkg/aws"
	"github.com/ffajarpratama/pos-wash-api/pkg/google"
	"github.com/ffajarpratama/pos-wash-api/pkg/redis"
	"gorm.io/gorm"
)

type Usecase struct {
	Cnf   *config.Config
	Repo  repository.IFaceRepository
	DB    *gorm.DB
	Redis redis.IFaceRedis
	AWS   aws.IFaceAWS
	FCM   google.IFaceFCM
}

func New(params *Usecase) IFaceUsecase {
	return &Usecase{
		Cnf:   params.Cnf,
		Repo:  params.Repo,
		DB:    params.DB,
		Redis: params.Redis,
		AWS:   params.AWS,
		FCM:   params.FCM,
	}
}
