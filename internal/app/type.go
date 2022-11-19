package app

type (
	Company struct {
		Id       uint
		Name     string
		Postcode uint
		State    string
		City     string
		Street   string
		House    string
		Office   string
		//SenderGrainCards   []GrainCard
		//CustomerGrainCards []GrainCard
	}

	Part struct {
		Id   uint
		Name string
	}

	ImpurityParameter struct {
		Id   uint
		Name string
	}

	PartOfImpurityParameter struct {
		Id                uint
		Value             int
		ImpurityParameter ImpurityParameter
		Part              Part
	}

	GrainType struct {
		Id       uint
		Name     string
		GostName string
		Sort     string
	}

	Consignment struct {
		Id                      uint
		GrainCard               GrainCard
		TransportModel          string
		TransportNumber         string
		FirstNameDriver         string
		SecondNameDriver        string
		MiddleNameDriver        string
		TrailerNumber           *string
		GrainType               *GrainType
		GrainImpurity           []PartOfImpurityParameter
		VolatilesImpurity       []PartOfImpurityParameter
		Moisture                *float64
		Origin                  *string
		Nature                  *int
		Color                   *string
		Smell                   *string
		SmallGrainsPercent      *float64
		Vitreous                *float64
		Gluten                  *float64
		GostCulture             *float64
		Filminess               *float64
		Contamitation           *float64
		TypeSubtypeComposition  *float64
		Core                    *float64
		Bug                     *float64
		AdditionalNotes         *string
		AcceptedGrossWeight     *float64
		AcceptedTransportWeight *float64
		SentGrossWeight         float64
		SentTransportWeight     float64
	}

	GrainCard struct {
		Id          uint
		Customer    Company
		Sender      Company
		Storage     uint
		ProcessType string
	}
)
