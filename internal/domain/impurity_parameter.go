package domain

type ImpurityParameter struct {
	id   uint
	name string
}

func (ip *ImpurityParameter) Id() uint {
	return ip.id
}

func (ip *ImpurityParameter) Name() string {
	return ip.name
}
