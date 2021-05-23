package model

type Volume struct {
	Height int `json:"height" `
	Width  int `json:"width" `
	Length int `json:"length" `
}

func (v *Volume) GetVolume() int {
	return v.Height * v.Length * v.Width
}
