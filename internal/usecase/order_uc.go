package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/model"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/google/uuid"
)

// CreateOrder implements IFaceUsecase.
func (u *Usecase) CreateOrder(ctx context.Context, req *request.CreateOrder) (*model.Order, error) {
	outlet, err := u.Repo.FindOneOutlet(ctx, "outlet_id = ?", req.OutletID)
	if err != nil {
		return nil, err
	}

	pm, err := u.Repo.FindOnePaymentMethod(ctx, "payment_method_id = ?", req.PaymentMethodID)
	if err != nil {
		return nil, err
	}

	perfume, err := u.Repo.FindOnePerfume(ctx, "perfume_id = ?", req.PerfumeID)
	if err != nil {
		return nil, err
	}

	smap := make(map[uuid.UUID]request.Service)
	serviceIDs := make([]uuid.UUID, 0, len(req.Services))
	for _, v := range req.Services {
		smap[v.ServiceID] = v
		serviceIDs = append(serviceIDs, v.ServiceID)
	}

	services, err := u.Repo.FindService(ctx, "service_id IN (?)", serviceIDs)
	if err != nil {
		return nil, err
	}

	var subtotal float64
	var total float64
	var fee float64
	var discount float64
	var est time.Duration

	orderID := uuid.New()
	odList := make([]*model.OrderDetail, 0, len(req.Services))

	for _, v := range services {
		tmp, ok := smap[v.ServiceID]
		if !ok {
			continue
		}

		var unit time.Duration
		switch v.EstCompletionUnit {
		case constant.EstCompletionMinute:
			unit = time.Minute
		case constant.EstCompletionHour:
			unit = time.Hour
		case constant.EstCompletionDay:
			unit = 24 * time.Hour
		case constant.EstCompletionWeek:
			unit = 24 * 7 * time.Hour
		case constant.EstCompletionMonth:
			t1 := time.Now()
			t2 := time.Date(t1.Year(), t1.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
			unit = 24 * time.Duration(t2) * time.Hour
		case constant.EstCompletionYear:
			unit = 24 * 365 * time.Hour
		}

		d := unit * time.Duration(v.EstCompletion)
		if est < d {
			est = d
		}

		odList = append(odList, &model.OrderDetail{
			OrderID:     orderID,
			OutletID:    req.OutletID,
			ServiceID:   v.ServiceID,
			Quantity:    tmp.Quantity,
			TotalAmount: v.Price * float64(tmp.Quantity),
		})

		subtotal += v.Price * float64(tmp.Quantity)
	}

	total += subtotal

	if pm.FeePercentage > 0 {
		fee = total * (pm.FeePercentage / 100)
	} else if pm.Fee > 0 {
		fee = pm.Fee
	}

	tx := u.DB.Begin()
	defer tx.Rollback()

	order := &model.Order{
		OrderID:          orderID,
		OutletID:         req.OutletID,
		PaymentMethodID:  req.PaymentMethodID,
		PerfumeID:        req.PerfumeID,
		CustomerID:       req.CustomerID,
		UserID:           req.UserID,
		InvoiceNumber:    fmt.Sprintf("%s-INV-%s", outlet.Code, util.GenerateRandomString(9, true)),
		SubtotalAmount:   subtotal,
		SubtotalFee:      fee,
		SubtotalDiscount: discount,
		TotalAmount:      (subtotal + fee + perfume.Price) - discount,
		Status:           constant.OrderAccepted,
		Note:             req.Note,
		EstCompletionAt:  time.Now().Add(est),
	}

	err = u.Repo.CreateOrder(ctx, order, tx)
	if err != nil {
		return nil, err
	}

	err = u.Repo.CreateManyOrderDetail(ctx, odList, tx)
	if err != nil {
		return nil, err
	}

	history := []*model.OrderHistoryStatus{{
		OrderID:   orderID,
		Status:    constant.OrderAccepted,
		CreatedBy: req.UserID,
	}}

	err = u.Repo.CreateManyOrderHistoryStatus(ctx, history, tx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return u.Repo.FindOneOrder(ctx, "order_id = ?", order.OrderID)
}

// FindAndCountOrder implements IFaceUsecase.
func (u *Usecase) FindAndCountOrder(ctx context.Context, params *request.ListOrderQuery) ([]*model.Order, int64, error) {
	return u.Repo.FindAndCountOrder(ctx, params)
}

// FindOneOrder implements IFaceUsecase.
func (u *Usecase) FindOneOrder(ctx context.Context, orderID uuid.UUID) (*model.Order, error) {
	return u.Repo.FindOneOrder(ctx, "order_id = ?", orderID)
}

// UpdateOrderStatus implements IFaceUsecase.
func (u *Usecase) UpdateOrderStatus(ctx context.Context, req *request.UpdateOrderStatus) error {
	order, err := u.Repo.FindOneOrder(ctx, "order_id = ?", req.OrderID)
	if err != nil {
		return err
	}

	if order.Status == constant.OrderComplete {
		err = custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "order has been completed",
		})

		return err
	}

	tx := u.DB.Begin()
	defer tx.Rollback()

	data := map[string]interface{}{
		"status": req.Status,
	}

	err = u.Repo.UpdateOrder(ctx, tx, data, "order_id = ?", req.OrderID)
	if err != nil {
		return err
	}

	history := []*model.OrderHistoryStatus{{
		OrderID:   req.OrderID,
		Status:    req.Status,
		CreatedBy: req.UserID,
	}}

	err = u.Repo.CreateManyOrderHistoryStatus(ctx, history, tx)
	if err != nil {
		return err
	}

	return tx.Commit().Error
}
