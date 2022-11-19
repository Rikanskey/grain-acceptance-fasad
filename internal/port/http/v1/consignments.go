package v1

import (
	"github.com/pkg/errors"
	"grain-acceptance/internal/app"
	"grain-acceptance/pkg/httperr"
	"net/http"
)

func (h handler) GetConsignment(w http.ResponseWriter, r *http.Request, consignmentId string) {
	q, ok := unmarshalConsignment(w, r, consignmentId)
	if !ok {
		return
	}

	c, err := h.app.Queries.GetConsignment.Handle(r.Context(), q)

	if err == nil {
		marshallConsignment(w, r, c)
	}

}

func (h handler) UpdateConsignmentAnalyse(w http.ResponseWriter, r *http.Request, consignmentId string) {
	q, ok := unmarshalConsignmentAnalyse(w, r, consignmentId)
	if !ok {
		return
	}

	err := h.app.Commands.UpdateConsignmentAnalyse.Handle(r.Context(), q)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if errors.Is(err, app.ErrConsignmentDoesNotExist) {
		httperr.NotFound("consignment-not-found", err, w, r)

		return
	}
	//err := h.app.Commands.
	httperr.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) UpdateConsignmentGrossWeight(w http.ResponseWriter, r *http.Request, consignmentId string) {
	q, ok := unmarshalConsignmentGrossWeight(w, r, consignmentId)
	if !ok {
		return
	}

	err := h.app.Commands.UpdateConsignmentGrossWeight.Handle(r.Context(), q)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if errors.Is(err, app.ErrConsignmentDoesNotExist) {
		httperr.NotFound("consignment-not-found", err, w, r)

		return
	}

	httperr.InternalServerError("unexpected-error", err, w, r)
}

func (h handler) UpdateConsignmentTransportWeight(w http.ResponseWriter, r *http.Request, consignmentId string) {
	q, ok := unmarshalConsignmentTransportWeight(w, r, consignmentId)
	if !ok {
		return
	}

	err := h.app.Commands.UpdateConsignmentTransportWeight.Handle(r.Context(), q)

	if err == nil {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if errors.Is(err, app.ErrConsignmentDoesNotExist) {
		httperr.NotFound("consignment-not-found", err, w, r)

		return
	}

	httperr.InternalServerError("unexpected-error", err, w, r)
}
