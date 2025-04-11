package graph

import (
	"context"

	"github.com/99designs/gqlgen/_examples/large-project-structure/main/public/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Define an interface for each resolver method
type ExternalQueryResolver interface {
	// Example query resolver
	Tezz(ctx context.Context) (*model.Test, error)
	// Example query resolver with args
	GetYaSome(context.Context, *model.CustomInput) ([]*model.CustomZeekIntel, error)

	GetAppCrowdStrike(ctx context.Context) (*model.AppCrowdStrike, error)
	AddIndicator(ctx context.Context, input model.IndicatorInput) (*model.Indicator, error)
}

type Resolver struct {
	ExternalQueryResolver
}
