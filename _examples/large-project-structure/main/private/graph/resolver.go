package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"

	"github.com/99designs/gqlgen/_examples/large-project-structure/main/private/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type PrivateIntegration interface {
	FetchCrowdStrikeHostDetail(ctx context.Context, connectorID int32) (*model.GeneralMessage, error)
	FetchCrowdStrikeHostDetailDirect(ctx context.Context, connectorID int32) ([]*model.HostDetail, error)
	FetchCrowdStrikeIntel(ctx context.Context, input *model.FetchCrowdStrikeIntelInput) (*model.GeneralMessage, error)
}

type Resolver struct {
	PrivateIntegration
}
