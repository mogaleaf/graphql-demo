package model

type WineColor string

const (
	RED   WineColor = "RED"
	WHITE WineColor = "WHITE"
	ROSE  WineColor = "ROSE"
)

type Wine struct {
	Id      int32
	Color   WineColor
	Name    string
	Country string
	Reviews []*Review
	Rate    float64
}

type Review struct {
	Name string
	Text string
	Rate int32
	WineId int32
}

type ReviewInput struct {
	Name string
	Text string
	Rate float32
}
