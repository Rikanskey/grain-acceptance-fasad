package queries

import (
	"context"
	"github.com/pkg/errors"
	"grain-acceptance/internal/app"
)

type getConsignmentModel interface {
	GetConsignmentById(ctx context.Context, consignmentId string) (app.Consignment, error)
}

type GetConsignmentHandler struct {
	getModel getConsignmentModel
}

func NewGetConsignmentHandler(getModel getConsignmentModel) GetConsignmentHandler {
	if getModel == nil {
		panic("GetModel is nil")
	}

	return GetConsignmentHandler{getModel: getModel}
}

func (h GetConsignmentHandler) Handle(ctx context.Context, query app.GetConsignmentQuery) (app.Consignment, error) {
	consignment, err := h.getModel.GetConsignmentById(ctx, query.ConsignmentId)

	return consignment, errors.Wrapf(err, "Getting consignment with id %d", consignment.Id)
}
