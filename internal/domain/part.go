package domain

type Part struct {
	id   uint
	name string
}

func (p *Part) Id() uint {
	return p.id
}

func (p *Part) Name() string {
	return p.name
}
