package resolver

import (
	"graph-ql-meetup/model"
	"graph-ql-meetup/data"
)

//Resolve Review object from schema file
type reviewResolver struct {
	r *model.Review
}

//Review property
func (r *reviewResolver) Name() string {
	return r.r.Name
}

//Review property
func (r *reviewResolver) Text() string {
	return r.r.Text
}

//Review property
func (r *reviewResolver) Rate() int32 {
	return r.r.Rate
}

//Retrieve Wine from an other datasource
func (r *reviewResolver) Wine() *wineResolver {
	for _,k := range  data.WineList {
		if k.Id == r.r.WineId {
			return &wineResolver{k}
		}
	}
	return nil
}
