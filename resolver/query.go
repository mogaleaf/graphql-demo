package resolver

import (
	"graph-ql-meetup/model"
	"graph-ql-meetup/data"
)

//Root Query resolver
type Resolver struct{}

//Query Search Wines
func (r *Resolver) SearchWines(args struct{ Color *model.WineColor }) []*wineResolver {
	var returnList []*wineResolver
	for _, k := range data.WineList {
		if args.Color == nil || k.Color == *args.Color {
			returnList = append(returnList, &wineResolver{k})
		}
	}
	return returnList
}

//Query Search Reviews
func (r *Resolver) SearchReviews(args struct{ MoreThanRate int32 }) []*reviewResolver {
	var returnList []*reviewResolver
	for _, k := range data.Reviews {
		for _, j  := range k {
			if j.Rate > args.MoreThanRate {
				returnList = append(returnList,&reviewResolver{j})
			}
		}
	}
	return returnList
}
