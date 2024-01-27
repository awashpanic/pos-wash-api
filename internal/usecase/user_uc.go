package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/google/uuid"
)

// FindOneUser implements IFaceUsecase.
func (u *Usecase) FindOneUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.Repo.FindOneUser(ctx, "user_id = ?", userID)
}
