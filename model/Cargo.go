package model

type CargoTemperature struct {
	MinTemp int
	MaxTemp int
}

type Door struct {
	Height int
	Width  int
}

type Cargo struct {
	MaxWeigth       int
	CurrentWeigth   int
	Volume          Volume
	CurrentVolume   int
	Temperature     CargoTemperature
	Door            Door
	RegisteredItems []Item
}
