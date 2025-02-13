package factory

import "fmt"

func NewPublication(pubType string, name string, pages int, publisher string) (IPublication, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pages, publisher), nil
	}

	if pubType == "magazine" {
		return createMagazine(name, pages, publisher), nil
	}

	return nil, fmt.Errorf("no such publication type")
}
