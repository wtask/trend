package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wtask/trend/internal/graph/model"
)

func (r *progressionQueryResolver) Percentage(ctx context.Context, obj *model.ProgressionQuery, first float64, percents float64) ([]*model.IncrementalProgressItem, error) {
	panic(fmt.Errorf("Percentage resolver not implemented"))
}

func (r *progressionQueryResolver) Immutable(ctx context.Context, obj *model.ProgressionQuery, first float64) ([]*model.IncrementalProgressItem, error) {
	panic(fmt.Errorf("Immutable resolver not implemented"))
}

// ProgressionQuery returns ProgressionQueryResolver implementation.
func (r *Resolver) ProgressionQuery() ProgressionQueryResolver { return &progressionQueryResolver{r} }

type progressionQueryResolver struct{ *Resolver }
