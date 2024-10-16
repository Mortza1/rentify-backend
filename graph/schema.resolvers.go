package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"rentify-backend/graph/generated"
	"rentify-backend/graph/models"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, name string, email string, password string, location string) (*models.User, error) {
	panic(fmt.Errorf("not implemented: RegisterUser - registerUser"))
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, name string, description *string, category *string, price float64, depositRequired bool) (*models.Item, error) {
	panic(fmt.Errorf("not implemented: CreateItem - createItem"))
}

// RentItem is the resolver for the rentItem field.
func (r *mutationResolver) RentItem(ctx context.Context, itemID string, startDate string, endDate string) (*models.Rental, error) {
	panic(fmt.Errorf("not implemented: RentItem - rentItem"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Items is the resolver for the items field.
func (r *queryResolver) Items(ctx context.Context) ([]*models.Item, error) {
	panic(fmt.Errorf("not implemented: Items - items"))
}

// Rentals is the resolver for the rentals field.
func (r *queryResolver) Rentals(ctx context.Context) ([]*models.Rental, error) {
	panic(fmt.Errorf("not implemented: Rentals - rentals"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }