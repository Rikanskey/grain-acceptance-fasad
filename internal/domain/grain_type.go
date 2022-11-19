package domain

type SortType string

const (
	fodder SortType = "Fodder"
	food   SortType = "Food"
)

func (st SortType) String() string {
	switch st {
	case fodder:
		return "Fodder"
	case food:
		return "Food"
	}

	return ""
}

type GrainType struct {
	id   int
	name string
	//Enum string type: fodder or food
	sort     SortType
	gostName string
}

func (gt *GrainType) Id() int {
	return gt.id
}

func (gt *GrainType) Name() string {
	return gt.name
}

func (gt *GrainType) Sort() SortType {
	return gt.sort
}

func (gt *GrainType) GostName() string {
	return gt.gostName
}
