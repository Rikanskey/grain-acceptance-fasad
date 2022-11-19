package v1

import (
	"github.com/go-chi/render"
	"grain-acceptance/internal/app"
	"grain-acceptance/pkg/httperr"
	"net/http"
)

func decode(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if err := render.Decode(r, v); err != nil {
		httperr.BadRequest("bad-request", err, w, r)
		return false
	}

	return true
}

func unmarshalConsignment(
	w http.ResponseWriter, r *http.Request,
	consignmentId string) (q app.GetConsignmentQuery, ok bool) {
	return app.GetConsignmentQuery{
		ConsignmentId: consignmentId,
	}, true
}

func unmarshalConsignmentAnalyse(
	w http.ResponseWriter, r *http.Request,
	consignmentId string) (q app.UpdateConsignmentAnalyseQuery, ok bool) {
	//var updateConsignmentAnalyseQuery app.UpdateConsignmentAnalyseQuery
	var analyseUpdateRequest AnalyseUpdateRequest
	if ok = decode(w, r, &analyseUpdateRequest); !ok {
		return app.UpdateConsignmentAnalyseQuery{}, ok
	}

	return app.UpdateConsignmentAnalyseQuery{
		ConsignmentId:          consignmentId,
		AdditionalNotes:        *analyseUpdateRequest.AdditionalNotes,
		Bug:                    *analyseUpdateRequest.Bug,
		Color:                  *analyseUpdateRequest.Color,
		Filminess:              *analyseUpdateRequest.Filminess,
		Gluten:                 *analyseUpdateRequest.Gluten,
		GostCulture:            *analyseUpdateRequest.GostCulture,
		Moisture:               *analyseUpdateRequest.Moisture,
		Nature:                 *analyseUpdateRequest.Nature,
		Origin:                 *analyseUpdateRequest.Origin,
		SmallGrainsPercent:     *analyseUpdateRequest.SmallGrainsPercent,
		Smell:                  *analyseUpdateRequest.Smell,
		TypeSubtypeComposition: *analyseUpdateRequest.TypeSubtypeComposition,
		Vitreous:               *analyseUpdateRequest.Vitreous,
	}, ok
}

func unmarshalConsignmentGrossWeight(
	w http.ResponseWriter, r *http.Request,
	consignmentId string) (q app.UpdateConsignmentGrossWeightQuery, ok bool) {
	var grossWeightUpdateRequest GrossWeightUpdateRequest
	if ok = decode(w, r, &grossWeightUpdateRequest); !ok {
		return app.UpdateConsignmentGrossWeightQuery{}, ok
	}

	return app.UpdateConsignmentGrossWeightQuery{
		ConsignmentId: consignmentId,
		GrossWeight:   grossWeightUpdateRequest.Value,
	}, ok
}

func unmarshalConsignmentTransportWeight(
	w http.ResponseWriter, r *http.Request,
	consignmentId string) (q app.UpdateConsignmentTransportWeightQuery, ok bool) {
	var transportWeightUpdateRequest TransportWeightUpdateRequest
	if ok = decode(w, r, &transportWeightUpdateRequest); !ok {
		return app.UpdateConsignmentTransportWeightQuery{}, ok
	}

	return app.UpdateConsignmentTransportWeightQuery{
		ConsignmentId:   consignmentId,
		TransportWeight: transportWeightUpdateRequest.Value,
	}, ok
}
