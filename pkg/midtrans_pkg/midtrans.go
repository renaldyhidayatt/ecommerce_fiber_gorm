package midtranspkg

import (
	midtransrequest "ecommerce_fiber/internal/domain/requests/midtrans_request"

	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/spf13/viper"
)

type SnapClient struct {
	client snap.Client
}

func setupGlobalMidtransConfig() {
	midtrans.ServerKey = viper.GetString("MIDTRANS_SERVER_KEY")
	midtrans.ClientKey = viper.GetString("MIDTRANS_CLIENT_KEY")
	midtrans.Environment = midtrans.Sandbox

}

func NewSnapClient() *SnapClient {
	setupGlobalMidtransConfig()

	var s snap.Client

	s.New(viper.GetString("MIDTRANS_SERVER"), midtrans.Sandbox)

	return &SnapClient{s}
}

func (s *SnapClient) CreateTransaction(request midtransrequest.CreateMidtransRequest) (*snap.Response, error) {
	params := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "order-csb-" + uuid.New().String(),
			GrossAmt: int64(request.GrossAmount),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: false,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.FirstName,
			LName: request.LastName,
			Email: request.Email,
			Phone: request.Phone,
		},
		Callbacks: &snap.Callbacks{
			Finish: "http://localhost:3000/user/oder",
		},
	}

	res, err := snap.CreateTransaction(params)

	if err != nil {
		return nil, err
	}

	return res, nil

}
