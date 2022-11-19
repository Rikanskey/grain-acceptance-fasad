package commands

import (
	"context"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
)

//type updateConsignmentAnalyse interface {
//	updateConsignment(ctx context.Context, consignment app.Consignment) error
//}

type UpdateConsignmentAnalyseHandler struct {
	consignmentRepository consignmentRepository
}

func NewUpdateConsignmentAnalyseHandler(repository consignmentRepository) UpdateConsignmentAnalyseHandler {
	if repository == nil {
		panic("repository is nil")
	}

	return UpdateConsignmentAnalyseHandler{
		consignmentRepository: repository,
	}
}

func (h UpdateConsignmentAnalyseHandler) Handle(ctx context.Context,
	query app.UpdateConsignmentAnalyseQuery) error {

	consignment, err := h.consignmentRepository.GetConsignmentById(ctx, query.ConsignmentId)
	if err != nil {
		return app.ErrConsignmentDoesNotExist
	}

	cons := domain.Consignment{}
	cons.InitByApp(consignment)
	ap := domain.AnalyseParams{
		AdditionalNotes:        query.AdditionalNotes,
		Bug:                    query.Bug,
		Color:                  query.Color,
		Contamitation:          query.Contamitation,
		Core:                   query.Core,
		Filminess:              query.Filminess,
		Gluten:                 query.Gluten,
		GostCulture:            query.GostCulture,
		Moisture:               query.Moisture,
		Nature:                 query.Nature,
		Origin:                 query.Origin,
		SmallGrainsPercent:     query.SmallGrainsPercent,
		Smell:                  query.Smell,
		TypeSubtypeComposition: query.TypeSubtypeComposition,
		Vitreous:               query.Vitreous,
	}
	cons.UpdateAnalyses(ap)

	err = h.consignmentRepository.UpdateConsignment(ctx, cons)
	if err != nil {
		return err
	}

	return nil
}
