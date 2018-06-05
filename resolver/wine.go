package resolver

import (
	"graph-ql-meetup/model"
	"graph-ql-meetup/data"
)

type wineResolver struct {
	w *model.Wine
}

func (r *wineResolver) Id() int32 {
	return r.w.Id
}

func (r *wineResolver) Color() *model.WineColor {
	return &r.w.Color
}
func (r *wineResolver) Name() string {
	return r.w.Name
}
func (r *wineResolver) Country() string {
	return r.w.Country
}

func (r *wineResolver) Reviews() []*reviewResolver {
	var reviews []*reviewResolver
	list := data.Reviews[r.w.Id]
	if list == nil || len(list) == 0 {
		return reviews
	}
	for _, k := range data.Reviews[r.w.Id] {
		reviews = append(reviews, &reviewResolver{k})
	}
	return reviews
}

func (r *wineResolver) Rate() float64 {
	var finalRate float64
	reviews := data.Reviews[r.w.Id]
	if reviews == nil || len(reviews) == 0 {
		return -1
	}
	for _, k := range reviews {
		finalRate += float64(k.Rate)
	}
	return finalRate / float64(len(reviews));
}
