package types

func NewSize(width, height int) *Size {
	return &Size{
		Width:  width,
		Height: height,
	}
}

type Size struct {
	Width  int `bson:"width,omitempty" json:"width,omitempty"`
	Height int `bson:"height,omitempty" json:"height,omitempty"`
}
