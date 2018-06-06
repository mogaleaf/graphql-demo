package model

type WineColor string

const (
	RED   WineColor = "RED"
	WHITE WineColor = "WHITE"
	ROSE  WineColor = "ROSE"
)

//Wine object from Graphql schema
type Wine struct {
	Id      int32
	Color   WineColor
	Name    string
	Country string
	Reviews []*Review
	Rate    float64
}
