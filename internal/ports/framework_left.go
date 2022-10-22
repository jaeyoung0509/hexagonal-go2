package ports

import (
	"context"
	"hexagonal-go/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationalParameters) (*pb.Answer, error)
	GetSubtraction(ctx context.Context, req *pb.OperationalParameters) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationalParameters) (*pb.Answer, error)
	GetDivision(ctx context.Context, req *pb.OperationalParameters) (*pb.Answer, error)
}
