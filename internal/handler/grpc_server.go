package handler

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"mock-cbr/internal/proto"
	"mock-cbr/internal/service"
)

type CurrencyServiceServer struct {
	proto.UnimplementedCurrencyServiceServer
}

func (s *CurrencyServiceServer) GetDaily(ctx context.Context, req *proto.DailyRequest) (*proto.DailyResponse, error) {
	if req.Date == "error" { // Имитация ошибки
		return nil, status.Error(codes.Internal, "mocked internal server error")
	}

	xmlStr, valuteList, err := service.GenerateRandomData(req.Date)
	if err != nil {
		return nil, status.Error(codes.Internal, "generation failed")
	}

	protoValutes := make([]*proto.Valute, 0, len(valuteList))
	for _, v := range valuteList {
		protoValutes = append(protoValutes, &proto.Valute{
			Id:       v.ID,
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Nominal:  int32(v.Nominal),
			Name:     v.Name,
			Value:    v.Value,
		})
	}

	return &proto.DailyResponse{
		Xml:     xmlStr,
		Valutes: protoValutes,
	}, nil
}
