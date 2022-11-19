package commands

import (
	"context"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
)

type UpdateConsignmentTransportWeightHandler struct {
	consignmentRepository consignmentRepository
}

func NewUpdateConsignmentTransportWeightHandler(repository consignmentRepository) UpdateConsignmentTransportWeightHandler {
	if repository == nil {
		panic("repository is nil")
	}
	return UpdateConsignmentTransportWeightHandler{
		consignmentRepository: repository,
	}
}

func (h UpdateConsignmentTransportWeightHandler) Handle(ctx context.Context,
	q app.UpdateConsignmentTransportWeightQuery) error {
	consignment, err := h.consignmentRepository.GetConsignmentById(ctx, q.ConsignmentId)
	if err != nil {
		return app.ErrConsignmentDoesNotExist
	}

	cons := domain.Consignment{}
	cons.InitByApp(consignment)

	return nil
}
