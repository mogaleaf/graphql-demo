package resolver

import (
	"graph-ql-meetup/model"
	"graph-ql-meetup/data"
	"errors"
)

//Mutation AddWine
func (r *Resolver) AddWine(args *struct {
	Wine *WineInput
}) *wineResolver {
	wineCreated := &model.Wine{
		Name:    args.Wine.Name,
		Country: args.Wine.Country,
		Color:   args.Wine.Color,
		Id:      data.WineSeq,
	}
	data.WineList = append(data.WineList, wineCreated)
	data.WineSeq++
	return &wineResolver{wineCreated}
}

//Mutation AddReview
func (r *Resolver) AddReview(args *struct {
	WineId int32
	Review *ReviewInput
}) (*reviewResolver, error) {
	exists := existWine(args.WineId)
	if !exists {
		return nil,errors.New("Wine does not exist")
	}
	reviewCreated := &model.Review{
		Name:   args.Review.Name,
		Text:   args.Review.Text,
		Rate:   args.Review.Rate,
		WineId: args.WineId,
	}
	data.Reviews[args.WineId] = append(data.Reviews[args.WineId], reviewCreated)
	return &reviewResolver{reviewCreated}, nil
}


//Input object from graphql schema
type WineInput struct {
	Name    string
	Color   model.WineColor
	Country string
}

type ReviewInput struct {
	Name string
	Text string
	Rate int32
}


func existWine(wineId int32) (bool){
	for _,k := range data.WineList {
		if k.Id == wineId {
			return true
		}
	}
	return false
}