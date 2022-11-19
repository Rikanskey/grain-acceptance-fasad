package dao

import "grain-acceptance/internal/app"

func unmarshallAppConsignment(c consignment) app.Consignment {
	return app.Consignment{
		Id:                      c.Id,
		GrainCard:               unmarshallAppGrainCard(c.GrainCard),
		TransportModel:          c.TransportModel,
		TransportNumber:         c.TransportNumber,
		FirstNameDriver:         c.FirstNameDriver,
		SecondNameDriver:        c.SecondNameDriver,
		MiddleNameDriver:        c.MiddleNameDriver,
		TrailerNumber:           c.TrailerNumber,
		GrainType:               unmarshallAppGrainType(c.GrainType),
		GrainImpurity:           unmarshallAppPartOfImpurityParameters(c.GrainImpurity),
		VolatilesImpurity:       unmarshallAppPartOfImpurityParameters(c.VolatilesImpurity),
		Moisture:                c.Moisture,
		Origin:                  c.Origin,
		Nature:                  c.Nature,
		Color:                   c.Color,
		Smell:                   c.Smell,
		SmallGrainsPercent:      c.SmallGrainsPercent,
		Vitreous:                c.Vitreous,
		Gluten:                  c.Gluten,
		GostCulture:             c.GostCulture,
		Filminess:               c.Filminess,
		Contamitation:           c.Contamitation,
		TypeSubtypeComposition:  c.TypeSubtypeComposition,
		Core:                    c.Core,
		Bug:                     c.Bug,
		AdditionalNotes:         c.AdditionalNotes,
		AcceptedGrossWeight:     c.AcceptedGrossWeight,
		AcceptedTransportWeight: c.AcceptedTransportWeight,
		SentGrossWeight:         c.SentGrossWeight,
		SentTransportWeight:     c.SentTransportWeight,
	}
}

func unmarshallAppGrainCard(gc grainCard) app.GrainCard {
	return app.GrainCard{
		Id:          gc.Id,
		Customer:    unmarshallAppCompany(gc.Customer),
		Sender:      unmarshallAppCompany(gc.Sender),
		ProcessType: gc.ProcessType,
	}
}

func unmarshallAppCompany(c company) app.Company {
	return app.Company{
		Id:       c.Id,
		Name:     c.Name,
		Postcode: c.Postcode,
		State:    c.State,
		Street:   c.Street,
		House:    c.House,
		Office:   c.Office,
	}
}

func unmarshallAppGrainType(gt *grainType) *app.GrainType {
	if gt != nil {
		return &app.GrainType{
			Id:       gt.Id,
			Name:     gt.Name,
			GostName: gt.GostName,
			Sort:     gt.Sort,
		}
	}
	return nil
}

func unmarshallAppPartOfImpurityParameters(poips []partOfImpurityParameter) []app.PartOfImpurityParameter {
	result := make([]app.PartOfImpurityParameter, len(poips))
	for _, poip := range poips {
		result = append(result, unmarshallAppImpurityParameter(poip))
	}
	return result
}

func unmarshallAppImpurityParameter(poip partOfImpurityParameter) app.PartOfImpurityParameter {
	return app.PartOfImpurityParameter{
		Id:    poip.Id,
		Value: poip.Value,
		ImpurityParameter: app.ImpurityParameter{
			Id:   poip.ImpurityParameter.Id,
			Name: poip.ImpurityParameter.Name,
		},
		Part: app.Part{
			Id:   poip.Part.Id,
			Name: poip.Part.Name,
		},
	}
}
