package dao

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
	"strconv"
)

type ImpurityType string

const (
	VOLATILES_IMPURITY_ID ImpurityType = "1"
	GRAIN_IMPURITY_ID     ImpurityType = "2"
)

type ConsignmentRepository struct {
	consignments *sql.DB
}

func NewConsignmentRepository(db *sql.DB) *ConsignmentRepository {
	return &ConsignmentRepository{consignments: db}
}

func (c *ConsignmentRepository) GetConsignmentById(ctx context.Context, consignmentId string) (app.Consignment, error) {
	//stmt, err := c.consignments.Prepare("SELECT * FROM consignment WHERE id=$1")
	stmt, err := c.consignments.Prepare("SELECT id, model_transport, number_transport, first_name_driver, second_name_driver, middle_name_driver, trailer_number, moisture, origin, nature, fall_number, color, smell, small_grains_percent, vitreous, gluten, gost_culture, filminess, contamitation, type_subtype_composition, core, bug, additional_notes, accepted_gross_weight, accepted_transport_weight, sent_gross_weight, sent_transport_weight FROM consignment WHERE id=$1")
	defer stmt.Close()

	if err != nil {
		logrus.WithError(err).Println("Failed with \"Get consignment by id\" query")
		return app.Consignment{}, err
	}

	result, err := stmt.QueryContext(ctx, consignmentId)

	if err != nil {
		logrus.WithError(err)
		return app.Consignment{}, err
	}

	consignment := consignment{GrainCard: grainCard{}, GrainType: &grainType{}}
	result.Next()
	if err := result.Scan(&consignment.Id, &consignment.TransportModel, &consignment.TransportNumber,
		&consignment.FirstNameDriver, &consignment.SecondNameDriver, &consignment.MiddleNameDriver,
		&consignment.TrailerNumber, &consignment.Moisture, &consignment.Origin, &consignment.Nature, &consignment.FallNumber,
		&consignment.Color, &consignment.Smell, &consignment.SmallGrainsPercent, &consignment.Vitreous,
		&consignment.Gluten, &consignment.GostCulture, &consignment.Filminess, &consignment.Contamitation,
		&consignment.TypeSubtypeComposition, &consignment.Core, &consignment.Bug, &consignment.AdditionalNotes,
		&consignment.AcceptedGrossWeight, &consignment.AcceptedTransportWeight, &consignment.SentGrossWeight,
		&consignment.SentTransportWeight); err != nil {
		logrus.WithError(err).Println("Failed with \"Get consignment by id\" query")
		return unmarshallAppConsignment(consignment), err
	}

	consignment.GrainType, err = c.getGonsignmentGrainType(ctx, consignmentId)
	if err != nil {
		consignment.GrainType = nil
		//logrus.WithError(err)
		//return unmarshallAppConsignment(consignment), err
	}

	consignment.GrainCard, err = c.getConsignmentGrainCard(ctx, consignmentId)
	if err != nil {
		logrus.WithError(err)
		return unmarshallAppConsignment(consignment), err
	}

	consignment.VolatilesImpurity, err = c.getConsignmentPartOfImpurityParameters(ctx, consignmentId,
		VOLATILES_IMPURITY_ID)
	if err != nil {
		consignment.VolatilesImpurity = make([]partOfImpurityParameter, 0)
		//logrus.WithError(err)
		//return unmarshallAppConsignment(consignment), err
	}

	consignment.GrainImpurity, err = c.getConsignmentPartOfImpurityParameters(ctx, consignmentId, GRAIN_IMPURITY_ID)
	if err != nil {
		consignment.GrainImpurity = make([]partOfImpurityParameter, 0)
		//logrus.WithError(err)
		//return unmarshallAppConsignment(consignment), err
	}

	return unmarshallAppConsignment(consignment), nil
}

