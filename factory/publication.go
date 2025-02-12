package main

type iPublication interface {
	setName(name string)
	setPages(pages int)
	setPublisher(publisher string)
	getName() string
	getPages() int
	getPublisher() string
}

type Publication struct {
	name      string
	pages     int
	publisher string
}

func (p *Publication) setName(name string) {
	p.name = name
}

func (p *Publication) setPages(pages int) {
	p.pages = pages
}

func (p *Publication) setPublisher(publisher string) {
	p.publisher = publisher
}

func (p *Publication) getName() string {
	return p.name
}

func (p *Publication) getPages() int {
	return p.pages
}

func (p *Publication) getPublisher() string {
	return p.publisher
}
