package model

type Item struct {
	ID          string `json:"id" `
	Type        string `json:"type" `
	Weight      int    `json:"weigth" `
	Temperature int    `json:"temperature" `
	Volume      Volume `json:"volume" `
}

func (i *Item) GetVolume() int {
	if i.Type == "L" {
		return 1
	}
	return i.Volume.GetVolume()
}
