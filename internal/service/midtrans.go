package service

import (
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"
	midtranspkg "ecommerce_fiber/pkg/midtrans_pkg"

	"github.com/midtrans/midtrans-go/snap"
)

type midtransService struct {
	snapClient *midtranspkg.SnapClient
}

func NewMidtransService(snapClient *midtranspkg.SnapClient) *midtransService {
	return &midtransService{
		snapClient: snapClient,
	}
}

func (s *midtransService) CreateTransaction(request *midtransrequest.CreateMidtransRequest) (*snap.Response, error) {

	res, err := s.snapClient.CreateTransaction(*request)

	if err != nil {
		return nil, err
	}

	return res, nil
}
