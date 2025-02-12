package main

import "fmt"

type Magazine struct {
	Publication
}

func (m Magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", m.name)
}

func createMagazine(name string, pages int, publisher string) iPublication {
	return &Magazine{
		Publication: Publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
