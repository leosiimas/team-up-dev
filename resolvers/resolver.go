//go:generate go run github.com/99designs/gqlgen --verbose
package resolvers

import (
	"github.com/go-pg/pg/v9"
	"github.com/leosiimas/team-up-dev/graph/generated"
)

type Resolver struct {
	DB *pg.DB
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

