package dao

import (
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
	"time"
)

type (
	company struct {
		Id       uint   `db:"id"`
		Name     string `db:"name"`
		Postcode uint   `db:"postcode"`
		State    string `db:"state"`
		City     string `db:"city"`
		Street   string `db:"street"`
		House    string `db:"house"`
		Office   string `db:"office"`
		//SenderGrainCards   []GrainCard
		//CustomerGrainCards []GrainCard
	}

	part struct {
		Id   uint   `db:"id"`
		Name string `db:"name"`
	}

	impurityParameter struct {
		Id   uint   `db:"id"`
		Name string `db:"name"`
	}

	partOfImpurityParameter struct {
		Id                uint `db:"id"`
		Value             int  `db:"name"`
		ImpurityParameter impurityParameter
		Part              part
	}

	grainType struct {
		Id       uint   `db:"id"`
		Name     string `db:"name"`
		GostName string `db:"gost_name"`
		Sort     string `db:"sort"`
	}

	consignment struct {
		Id                      uint `db:"id"`
		GrainCard               grainCard
		TransportModel          string  `db:"model_transport"`
		TransportNumber         string  `db:"number_transport"`
		FirstNameDriver         string  `db:"first_name_driver"`
		SecondNameDriver        string  `db:"second_name_driver"`
		MiddleNameDriver        string  `db:"middle_name_driver"`
		TrailerNumber           *string `db:"trailer_number, omitempty"`
		SentGrossWeight         float64 `db:"sent_gross_weight"`
		SentTransportWeight     float64 `db:"sent_transport_weight"`
		GrainType               *grainType
		VolatilesImpurity       []partOfImpurityParameter
		GrainImpurity           []partOfImpurityParameter
		Moisture                *float64 `db:"moisture, omitempty"`
		Origin                  *string  `db:"origin, omitempty"`
		Nature                  *int     `db:"nature, omitempty"`
		FallNumber              *float64 `db:"fall_number, omitempty"`
		Color                   *string  `db:"color, omitempty"`
		Smell                   *string  `db:"smell, omitempty"`
		SmallGrainsPercent      *float64 `db:"small_grains_percent, omitempty"`
		Vitreous                *float64 `db:"vitreous, omitempty"`
		Gluten                  *float64 `db:"gluten, omitempty"`
		GostCulture             *float64 `db:"gost_culture, omitempty"`
		Filminess               *float64 `db:"filminess, omitempty"`
		Contamitation           *float64 `db:"contamitation, omitempty"`
		TypeSubtypeComposition  *float64 `db:"type_subtype_composition, omitempty"`
		Core                    *float64 `db:"core, omitempty"`
		Bug                     *float64 `db:"bug, omitempty"`
		AdditionalNotes         *string  `db:"additional_notes, omitempty"`
		AcceptedGrossWeight     *float64 `db:"accepted_gross_weight, omitempty"`
		AcceptedTransportWeight *float64 `db:"accepted_transport_weight, omitempty"`
	}

	grainCard struct {
		Id          uint `db:"id"`
		Customer    company
		Sender      company
		Date        time.Time
		Storage     uint
		ProcessType string `db:"process_type"`
	}
)

func marshallAppConsignment(c app.Consignment) consignment {
	return consignment{
		Id: c.Id,
		//GrainCard: marshallGrainCard(c.GrainCard),
		TransportModel:   c.TransportModel,
		TransportNumber:  c.TransportNumber,
		FirstNameDriver:  c.FirstNameDriver,
		SecondNameDriver: c.SecondNameDriver,
		MiddleNameDriver: c.MiddleNameDriver,
	}
}

func marshallConsignment(c domain.Consignment) consignment {
	volatilesImpurity, grainImpurity := make([]partOfImpurityParameter, len(c.VolatilesImpurity())),
		make([]partOfImpurityParameter, len(c.GrainImpurity()))

	for _, el := range c.GrainImpurity() {
		grainImpurity = append(grainImpurity, marshallPartOfImpurityParameter(el))
	}

	for _, el := range c.VolatilesImpurity() {
		volatilesImpurity = append(volatilesImpurity, marshallPartOfImpurityParameter(el))
	}

	return consignment{
		Id:                      c.Id(),
		GrainCard:               marshallGrainCard(c.GrainCard()),
		TransportModel:          c.TransportModel(),
		TransportNumber:         c.TransportNumber(),
		FirstNameDriver:         c.FirstNameDriver(),
		SecondNameDriver:        c.SecondNameDriver(),
		MiddleNameDriver:        c.MiddleNameDriver(),
		TrailerNumber:           c.TrailerNumber(),
		SentGrossWeight:         c.SentGrossWeight(),
		SentTransportWeight:     c.SentTransportWeight(),
		GrainImpurity:           grainImpurity,
		VolatilesImpurity:       volatilesImpurity,
		Moisture:                c.Moisture(),
		Origin:                  c.Origin(),
		Nature:                  c.Nature(),
		FallNumber:              c.FallNumber(),
		Color:                   c.Color(),
		Smell:                   c.Smell(),
		SmallGrainsPercent:      c.SmallGrainsPercent(),
		Vitreous:                c.Vitreous(),
		Gluten:                  c.Gluten(),
		GostCulture:             c.GostCulture(),
		Filminess:               c.Filminess(),
		Contamitation:           c.Contamitation(),
		TypeSubtypeComposition:  c.TypeSubtypeComposition(),
		Core:                    c.Core(),
		Bug:                     c.Bug(),
		AdditionalNotes:         c.AdditionalNotes(),
		AcceptedGrossWeight:     c.AcceptedGrossWeight(),
		AcceptedTransportWeight: c.AcceptedTransportWeight(),
	}
}

func marshallGrainCard(gc domain.GrainCard) grainCard {
	return grainCard{
		Id:          gc.Id(),
		Customer:    marshallCompany(gc.Customer()),
		Sender:      marshallCompany(gc.Sender()),
		Date:        gc.Date(),
		ProcessType: gc.ProcessType().String(),
	}
}

func marshallCompany(c domain.Company) company {
	return company{
		Id:       c.Id(),
		Name:     c.Name(),
		Postcode: c.Postcode(),
		State:    c.State(),
		City:     c.City(),
		Street:   c.Street(),
		House:    c.House(),
		Office:   c.Office(),
	}
}

func marshallPartOfImpurityParameter(poip domain.PartOfImpurityParameter) partOfImpurityParameter {
	return partOfImpurityParameter{
		Id:                poip.Id(),
		Value:             poip.Value(),
		Part:              marshallPart(poip.Part()),
		ImpurityParameter: marshallImpurityParameter(poip.ImpurityParameter()),
	}
}

func marshallPart(p domain.Part) part {
	return part{
		Id:   p.Id(),
		Name: p.Name(),
	}
}

func marshallImpurityParameter(ip domain.ImpurityParameter) impurityParameter {
	return impurityParameter{
		Id:   ip.Id(),
		Name: ip.Name(),
	}
}
