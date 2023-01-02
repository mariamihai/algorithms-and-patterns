package factorymethod

import "fmt"

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAK47(), nil
	}

	if gunType == "musket" {
		return NewMusket(), nil
	}

	return nil, fmt.Errorf("invalid gun type")
}
