package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wtask/trend/internal/graph/model"
)

func (r *progressionRootQueryResolver) Percentage(ctx context.Context, obj *model.ProgressionRootQuery, first float64, percents float64) ([]*model.IncrementalProgressItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *progressionRootQueryResolver) Immutable(ctx context.Context, obj *model.ProgressionRootQuery, first float64) ([]*model.IncrementalProgressItem, error) {
	panic(fmt.Errorf("not implemented"))
}

// ProgressionRootQuery returns ProgressionRootQueryResolver implementation.
func (r *Resolver) ProgressionRootQuery() ProgressionRootQueryResolver {
	return &progressionRootQueryResolver{r}
}

type progressionRootQueryResolver struct{ *Resolver }
