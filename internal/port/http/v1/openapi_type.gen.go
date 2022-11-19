// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package v1

// AnalyseUpdateRequest defines model for AnalyseUpdateRequest.
type AnalyseUpdateRequest struct {
	AdditionalNotes        *string  `json:"additionalNotes,omitempty"`
	Bug                    *float64 `json:"bug,omitempty"`
	Color                  *string  `json:"color,omitempty"`
	Contamitation          *float64 `json:"contamitation,omitempty"`
	Core                   *float64 `json:"core,omitempty"`
	Filminess              *float64 `json:"filminess,omitempty"`
	Gluten                 *float64 `json:"gluten,omitempty"`
	GostCulture            *float64 `json:"gostCulture,omitempty"`
	Moisture               *float64 `json:"moisture,omitempty"`
	Nature                 *int     `json:"nature,omitempty"`
	Origin                 *string  `json:"origin,omitempty"`
	SmallGrainsPercent     *float64 `json:"smallGrainsPercent,omitempty"`
	Smell                  *string  `json:"smell,omitempty"`
	TypeSubtypeComposition *float64 `json:"typeSubtypeComposition,omitempty"`
	Vitreous               *float64 `json:"vitreous,omitempty"`
}

// CompanyRequest defines model for CompanyRequest.
type CompanyRequest struct {
	City     string  `json:"city"`
	House    string  `json:"house"`
	Name     string  `json:"name"`
	Office   *string `json:"office,omitempty"`
	Postcode int     `json:"postcode"`
	State    string  `json:"state"`
	Street   string  `json:"street"`
}

// CompanyResponse defines model for CompanyResponse.
type CompanyResponse struct {
	City     string  `json:"city"`
	House    string  `json:"house"`
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Office   *string `json:"office,omitempty"`
	Postcode int     `json:"postcode"`
	State    string  `json:"state"`
	Street   string  `json:"street"`
}

// ConsignmentGrainCardResponse defines model for ConsignmentGrainCardResponse.
type ConsignmentGrainCardResponse struct {
	Customer    *CompanyResponse `json:"customer,omitempty"`
	Id          *int             `json:"id,omitempty"`
	ProcessType *string          `json:"processType,omitempty"`
	Sender      *CompanyResponse `json:"sender,omitempty"`
}

// ConsignmentResponse defines model for ConsignmentResponse.
type ConsignmentResponse struct {
	AcceptedGrossWeight     *float64                      `json:"acceptedGrossWeight,omitempty"`
	AcceptedTransportWeight *float64                      `json:"acceptedTransportWeight,omitempty"`
	AdditionalNotes         *string                       `json:"additionalNotes,omitempty"`
	Bug                     *float64                      `json:"bug,omitempty"`
	Color                   *string                       `json:"color,omitempty"`
	Contamitation           *float64                      `json:"contamitation,omitempty"`
	Core                    *float64                      `json:"core,omitempty"`
	Filminess               *float64                      `json:"filminess,omitempty"`
	FirstNameDriver         string                        `json:"firstNameDriver"`
	Gluten                  *float64                      `json:"gluten,omitempty"`
	GostCulture             *float64                      `json:"gostCulture,omitempty"`
	GrainCard               *ConsignmentGrainCardResponse `json:"grainCard,omitempty"`
	GrainImpurity           *[]PartOfImpurityResponse     `json:"grainImpurity,omitempty"`
	GrainType                *GrainTypeResponse            `json:"grainTye,omitempty"`
	Id                      int                           `json:"id"`
	MiddleNameDriver        string                        `json:"middleNameDriver"`
	Moisture                *float64                      `json:"moisture,omitempty"`
	Nature                  *int                          `json:"nature,omitempty"`
	Origin                  *string                       `json:"origin,omitempty"`
	SecondNameDriver        string                        `json:"secondNameDriver"`
	SentGrossWeight         float64                       `json:"sentGrossWeight"`
	SentTransportWeight     float64                       `json:"sentTransportWeight"`
	SmallGrainsPercent      *float64                      `json:"smallGrainsPercent,omitempty"`
	Smell                   *string                       `json:"smell,omitempty"`
	TrailerNumber           *string                       `json:"trailerNumber,omitempty"`
	TransportModel          string                        `json:"transportModel"`
	TransportNumber         string                       `json:"transportNumber"`
	TypeSubtypeComposition  *float64                      `json:"typeSubtypeComposition,omitempty"`
	Vitreous                *float64                      `json:"vitreous,omitempty"`
	VolatilesImpurity       *[]PartOfImpurityResponse     `json:"volatilesImpurity,omitempty"`
}

// CreateCardRequest defines model for CreateCardRequest.
type CreateCardRequest struct {
	Customer *CompanyRequest `json:"customer,omitempty"`
	Sender   *CompanyRequest `json:"sender,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Details string `json:"details"`
	Slug    string `json:"slug"`
}

// GrainTypeResponse defines model for GrainTypeResponse.
type GrainTypeResponse struct {
	GostName string `json:"gostName"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Sort     string `json:"sort"`
}

// GrossWeightUpdateRequest defines model for GrossWeightUpdateRequest.
type GrossWeightUpdateRequest struct {
	Value float32 `json:"value"`
}

// ImpurityParameterResponse defines model for ImpurityParameterResponse.
type ImpurityParameterResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// PartOfImpurityResponse defines model for PartOfImpurityResponse.
type PartOfImpurityResponse struct {
	Id                int                       `json:"id,omitempty"`
	ImpurityParameter ImpurityParameterResponse `json:"impurityParameter,omitempty"`
	Part              PartResponse              `json:"part,omitempty"`
	Value             int                       `json:"value,omitempty"`
}

// PartResponse defines model for PartResponse.
type PartResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// TransportWeightUpdateRequest defines model for TransportWeightUpdateRequest.
type TransportWeightUpdateRequest struct {
	Value float64 `json:"value"`
}

// CreateCardJSONBody defines parameters for CreateCard.
type CreateCardJSONBody = CreateCardRequest

// CreateCardJSONRequestBody defines body for CreateCard for application/json ContentType.
type CreateCardJSONRequestBody = CreateCardJSONBody