package model


//Review object from Graphql schema
type Review struct {
	Name string
	Text string
	Rate int32
	WineId int32
}
