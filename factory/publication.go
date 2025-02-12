package factory

type IPublication interface {
	SetName(name string)
	SetPages(pages int)
	SetPublisher(publisher string)
	GetName() string
	GetPages() int
	GetPublisher() string
}

type Publication struct {
	name      string
	pages     int
	publisher string
}

func (p *Publication) SetName(name string) {
	p.name = name
}

func (p *Publication) SetPages(pages int) {
	p.pages = pages
}

func (p *Publication) SetPublisher(publisher string) {
	p.publisher = publisher
}

func (p *Publication) GetName() string {
	return p.name
}

func (p *Publication) GetPages() int {
	return p.pages
}

func (p *Publication) GetPublisher() string {
	return p.publisher
}
