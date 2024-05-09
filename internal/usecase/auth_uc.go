package usecase

import (
	"context"
	"net/http"
	"net/mail"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/internal/repository"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"github.com/ffajarpratama/pos-wash-api/pkg/hash"
	"github.com/ffajarpratama/pos-wash-api/pkg/jwt"
	"github.com/ffajarpratama/pos-wash-api/pkg/types"
	"github.com/google/uuid"
)

// Register implements IFaceUsecase.
func (u *Usecase) Register(ctx context.Context, req *request.Register) (*model.User, error) {
	tx := u.DB.Begin()
	defer tx.Rollback()

	user := &model.User{
		Name:        req.User.Name,
		Email:       req.User.Email,
		PhoneNumber: req.User.PhoneNumber.Format(),
		Password:    req.User.Password.Hash(),
	}

	err := u.Repo.CreateUser(ctx, user, tx)
	if err != nil {
		if repository.IsDuplicateErr(err) {
			err = custom_error.SetCustomError(&custom_error.ErrorContext{
				HTTPCode: http.StatusConflict,
				Code:     constant.DefaultDuplicateDataError,
				Message:  "email or phone already used",
			})
		}

		return nil, err
	}

	outlet := &model.Outlet{
		UserID:  user.UserID,
		Name:    req.Outlet.Name,
		Address: req.Outlet.Address,
		LogoID:  req.Outlet.LogoID,
	}

	err = u.Repo.CreateOutlet(ctx, outlet, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomClaims{
		ID:       user.UserID.String(),
		OutletID: outlet.OutletID.String(),
		Role:     string(constant.Owner),
	}

	user.AccessToken, err = jwt.GenerateToken(claims, u.Cnf.JWTConfig.Secret)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Login implements IFaceUsecase.
func (u *Usecase) Login(ctx context.Context, req *request.Login) (*model.User, error) {
	_, err := mail.ParseAddress(req.Identifier)
	if err != nil {
		req.Identifier = string(types.PhoneNumber(req.Identifier).Format())
	}

	res, err := u.Repo.FindOneUser(ctx, "email = ? OR phone_number = ?", req.Identifier, req.Identifier)
	if err != nil {
		return nil, err
	}

	err = hash.Compare([]byte(res.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	outlet, err := u.Repo.FindOneOutlet(ctx, "user_id = ?", res.UserID)
	if err != nil {
		return nil, err
	}

	claims := &jwt.CustomClaims{
		ID:   res.UserID.String(),
		Role: string(constant.Owner),
	}

	if outlet != nil {
		res.Outlet = outlet
		claims.OutletID = outlet.OutletID.String()
	}

	res.AccessToken, err = jwt.GenerateToken(claims, u.Cnf.JWTConfig.Secret)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// FindOneUser implements IFaceUsecase.
func (u *Usecase) FindOneUser(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	return u.Repo.FindOneUser(ctx, "user_id = ?", userID)
}
