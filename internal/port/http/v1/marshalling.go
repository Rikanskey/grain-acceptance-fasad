package v1

import (
	"github.com/go-chi/render"
	"grain-acceptance/internal/app"
	"net/http"
)

func marshallConsignment(w http.ResponseWriter, r *http.Request, consignment app.Consignment) {
	response := marshallConsignmentToResponse(consignment)

	render.Respond(w, r, response)
}

func marshallConsignmentToResponse(consignment app.Consignment) ConsignmentResponse {
	c := ConsignmentResponse{
		Id: int(consignment.Id),
		//gc
		GrainType:               marshallGrainType(consignment.GrainType),
		TransportModel:          consignment.TransportModel,
		TransportNumber:         consignment.TransportNumber,
		FirstNameDriver:         consignment.FirstNameDriver,
		SecondNameDriver:        consignment.SecondNameDriver,
		MiddleNameDriver:        consignment.MiddleNameDriver,
		TrailerNumber:           consignment.TrailerNumber,
		Moisture:                consignment.Moisture,
		Origin:                  consignment.Origin,
		Nature:                  consignment.Nature,
		Color:                   consignment.Color,
		Smell:                   consignment.Smell,
		SmallGrainsPercent:      consignment.SmallGrainsPercent,
		Vitreous:                consignment.Vitreous,
		Gluten:                  consignment.Gluten,
		GostCulture:             consignment.GostCulture,
		Filminess:               consignment.Filminess,
		Contamitation:           consignment.Contamitation,
		TypeSubtypeComposition:  consignment.TypeSubtypeComposition,
		Core:                    consignment.Core,
		Bug:                     consignment.Bug,
		AdditionalNotes:         consignment.AdditionalNotes,
		AcceptedGrossWeight:     consignment.AcceptedGrossWeight,
		AcceptedTransportWeight: consignment.AcceptedTransportWeight,
		SentGrossWeight:         consignment.SentGrossWeight,
		SentTransportWeight:     consignment.SentTransportWeight,
	}

	volatilesImpurityResponse := make([]PartOfImpurityResponse, len(consignment.VolatilesImpurity))
	for _, el := range consignment.VolatilesImpurity {
		volatilesImpurityResponse = append(volatilesImpurityResponse, marshallPartOfImpurityParameterToResponse(el))
	}

	grainImpurityResponse := make([]PartOfImpurityResponse, len(consignment.GrainImpurity))
	for _, el := range consignment.GrainImpurity {
		grainImpurityResponse = append(grainImpurityResponse, marshallPartOfImpurityParameterToResponse(el))
	}

	return c
}

func marshallGrainType(gt *app.GrainType) *GrainTypeResponse {
	if gt != nil {
		return &GrainTypeResponse{
			Id:       int(gt.Id),
			Name:     gt.Name,
			GostName: gt.GostName,
			Sort:     gt.Sort,
		}
	}
	return nil
}

func marshallPartOfImpurityParameterToResponse(poip app.PartOfImpurityParameter) PartOfImpurityResponse {
	return PartOfImpurityResponse{
		Id:                int(poip.Id),
		Value:             poip.Value,
		Part:              marshallPartToResponse(poip.Part),
		ImpurityParameter: marshallImpurityParameter(poip.ImpurityParameter),
	}
}

func marshallPartToResponse(p app.Part) PartResponse {
	return PartResponse{
		Id:   int(p.Id),
		Name: p.Name,
	}
}

func marshallImpurityParameter(ip app.ImpurityParameter) ImpurityParameterResponse {
	return ImpurityParameterResponse{
		Id:   int(ip.Id),
		Name: ip.Name,
	}
}