package usecase

import (
	"context"

	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/google/uuid"
)

// GetDashboardSummary implements IFaceUsecase.
func (u *Usecase) GetDashboardSummary(ctx context.Context, outletID uuid.UUID) (*response.DashoardSummary, error) {
	order, err := u.Repo.GetOrderTrend(ctx, outletID)
	if err != nil {
		return nil, err
	}

	customer, err := u.Repo.GetCustomerTrend(ctx, outletID)
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
			Yesterday: order.Rev1,
			Today:     order.Rev2,
			Delta:     odelta,
		},
		CustomerTrend: &response.CustomerTrend{
			Yesterday: customer.Count1,
			Today:     customer.Count2,
			Delta:     cdelta,
		},
	}

	return res, nil
}
