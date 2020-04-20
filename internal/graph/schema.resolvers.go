package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/wtask/trend/internal/graph/model"
)

func (r *queryResolver) Progression(ctx context.Context) (*model.ProgressionQuery, error) {
	// Это корневой ресолвер для query-поля "progression"
	// Если здесь вернуть ошибку или вызвать панику, то дальнейшая маршрутизация остановится
	return &model.ProgressionQuery{}, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
