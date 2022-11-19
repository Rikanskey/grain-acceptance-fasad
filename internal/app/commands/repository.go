package commands

import (
	"context"
	"grain-acceptance/internal/app"
	"grain-acceptance/internal/domain"
)

type consignmentRepository interface {
	GetConsignmentById(ctx context.Context, consignmentId string) (app.Consignment, error)
	UpdateConsignment(ctx context.Context, consignment domain.Consignment) error
}
