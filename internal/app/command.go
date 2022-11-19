package app

type (
	UpdateConsignmentGrossWeightQuery struct {
		ConsignmentId string
		GrossWeight   float32
	}

	UpdateConsignmentTransportWeightQuery struct {
		ConsignmentId   string
		TransportWeight float64
	}

	UpdateConsignmentAnalyseQuery struct {
		ConsignmentId          string
		AdditionalNotes        string
		Bug                    float64
		Color                  string
		Contamitation          float64
		Core                   float64
		Filminess              float64
		Gluten                 float64
		GostCulture            float64
		Moisture               float64
		Nature                 int
		Origin                 string
		SmallGrainsPercent     float64
		Smell                  string
		TypeSubtypeComposition float64
		Vitreous               float64
	}
)
