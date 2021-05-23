package service

import (
	"errors"
	"fmt"
	"zanuarmirza/goFiberSample/model"
)

type CargoService struct {
	Cargo model.Cargo
}

func NewCargoService(cargo model.Cargo) *CargoService {
	return &CargoService{cargo}
}

func (c *CargoService) InputItem(item model.Item) bool {
	err := c.isCargoReachMax(item)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if item.Type == "L" {
		item.Volume.Height = 1
		item.Volume.Length = 1
		item.Volume.Width = 1
		err = c.tempIsOk(item)
		if err != nil {
			fmt.Println(err)
			return false
		}
	} else {
		err = c.isPassTheDoor(item)
		if err != nil {
			fmt.Println(err)
			return false
		}
	}
	c.Cargo.RegisteredItems = append(c.Cargo.RegisteredItems, item)
	c.Cargo.CurrentVolume = c.Cargo.CurrentVolume + item.Volume.GetVolume()
	c.Cargo.CurrentWeigth = c.Cargo.CurrentVolume + item.Weight
	return true
}

func (c *CargoService) isCargoReachMax(item model.Item) error {
	if c.Cargo.CurrentWeigth+item.Weight > c.Cargo.MaxWeigth {
		return errors.New("cargo has react Max weigth")
	}

	if c.Cargo.CurrentVolume+item.GetVolume() > c.Cargo.Volume.GetVolume() {
		return errors.New("cargo has react Max Volume")
	}
	return nil
}

func (c *CargoService) isPassTheDoor(item model.Item) error {
	if c.Cargo.Door.Height >= item.Volume.Height && c.Cargo.Door.Width >= item.Volume.Width {
		return nil
	}
	if c.Cargo.Door.Height >= item.Volume.Width && c.Cargo.Door.Width >= item.Volume.Height {
		return nil
	}

	return errors.New("cannot pass the door")
}

func (c *CargoService) tempIsOk(item model.Item) error {
	if item.Temperature < c.Cargo.Temperature.MinTemp || item.Temperature > c.Cargo.Temperature.MaxTemp {
		return errors.New("temperature not in range")
	}
	return nil

}
