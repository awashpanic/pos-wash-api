package usecase

import (
	"context"
	"net/mail"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/hash"
	"github.com/ffajarpratama/pos-wash-api/pkg/jwt"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.ReqRegister) (*model.User, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	user := &model.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber.Format(),
		Password:    req.Password.Hash(),
	}

	err := u.Repo.CreateUser(ctx, user, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomUserClaims{
		ID:   user.UserID.String(),
		Role: string(constant.Owner),
	}

	tokens, err := jwt.GenerateToken(claims, u.Cnf)
	if err != nil {
		return nil, err
	}

	user.AccessToken = tokens.AccessToken
	user.RefreshToken = tokens.RefreshToken

	return user, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.ReqLogin) (*model.User, error) {
	_, err := mail.ParseAddress(req.Identifier)
	if err != nil {
		req.Identifier = string(types.PhoneNumber(req.Identifier).Format())
	}

	res, err := u.Repo.FindOneUser(ctx, "email = ? OR phone_number = ?", req.Identifier, req.Identifier)
	if err != nil {
		return nil, err
	}

	err = hash.Compare(string(res.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	outletOwner, err := u.Repo.FindOneOutletOwner(ctx, "user_id = ?", res.UserID)
	if err != nil && !repository.IsRecordNotfound(err) {
		return nil, err
	}

	if outletOwner != nil {
		res.Outlet = outletOwner.Outlet
	}

	claims := &jwt.CustomUserClaims{
		ID:   res.UserID.String(),
		Role: string(constant.Owner),
	}

	tokens, err := jwt.GenerateToken(claims, u.Cnf)
	if err != nil {
		return nil, err
	}

	res.AccessToken = tokens.AccessToken
	res.RefreshToken = tokens.RefreshToken

	return res, nil
}
