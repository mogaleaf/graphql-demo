package data

import "graph-ql-meetup/model"

var WineSeq int32

var WineList = []*model.Wine{
	{
		Color:   model.RED,
		Id:      1,
		Name:    "Nuit Saint Georges",
		Country: "France",
	},
}

var Reviews = make(map[int32][]*model.Review)

func init() {
	WineSeq = 2
	Reviews[1] = []*model.Review{
		{
			Rate: 5,
			Name: "Cecile",
			Text: "Super good wine",
			WineId: 1,
		},
	}
}
