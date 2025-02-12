package factory

import "fmt"

type Newspaper struct {
	Publication
}

func (n Newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

func createNewspaper(name string, pages int, publisher string) IPublication {
	return &Newspaper{
		Publication: Publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
