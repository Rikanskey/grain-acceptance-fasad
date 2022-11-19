package domain

type PartOfImpurityParameter struct {
	id                uint
	value             int
	part              Part
	impurityParameter ImpurityParameter
}

func (poip *PartOfImpurityParameter) Id() uint {
	return poip.id
}

func (poip *PartOfImpurityParameter) Value() int {
	return poip.value
}

func (poip *PartOfImpurityParameter) Part() Part {
	return poip.part
}

func (poip *PartOfImpurityParameter) ImpurityParameter() ImpurityParameter {
	return poip.impurityParameter
}
