package gapi

import (
	"context"
	rajaongkirrequest "ecommerce_fiber/internal/domain/requests/rajaongkir_request"
	"ecommerce_fiber/internal/pb"
	"ecommerce_fiber/internal/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type rajaOngkirHandleGrpc struct {
	pb.UnimplementedRajaOngkirServiceServer
	rajaongkir service.RajaOngkirService
}

func NewRajaOngirHandleGrpc(rajaongkir service.RajaOngkirService) *rajaOngkirHandleGrpc {
	return &rajaOngkirHandleGrpc{
		rajaongkir: rajaongkir,
	}
}

func (h *rajaOngkirHandleGrpc) GetProvinsi(ctx context.Context, empty *emptypb.Empty) (*pb.RajaOngkirResponseProvinsi, error) {
	res, err := h.rajaongkir.GetProvinsi()

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	rajaongkirResponse := &pb.RajaOngkirResponseProvinsi{
		Rajaongkir: &pb.RajaOngkirResponseProvinsi_Rajaongkir{
			Status: &pb.StatusProvinsi{
				Code:        int32(res.Rajaongkir.Status.Code),
				Description: res.Rajaongkir.Status.Description,
			},
		},
	}

	for _, query := range res.Rajaongkir.Query {
		rajaongkirResponse.Rajaongkir.Query = append(rajaongkirResponse.Rajaongkir.Query, &pb.QueryProvinsi{
			Id: query.ID,
		})
	}

	for _, result := range res.Rajaongkir.Results {
		rajaongkirResponse.Rajaongkir.Results = append(rajaongkirResponse.Rajaongkir.Results, &pb.ResultsProvinsi{
			ProvinceId: result.ProvinceID,
			Province:   result.Province,
		})
	}

	return rajaongkirResponse, nil
}

func (h *rajaOngkirHandleGrpc) GetCity(ctx context.Context, req *pb.CityRequest) (*pb.RajaOngkirCityResponse, error) {
	res, err := h.rajaongkir.GetCity(int(req.GetId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	rajaOngkirCityResponse := &pb.RajaOngkirCityResponse{
		Rajaongkir: &pb.RajaOngkirCityResponse_Rajaongkir{
			Status: &pb.StatusCity{
				Code:        int32(res.RajaOngkir.Status.Code),
				Description: res.RajaOngkir.Status.Description,
			},
			Query: &pb.QueryCity{
				Id:       res.RajaOngkir.Query.ID,
				Province: res.RajaOngkir.Query.Province,
			},
		},
	}

	for _, result := range res.RajaOngkir.Results {
		rajaOngkirCityResponse.Rajaongkir.Results = append(rajaOngkirCityResponse.Rajaongkir.Results, &pb.ResultsCity{
			CityId:     result.CityID,
			ProvinceId: result.ProvinceID,
			Province:   result.Province,
			Type:       result.Type,
			CityName:   result.CityName,
			PostalCode: result.PostalCode,
		})
	}

	return rajaOngkirCityResponse, nil
}

func (h *rajaOngkirHandleGrpc) GetCost(ctx context.Context, req *pb.OngkosRequest) (*pb.RajaOngkirOngkosResponse, error) {
	ongkorReq := rajaongkirrequest.OngkosRequest{
		Asal:   req.Asal,
		Tujuan: req.Tujuan,
		Berat:  int(req.Berat),
		Kurir:  req.Kurir,
	}

	res, err := h.rajaongkir.GetCost(ongkorReq)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %s", err.Error())
	}

	rajaOngkirOngkosResponse := &pb.RajaOngkirOngkosResponse{
		Rajaongkir: &pb.RajaOngkirOngkosResponse_RajaOngkir{
			Status: &pb.Status{
				Code:        int32(res.RajaOngkir.Status.Code),
				Description: res.RajaOngkir.Status.Description,
			},
			OriginDetails: &pb.CityDetails{
				CityId:     res.RajaOngkir.OriginDetails.CityID,
				ProvinceId: res.RajaOngkir.OriginDetails.ProvinceID,
				Province:   res.RajaOngkir.OriginDetails.Province,
				Type:       res.RajaOngkir.OriginDetails.Type,
				CityName:   res.RajaOngkir.OriginDetails.CityName,
				PostalCode: res.RajaOngkir.OriginDetails.PostalCode,
			},
			DestinationDetails: &pb.CityDetails{
				CityId:     res.RajaOngkir.DestinationDetails.CityID,
				ProvinceId: res.RajaOngkir.DestinationDetails.ProvinceID,
				Province:   res.RajaOngkir.DestinationDetails.Province,
				Type:       res.RajaOngkir.DestinationDetails.Type,
				CityName:   res.RajaOngkir.DestinationDetails.CityName,
				PostalCode: res.RajaOngkir.DestinationDetails.PostalCode,
			},
		},
	}

	for _, result := range res.RajaOngkir.Results {
		courierResult := &pb.CourierResult{
			Code: result.Code,
			Name: result.Name,
		}
		for _, cost := range result.Costs {
			pbCost := &pb.Cost{
				Service:     cost.Service,
				Description: cost.Description,
			}
			for _, costDetail := range cost.CostDetails {
				pbCostDetail := &pb.CostDetail{
					Value: int32(costDetail.Value),
					Etd:   costDetail.Etd,
					Note:  costDetail.Note,
				}
				pbCost.CostDetails = append(pbCost.CostDetails, pbCostDetail)
			}
			courierResult.Costs = append(courierResult.Costs, pbCost)
		}
		rajaOngkirOngkosResponse.Rajaongkir.Results = append(rajaOngkirOngkosResponse.Rajaongkir.Results, courierResult)
	}

	return rajaOngkirOngkosResponse, nil
}
