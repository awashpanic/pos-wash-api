package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/ffajarpratama/pos-wash-api/internal/http/request"
	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/custom_error"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/goodsign/monday"
	"github.com/google/uuid"
)

// GetDashboardSummary implements IFaceUsecase.
func (u *Usecase) GetDashboardSummary(ctx context.Context, outletID uuid.UUID) (*response.DashoardSummary, error) {
	params := &request.OrderTrendQuery{
		OutletID: outletID,
	}

	order, err := u.Repo.GetOrderSummary(ctx, params)
	if err != nil {
		return nil, err
	}

	customer, err := u.Repo.GetCustomerSummary(ctx, outletID)
	if err != nil {
		return nil, err
	}

	odelta := order.Rev2 - order.Rev1
	if odelta != 0 {
		odelta = (odelta / order.Rev1) * 100
	}

	cdelta := customer.Count2 - customer.Count1
	if cdelta != 0 {
		cdelta = (cdelta / customer.Count1) * 100
	}

	res := &response.DashoardSummary{
		OrderCount: &response.OrderCount{
			Accepted:  order.Accepted,
			OnProcess: order.OnProcess,
			Complete:  order.Complete,
		},
		RevenueTrend: &response.RevenueTrend{
			Initial: order.Rev1,
			Final:   order.Rev2,
			Delta:   odelta,
		},
		CustomerTrend: &response.CustomerTrend{
			Initial: customer.Count1,
			Final:   customer.Count2,
			Delta:   cdelta,
		},
	}

	return res, nil
}

// GetOrderTrend implements IFaceUsecase.
func (u *Usecase) GetOrderTrend(ctx context.Context, params *request.OrderTrendQuery) (*response.OrderTrend, error) {
	if params.Type == "" {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "`Type` cannot be empty",
		})

		return nil, err
	}

	timeframe := map[string]bool{
		"weekly":  true,
		"monthly": true,
	}

	if !timeframe[params.Type] {
		err := custom_error.SetCustomError(&custom_error.ErrorContext{
			HTTPCode: http.StatusBadRequest,
			Message:  "invalid time frame",
		})

		return nil, err
	}

	tn := time.Now()

	if params.Start == "" {
		switch params.Type {
		case "weekly":
			params.Start = util.WeekStart(tn.ISOWeek()).Format("2006-01-02")
		case "monthly":
			params.Start = time.Date(tn.Year(), time.January, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		}
	}

	if params.End == "" {
		switch params.Type {
		case "weekly":
			params.End = util.WeekStart(tn.ISOWeek()).Add(24 * 6 * time.Hour).Format("2006-01-02")
		case "monthly":
			params.End = time.Date(tn.Year(), time.December, 31, 0, 0, 0, 0, time.UTC).Format("2006-01-02")
		}
	}

	trend, err := u.Repo.GetOrderTrend(ctx, params)
	if err != nil {
		return nil, err
	}

	for _, v := range trend {
		if params.Type == "weekly" {
			v.Label = monday.Format(v.Date, "Monday", "id_ID")
		} else {
			v.Label = monday.Format(v.Date, "January", "id_ID")
		}
	}

	order, err := u.Repo.GetOrderSummary(ctx, params)
	if err != nil {
		return nil, err
	}

	odelta := order.Rev2 - order.Rev1
	if odelta != 0 {
		odelta = (odelta / order.Rev1) * 100
	}

	res := &response.OrderTrend{
		OrderCount: &response.OrderCount{
			Accepted:  order.Accepted,
			OnProcess: order.OnProcess,
			Complete:  order.Complete,
		},
		RevenueTrend: &response.RevenueTrend{
			Initial: order.Rev1,
			Final:   order.Rev2,
			Delta:   odelta,
		},
		Trend: trend,
	}

	return res, nil
}
