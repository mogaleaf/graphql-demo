package resolver

import (
	"graph-ql-meetup/model"
	"graph-ql-meetup/data"
)

type reviewResolver struct {
	r *model.Review
}

func (r *reviewResolver) Name() string {
	return r.r.Name
}
func (r *reviewResolver) Text() string {
	return r.r.Text
}
func (r *reviewResolver) Rate() int32 {
	return r.r.Rate
}

func (r *reviewResolver) Wine() *wineResolver {
	for _,k := range  data.WineList {
		if k.Id == r.r.WineId {
			return &wineResolver{k}
		}
	}
	return nil
}
