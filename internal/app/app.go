package app

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type (
	Commands struct {
		UpdateConsignmentGrossWeight     updateConsignmentGrossWeight
		UpdateConsignmentAnalyse         updateConsignmentAnalyse
		UpdateConsignmentTransportWeight updateConsignmentTransportWeight
	}
	updateConsignmentGrossWeight interface {
		Handle(ctx context.Context, q UpdateConsignmentGrossWeightQuery) error
	}
	updateConsignmentAnalyse interface {
		Handle(ctx context.Context, q UpdateConsignmentAnalyseQuery) error
	}
	updateConsignmentTransportWeight interface {
		Handle(ctx context.Context, q UpdateConsignmentTransportWeightQuery) error
	}
)

type (
	Queries struct {
		GetConsignment getConsignment
	}
	getConsignment interface {
		Handle(ctx context.Context, q GetConsignmentQuery) (Consignment, error)
	}
)
