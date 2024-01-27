package usecase_test

import (
	"context"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ffajarpratama/pos-wash-api/config"
	mock_repo "github.com/ffajarpratama/pos-wash-api/internal/mock"
	"github.com/ffajarpratama/pos-wash-api/internal/usecase"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UsecaseTestSuite struct {
	suite.Suite
	uc       usecase.IFaceUsecase
	mockRepo *mock_repo.MockIFaceRepository
	ctx      context.Context
	goMock   *gomock.Controller
	dbMock   sqlmock.Sqlmock
}

func (s *UsecaseTestSuite) SetupTest() {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		log.Printf("[error-sql-mock]: %v\n", err.Error())
	}

	dbClient, err := gorm.Open(postgres.New(postgres.Config{Conn: dbMock}))
	if err != nil {
		log.Printf("[error-postgres-connection]: %v\n", err)
	}

	s.ctx = context.TODO()
	s.goMock = gomock.NewController(s.T())
	s.mockRepo = mock_repo.NewMockIFaceRepository(s.goMock)

	require.NoError(s.T(), err)

	s.dbMock = mock
	s.uc = usecase.New(&usecase.Usecase{
		Cnf: &config.Config{
			JWTConfig: config.JWTConfig{
				Admin:   "1234",
				User:    "4321",
				Refresh: "4422",
			},
		},
		Repo: s.mockRepo,
		DB:   dbClient,
	})
}

func (s *UsecaseTestSuite) AfterTest(suiteName, testName string) {
	defer s.goMock.Finish()
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UsecaseTestSuite))
}
