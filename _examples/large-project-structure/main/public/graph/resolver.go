package graph

import (
	"context"
	"github.com/99designs/gqlgen/_examples/large-project-structure/main/public/graph/model"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type ExternalQueryResolver interface {
	Tezz(ctx context.Context) (*model.Test, error)
	GetYaSome(ctx context.Context, input *model.CustomInput) ([]*model.CustomZeekIntel, error)
	GetAppCrowdStrike(ctx context.Context) (*model.AppCrowdStrike, error)
	AddIndicator(ctx context.Context, input model.IndicatorInput) (*model.Indicator, error)
}

type Resolver struct {
	ExternalQueryResolver
	executableSchema graphql.ExecutableSchema
}

// NewResolver initializes a Resolver with the given executable schema
func NewResolver(schema graphql.ExecutableSchema) *Resolver {
	return &Resolver{
		executableSchema: schema,
	}
}

// GetExecutableSchema returns the executable schema
func (r *Resolver) GetExecutableSchema() graphql.ExecutableSchema {
	return r.executableSchema
}

// Method to access the parsed schema
func (r *Resolver) GetParsedSchema() *ast.Schema {
	if e, ok := r.executableSchema.(*executableSchema); ok {
		return e.Schema()
	}
	return nil
}
