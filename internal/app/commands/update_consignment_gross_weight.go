package commands

import (
	"context"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
)

//type updateConsignmentGrossWeightModel interface {
//	UpdateConsignmentGrossWeight(ctx context.Context)
//}

type UpdateConsignmentGrossWeightHandler struct {
	consignmentRepository consignmentRepository
}

func NewUpdateConsignmentGrossWeightHandler(repository consignmentRepository) UpdateConsignmentGrossWeightHandler {
	if repository == nil {
		panic("repository is nil")
	}
	return UpdateConsignmentGrossWeightHandler{
		consignmentRepository: repository,
	}
}

func (h UpdateConsignmentGrossWeightHandler) Handle(ctx context.Context,
	q app.UpdateConsignmentGrossWeightQuery) error {
	consignment, err := h.consignmentRepository.GetConsignmentById(ctx, q.ConsignmentId)
	if err != nil {
		return app.ErrConsignmentDoesNotExist
	}

	cons := domain.Consignment{}
	cons.InitByApp(consignment)

	return nil
}
