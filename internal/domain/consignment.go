package domain

import "grain-acceptance/internal/app"

type Consignment struct {
	id                      uint
	grainCard               GrainCard
	transportModel          string
	transportNumber         string
	firstNameDriver         string
	secondNameDriver        string
	middleNameDriver        string
	trailerNumber           *string
	grainType               *GrainType
	grainImpurity           []PartOfImpurityParameter
	volatilesImpurity       []PartOfImpurityParameter
	moisture                *float64
	origin                  *string
	nature                  *int
	fallNumber              *float64
	color                   *string
	smell                   *string
	smallGrainsPercent      *float64
	vitreous                *float64
	gluten                  *float64
	gostCulture             *float64
	filminess               *float64
	contamitation           *float64
	typeSubtypeComposition  *float64
	core                    *float64
	bug                     *float64
	additionalNotes         *string
	sentGrossWeight         float64
	sentTransportWeight     float64
	acceptedGrossWeight     *float64
	acceptedTransportWeight *float64
}

type AnalyseParams struct {
	GrainType              GrainType
	GrainImpurity          []PartOfImpurityParameter
	VolatilesImpurity      []PartOfImpurityParameter
	Moisture               float64
	Origin                 string
	Nature                 int
	FallNumber             float64
	Color                  string
	Smell                  string
	SmallGrainsPercent     float64
	Vitreous               float64
	Gluten                 float64
	GostCulture            float64
	Filminess              float64
	Contamitation          float64
	TypeSubtypeComposition float64
	Core                   float64
	Bug                    float64
	AdditionalNotes        string
}

func (c *Consignment) Id() uint {
	return c.id
}

func (c *Consignment) GrainCard() GrainCard {
	return c.grainCard
}

func (c *Consignment) GrainType() *GrainType {
	return c.grainType
}

func (c *Consignment) GrainImpurity() []PartOfImpurityParameter {
	return c.grainImpurity
}

func (c *Consignment) VolatilesImpurity() []PartOfImpurityParameter {
	return c.volatilesImpurity
}

func (c *Consignment) TransportModel() string {
	return c.transportModel
}

func (c *Consignment) TransportNumber() string {
	return c.transportNumber
}

func (c *Consignment) FirstNameDriver() string {
	return c.firstNameDriver
}

func (c *Consignment) SecondNameDriver() string {
	return c.secondNameDriver
}

func (c *Consignment) MiddleNameDriver() string {
	return c.middleNameDriver
}

func (c *Consignment) TrailerNumber() *string {
	return c.trailerNumber
}

func (c *Consignment) Moisture() *float64 {
	return c.moisture
}

func (c *Consignment) Origin() *string {
	return c.origin
}

func (c *Consignment) Nature() *int {
	return c.nature
}

func (c *Consignment) FallNumber() *float64 {
	return c.fallNumber
}

func (c *Consignment) Color() *string {
	return c.color
}

func (c *Consignment) Smell() *string {
	return c.smell
}

func (c *Consignment) SmallGrainsPercent() *float64 {
	return c.smallGrainsPercent
}

func (c *Consignment) Vitreous() *float64 {
	return c.vitreous
}

func (c *Consignment) Gluten() *float64 {
	return c.gluten
}

func (c *Consignment) GostCulture() *float64 {
	return c.gostCulture
}

func (c *Consignment) Filminess() *float64 {
	return c.filminess
}

func (c *Consignment) Contamitation() *float64 {
	return c.contamitation
}

func (c *Consignment) TypeSubtypeComposition() *float64 {
	return c.typeSubtypeComposition
}

func (c *Consignment) Core() *float64 {
	return c.core
}

func (c *Consignment) Bug() *float64 {
	return c.bug
}

func (c *Consignment) AdditionalNotes() *string {
	return c.additionalNotes
}

func (c *Consignment) SentGrossWeight() float64 {
	return c.sentGrossWeight
}

func (c *Consignment) SentTransportWeight() float64 {
	return c.sentTransportWeight
}

func (c *Consignment) AcceptedGrossWeight() *float64 {
	return c.acceptedGrossWeight
}

func (c *Consignment) AcceptedTransportWeight() *float64 {
	return c.acceptedTransportWeight
}

func (c *Consignment) UpdateSentGrossWeight(grossWeight float64) {
	c.acceptedGrossWeight = &grossWeight
}

func (c *Consignment) UpdateSentTransportWeight(transportWeight float64) {
	c.acceptedTransportWeight = &transportWeight
}

func (c *Consignment) UpdateAnalyses(ap AnalyseParams) {
	c.grainType = &ap.GrainType
	c.grainImpurity = ap.GrainImpurity
	c.volatilesImpurity = ap.VolatilesImpurity
	c.moisture = &ap.Moisture
	c.origin = &ap.Origin
	c.nature = &ap.Nature
	c.fallNumber = &ap.FallNumber
	c.color = &ap.Color
	c.smell = &ap.Smell
	c.smallGrainsPercent = &ap.SmallGrainsPercent
	c.vitreous = &ap.Vitreous
	c.gluten = &ap.Gluten
	c.gostCulture = &ap.GostCulture
	c.filminess = &ap.Filminess
	c.contamitation = &ap.Contamitation
	c.typeSubtypeComposition = &ap.TypeSubtypeComposition
	c.core = &ap.Core
	c.bug = &ap.Bug
	c.additionalNotes = &ap.AdditionalNotes
}

func (c *Consignment) InitByApp(consignment app.Consignment) {
	c.id = consignment.Id

	c.transportModel = consignment.TransportModel
	c.transportNumber = consignment.TransportNumber
	c.firstNameDriver = consignment.FirstNameDriver
	c.secondNameDriver = consignment.SecondNameDriver
	c.middleNameDriver = consignment.MiddleNameDriver
	c.trailerNumber = consignment.TrailerNumber

	c.moisture = consignment.Moisture
	c.origin = consignment.Origin
	c.nature = consignment.Nature
	c.color = consignment.Color
	c.smell = consignment.Smell
	c.smallGrainsPercent = consignment.SmallGrainsPercent
	c.vitreous = consignment.Vitreous
	c.gluten = consignment.Gluten
	c.gostCulture = consignment.GostCulture
	c.filminess = consignment.Filminess
	c.contamitation = consignment.Contamitation
	c.typeSubtypeComposition = consignment.TypeSubtypeComposition
	c.core = consignment.Core
	c.bug = consignment.Bug
	c.additionalNotes = consignment.AdditionalNotes
	c.acceptedTransportWeight = consignment.AcceptedTransportWeight
	c.acceptedGrossWeight = consignment.AcceptedGrossWeight
	c.sentTransportWeight = consignment.SentTransportWeight
	c.sentGrossWeight = consignment.SentGrossWeight
}