func (c *ConsignmentRepository) getGonsignmentGrainType(ctx context.Context, consignmentId string) (*grainType, error) {
	stmt, err := c.consignments.Prepare("SELECT * FROM grain_type JOIN consignment c on grain_type.id = c.grain_type_id WHERE c.id=$1 ")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return &grainType{}, err
	}

	result, err := stmt.QueryContext(ctx, consignmentId)
	if err != nil {
		logrus.WithError(err)
		return &grainType{}, err
	}

	gt := &grainType{}
	result.Next()
	if err := result.Scan(&gt.Id, &gt.Name, &gt.GostName, &gt.Sort); err != nil {
		//logrus.WithError(err)
		return gt, err
	}

	return gt, nil
}

func (c *ConsignmentRepository) getConsignmentGrainCard(ctx context.Context, consignmentId string) (grainCard, error) {
	//stmt, err := c.consignments.Prepare("SELECT * FROM grain_card JOIN consignment c on grain_card.id = c.grain_card_id WHERE c.id=$1")
	stmt, err := c.consignments.Prepare("SELECT grain_card.id, date, process_type FROM grain_card JOIN consignment c on grain_card.id = c.grain_card_id WHERE c.id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return grainCard{}, err
	}

	result, err := stmt.QueryContext(ctx, consignmentId)
	if err != nil {
		logrus.WithError(err)
		return grainCard{}, err
	}

	gc := grainCard{}
	result.Next()
	if err := result.Scan(&gc.Id, &gc.Date, &gc.ProcessType); err != nil {
		logrus.WithError(err)
		return gc, err
	}

	gc.Customer, err = c.getConsignmentGrainCardCustomer(ctx, strconv.Itoa(int(gc.Id)))
	if err != nil {
		logrus.WithError(err)
		return gc, err
	}

	gc.Sender, _ = c.getConsignmentGrainCardSender(ctx, strconv.Itoa(int(gc.Id)))
	if err != nil {
		logrus.WithError(err)
		return gc, err
	}

	return gc, nil
}

func (c *ConsignmentRepository) getConsignmentGrainCardCustomer(ctx context.Context, grainCardId string) (company, error) {
	stmt, err := c.consignments.Prepare("SELECT company.id, name, postcode, state, city, street, house, office FROM company JOIN grain_card gc on company.id = gc.customer_id WHERE gc.id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return company{}, err
	}

	result, err := stmt.QueryContext(ctx, grainCardId)
	if err != nil {
		logrus.WithError(err)
		return company{}, err
	}

	customer := company{}
	result.Next()
	if err := result.Scan(&customer.Id, &customer.Name, &customer.State, &customer.City, &customer.Street,
		&customer.House, &customer.Office, &customer.Postcode); err != nil {
		logrus.WithError(err)
		return customer, err
	}

	return customer, nil
}

func (c *ConsignmentRepository) getConsignmentGrainCardSender(ctx context.Context, grainCardId string) (company, error) {
	stmt, err := c.consignments.Prepare("SELECT * FROM company JOIN grain_card gc on company.id = gc.sender_id WHERE gc.id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return company{}, err
	}

	result, err := stmt.QueryContext(ctx, grainCardId)
	if err != nil {
		logrus.WithError(err)
		return company{}, err
	}

	sender := company{}
	result.Next()
	if err := result.Scan(&sender.Id, &sender.Name, &sender.State, &sender.City, &sender.Street,
		&sender.House, &sender.Office, &sender.Postcode); err != nil {
		logrus.WithError(err)
		return sender, err
	}

	return sender, nil
}

func (c *ConsignmentRepository) getConsignmentPartOfImpurityParameters(ctx context.Context, consignmentId string,
	impurityId ImpurityType) ([]partOfImpurityParameter, error) {
	stmt, err := c.consignments.Prepare("SELECT * FROM part_of_impurity_parameter poip JOIN consignment c on c.id = poip.consignment_id WHERE c.id=$1 AND poip.impurity_parameter_id=$2")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return make([]partOfImpurityParameter, 0), err
	}

	result, err := stmt.QueryContext(ctx, consignmentId, impurityId)
	if err != nil {
		logrus.WithError(err)
		return make([]partOfImpurityParameter, 0), err
	}

	ip, err := c.getImpurityParameterById(ctx, impurityId)
	if err != nil {
		logrus.WithError(err)
		return make([]partOfImpurityParameter, 0), err
	}

	var poips []partOfImpurityParameter
	for result.Next() {
		poip := partOfImpurityParameter{}
		result.Scan(&poip.Id, &poip.Value)
		poip.ImpurityParameter = ip
		poip.Part, err = c.getPartByPartOfImpurityParameter(ctx, strconv.Itoa(int(poip.Id)))
		if err != nil {
			logrus.WithError(err)
			return poips, err
		}
		poips = append(poips, poip)
	}

	return poips, nil
}

