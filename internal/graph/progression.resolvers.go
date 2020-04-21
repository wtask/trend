package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wtask/trend/internal/graph/model"
)

func (r *progressionQueryResolver) Percentage(ctx context.Context, obj *model.ProgressionQuery, first float64, percents float64) ([]*model.IncrementalProgressItem, error) {
	r.Logger.Println("progression.Percentage", "(", first, percents, ")", "operational.length:", r.ProgressionSettings.Length(), "parent obj:", obj)
	panic(fmt.Errorf("Percentage resolver not implemented"))
}

func (r *progressionQueryResolver) Immutable(ctx context.Context, obj *model.ProgressionQuery, first float64) ([]*model.IncrementalProgressItem, error) {
	oplen := r.ProgressionSettings.Length()
	r.Logger.Println("progression.Immutable", "(", first, ")", "operational.length:", oplen, "parent obj:", obj)
	if oplen < 0 || oplen > r.ProgressionSettings.LengthRange[1] {
		return nil, fmt.Errorf("cannot compute immutable, invalid operational length (%d)", oplen)
	}
	result := make([]*model.IncrementalProgressItem, 0, oplen)
	for i := 0; i < oplen; i++ {
		result = append(result, &model.IncrementalProgressItem{
			Index: i,
			Value: first,
		})
	}
	return result, nil
}

// ProgressionQuery returns ProgressionQueryResolver implementation.
func (r *Resolver) ProgressionQuery() ProgressionQueryResolver { return &progressionQueryResolver{r} }

type progressionQueryResolver struct{ *Resolver }