func (c *ConsignmentRepository) getImpurityParameterById(ctx context.Context, impurityParameterId ImpurityType) (impurityParameter, error) {
	stmt, err := c.consignments.Prepare("SELECT * FROM impurity_parameter WHERE id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return impurityParameter{}, err
	}

	result, err := stmt.QueryContext(ctx, impurityParameterId)
	if err != nil {
		logrus.WithError(err)
		return impurityParameter{}, err
	}

	ip := impurityParameter{}
	result.Next()
	if err := result.Scan(&ip.Id, &ip.Name); err != nil {
		logrus.WithError(err)
		return ip, err
	}

	return ip, nil
}

func (c *ConsignmentRepository) getPartByPartOfImpurityParameter(ctx context.Context, poipId string) (part, error) {
	stmt, err := c.consignments.Prepare("SELECT * FROM part JOIN part_of_impurity_parameter poip on part.id = poip.part_id WHERE poip.id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return part{}, err
	}

	result, err := stmt.QueryContext(ctx, poipId)
	if err != nil {
		logrus.WithError(err)
		return part{}, err
	}

	p := part{}
	result.Next()
	if err := result.Scan(&p.Id, &p.Name); err != nil {
		logrus.WithError(err)
		return p, err
	}

	return p, nil
}

func (c *ConsignmentRepository) UpdateConsignment(ctx context.Context, cons domain.Consignment) error {
	consignment := marshallConsignment(cons)
	stmt, err := c.consignments.Prepare("UPDATE consignment SET grain_card_id=$2, model_transport=$3, number_transport=$4, first_name_driver=$5, second_name_driver=$6, middle_name_driver=$7, trailer_number=$8, grain_type_id=$9, moisture=$10, origin=$11, nature=$12, fall_number=$13, color=$14, smell=$15, small_grains_percent=$16, vitreous=$17, gluten=$18, gost_culture=$19, filminess=$20, contamitation=$21, type_subtype_composition=$22, core=$23, bug=$24, additional_notes=$25, accepted_gross_weight=$26, accepted_transport_weight=$27, sent_gross_weight=$28, sent_transport_weight=$29 WHERE id=$1")
	defer stmt.Close()
	if err != nil {
		logrus.WithError(err)
		return err
	}

	var grainTypeId uint
	if consignment.GrainType != nil {
		grainTypeId = consignment.GrainType.Id
	} else {
		grainTypeId = 0
	}

	_, err = stmt.QueryContext(ctx, consignment.Id, consignment.GrainCard.Id, consignment.TransportModel,
		consignment.TransportNumber, consignment.FirstNameDriver, consignment.SecondNameDriver,
		consignment.MiddleNameDriver, consignment.TrailerNumber, grainTypeId, consignment.Moisture, consignment.Origin,
		consignment.Nature, consignment.FallNumber, consignment.Color, consignment.Smell,
		consignment.SmallGrainsPercent, consignment.Vitreous, consignment.Gluten, consignment.GostCulture,
		consignment.Filminess, consignment.Contamitation, consignment.TypeSubtypeComposition, consignment.Core,
		consignment.Bug, consignment.AdditionalNotes, consignment.AcceptedTransportWeight,
		consignment.AcceptedTransportWeight, consignment.SentGrossWeight, consignment.SentTransportWeight)

	if err != nil {
		logrus.WithError(err)
		return err
	}

	return nil
}
