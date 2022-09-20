package sap_api_output_formatter

type ProductionRouting struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant"`
	Deleted       string `json:"deleted"`
}

type Header struct {
	ProductionRoutingGroup        string `json:"ProductionRoutingGroup"`
	ProductionRouting             string `json:"ProductionRouting"`
	ProductionRoutingInternalVers string `json:"ProductionRoutingInternalVers"`
	IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
	BillOfOperationsDesc          string `json:"BillOfOperationsDesc"`
	Plant                         string `json:"Plant"`
	BillOfOperationsUsage         string `json:"BillOfOperationsUsage"`
	BillOfOperationsStatus        string `json:"BillOfOperationsStatus"`
	ResponsiblePlannerGroup       string `json:"ResponsiblePlannerGroup"`
	MinimumLotSizeQuantity        string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity        string `json:"MaximumLotSizeQuantity"`
	BillOfOperationsUnit          string `json:"BillOfOperationsUnit"`
	CreationDate                  string `json:"CreationDate"`
	CreatedByUser                 string `json:"CreatedByUser"`
	LastChangeDate                string `json:"LastChangeDate"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	PlainLongText                 string `json:"PlainLongText"`
	ToMaterialAssignment          string `json:"to_MatlAssgmt"`
	ToSequence                    string `json:"to_Sequence"`
}

type MaterialAssignment struct {
	Product                        string `json:"Product"`
	Plant                          string `json:"Plant"`
	ProductionRoutingGroup         string `json:"ProductionRoutingGroup"`
	ProductionRouting              string `json:"ProductionRouting"`
	ProductionRoutingMatlAssgmt    string `json:"ProductionRoutingMatlAssgmt"`
	ProductionRtgMatlAssgmtIntVers string `json:"ProductionRtgMatlAssgmtIntVers"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ChangeNumber                   string `json:"ChangeNumber"`
}

type Sequence struct {
	ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
	ProductionRouting            string `json:"ProductionRouting"`
	ProductionRoutingSequence    string `json:"ProductionRoutingSequence"`
	ProductionRoutingSqncIntVers string `json:"ProductionRoutingSqncIntVers"`
	ChangeNumber                 string `json:"ChangeNumber"`
	ValidityStartDate            string `json:"ValidityStartDate"`
	ValidityEndDate              string `json:"ValidityEndDate"`
	SequenceCategory             string `json:"SequenceCategory"`
	BillOfOperationsRefSequence  string `json:"BillOfOperationsRefSequence"`
	MinimumLotSizeQuantity       string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity       string `json:"MaximumLotSizeQuantity"`
	BillOfOperationsUnit         string `json:"BillOfOperationsUnit"`
	SequenceText                 string `json:"SequenceText"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	ToOperation                  string `json:"to_Operation"`
}

type Operation struct {
	ProductionRoutingGroup        string `json:"ProductionRoutingGroup"`
	ProductionRouting             string `json:"ProductionRouting"`
	ProductionRoutingSequence     string `json:"ProductionRoutingSequence"`
	ProductionRoutingOpIntID      string `json:"ProductionRoutingOpIntID"`
	ProductionRoutingOpIntVersion string `json:"ProductionRoutingOpIntVersion"`
	Operation                     string `json:"Operation"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	OperationText                 string `json:"OperationText"`
	LongTextLanguageCode          string `json:"LongTextLanguageCode"`
	Plant                         string `json:"Plant"`
	OperationControlProfile       string `json:"OperationControlProfile"`
	OperationStandardTextCode     string `json:"OperationStandardTextCode"`
	WorkCenterTypeCode            string `json:"WorkCenterTypeCode"`
	WorkCenterInternalID          string `json:"WorkCenterInternalID"`
	CapacityCategoryCode          string `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType string `json:"OperationCostingRelevancyType"`
	NumberOfTimeTickets           string `json:"NumberOfTimeTickets"`
	NumberOfConfirmationSlips     string `json:"NumberOfConfirmationSlips"`
	OperationSetupType            string `json:"OperationSetupType"`
	OperationSetupGroupCategory   string `json:"OperationSetupGroupCategory"`
	OperationSetupGroup           string `json:"OperationSetupGroup"`
	OperationReferenceQuantity    string `json:"OperationReferenceQuantity"`
	OperationUnit                 string `json:"OperationUnit"`
	OpQtyToBaseQtyNmrtr           string `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr          string `json:"OpQtyToBaseQtyDnmntr"`
	MaximumWaitDuration           string `json:"MaximumWaitDuration"`
	MaximumWaitDurationUnit       string `json:"MaximumWaitDurationUnit"`
	MinimumWaitDuration           string `json:"MinimumWaitDuration"`
	MinimumWaitDurationUnit       string `json:"MinimumWaitDurationUnit"`
	StandardQueueDuration         string `json:"StandardQueueDuration"`
	StandardQueueDurationUnit     string `json:"StandardQueueDurationUnit"`
	MinimumQueueDuration          string `json:"MinimumQueueDuration"`
	MinimumQueueDurationUnit      string `json:"MinimumQueueDurationUnit"`
	StandardMoveDuration          string `json:"StandardMoveDuration"`
	StandardMoveDurationUnit      string `json:"StandardMoveDurationUnit"`
	MinimumMoveDuration           string `json:"MinimumMoveDuration"`
	MinimumMoveDurationUnit       string `json:"MinimumMoveDurationUnit"`
	OpIsExtlyProcdWithSubcontrg   bool   `json:"OpIsExtlyProcdWithSubcontrg"`
	PurchasingInfoRecord          string `json:"PurchasingInfoRecord"`
	PurchasingOrganization        string `json:"PurchasingOrganization"`
	PlannedDeliveryDuration       string `json:"PlannedDeliveryDuration"`
	MaterialGroup                 string `json:"MaterialGroup"`
	PurchasingGroup               string `json:"PurchasingGroup"`
	Supplier                      string `json:"Supplier"`
	NumberOfOperationPriceUnits   string `json:"NumberOfOperationPriceUnits"`
	CostElement                   string `json:"CostElement"`
	OpExternalProcessingPrice     string `json:"OpExternalProcessingPrice"`
	OpExternalProcessingCurrency  string `json:"OpExternalProcessingCurrency"`
	OperationScrapPercent         string `json:"OperationScrapPercent"`
	ChangedDateTime               string `json:"ChangedDateTime"`
	PlainLongText                 string `json:"PlainLongText"`
	ToComponentAllocation         string `json:"to_OpCompAlloc"`
}

type ToMaterialAssignment struct {
	Product                        string `json:"Product"`
	Plant                          string `json:"Plant"`
	ProductionRoutingGroup         string `json:"ProductionRoutingGroup"`
	ProductionRouting              string `json:"ProductionRouting"`
	ProductionRoutingMatlAssgmt    string `json:"ProductionRoutingMatlAssgmt"`
	ProductionRtgMatlAssgmtIntVers string `json:"ProductionRtgMatlAssgmtIntVers"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	ValidityStartDate              string `json:"ValidityStartDate"`
	ValidityEndDate                string `json:"ValidityEndDate"`
	ChangeNumber                   string `json:"ChangeNumber"`
}

type ToSequence struct {
	ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
	ProductionRouting            string `json:"ProductionRouting"`
	ProductionRoutingSequence    string `json:"ProductionRoutingSequence"`
	ProductionRoutingSqncIntVers string `json:"ProductionRoutingSqncIntVers"`
	ChangeNumber                 string `json:"ChangeNumber"`
	ValidityStartDate            string `json:"ValidityStartDate"`
	ValidityEndDate              string `json:"ValidityEndDate"`
	SequenceCategory             string `json:"SequenceCategory"`
	BillOfOperationsRefSequence  string `json:"BillOfOperationsRefSequence"`
	MinimumLotSizeQuantity       string `json:"MinimumLotSizeQuantity"`
	MaximumLotSizeQuantity       string `json:"MaximumLotSizeQuantity"`
	BillOfOperationsUnit         string `json:"BillOfOperationsUnit"`
	SequenceText                 string `json:"SequenceText"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	ToOperation                  string `json:"to_Operation"`
}

type ToOperation struct {
	ProductionRoutingGroup        string `json:"ProductionRoutingGroup"`
	ProductionRouting             string `json:"ProductionRouting"`
	ProductionRoutingSequence     string `json:"ProductionRoutingSequence"`
	ProductionRoutingOpIntID      string `json:"ProductionRoutingOpIntID"`
	ProductionRoutingOpIntVersion string `json:"ProductionRoutingOpIntVersion"`
	Operation                     string `json:"Operation"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	ChangeNumber                  string `json:"ChangeNumber"`
	ValidityStartDate             string `json:"ValidityStartDate"`
	ValidityEndDate               string `json:"ValidityEndDate"`
	OperationText                 string `json:"OperationText"`
	LongTextLanguageCode          string `json:"LongTextLanguageCode"`
	Plant                         string `json:"Plant"`
	OperationControlProfile       string `json:"OperationControlProfile"`
	OperationStandardTextCode     string `json:"OperationStandardTextCode"`
	WorkCenterTypeCode            string `json:"WorkCenterTypeCode"`
	WorkCenterInternalID          string `json:"WorkCenterInternalID"`
	CapacityCategoryCode          string `json:"CapacityCategoryCode"`
	OperationCostingRelevancyType string `json:"OperationCostingRelevancyType"`
	NumberOfTimeTickets           string `json:"NumberOfTimeTickets"`
	NumberOfConfirmationSlips     string `json:"NumberOfConfirmationSlips"`
	OperationSetupType            string `json:"OperationSetupType"`
	OperationSetupGroupCategory   string `json:"OperationSetupGroupCategory"`
	OperationSetupGroup           string `json:"OperationSetupGroup"`
	OperationReferenceQuantity    string `json:"OperationReferenceQuantity"`
	OperationUnit                 string `json:"OperationUnit"`
	OpQtyToBaseQtyNmrtr           string `json:"OpQtyToBaseQtyNmrtr"`
	OpQtyToBaseQtyDnmntr          string `json:"OpQtyToBaseQtyDnmntr"`
	MaximumWaitDuration           string `json:"MaximumWaitDuration"`
	MaximumWaitDurationUnit       string `json:"MaximumWaitDurationUnit"`
	MinimumWaitDuration           string `json:"MinimumWaitDuration"`
	MinimumWaitDurationUnit       string `json:"MinimumWaitDurationUnit"`
	StandardQueueDuration         string `json:"StandardQueueDuration"`
	StandardQueueDurationUnit     string `json:"StandardQueueDurationUnit"`
	MinimumQueueDuration          string `json:"MinimumQueueDuration"`
	MinimumQueueDurationUnit      string `json:"MinimumQueueDurationUnit"`
	StandardMoveDuration          string `json:"StandardMoveDuration"`
	StandardMoveDurationUnit      string `json:"StandardMoveDurationUnit"`
	MinimumMoveDuration           string `json:"MinimumMoveDuration"`
	MinimumMoveDurationUnit       string `json:"MinimumMoveDurationUnit"`
	OpIsExtlyProcdWithSubcontrg   bool   `json:"OpIsExtlyProcdWithSubcontrg"`
	PurchasingInfoRecord          string `json:"PurchasingInfoRecord"`
	PurchasingOrganization        string `json:"PurchasingOrganization"`
	PlannedDeliveryDuration       string `json:"PlannedDeliveryDuration"`
	MaterialGroup                 string `json:"MaterialGroup"`
	PurchasingGroup               string `json:"PurchasingGroup"`
	Supplier                      string `json:"Supplier"`
	NumberOfOperationPriceUnits   string `json:"NumberOfOperationPriceUnits"`
	CostElement                   string `json:"CostElement"`
	OpExternalProcessingPrice     string `json:"OpExternalProcessingPrice"`
	OpExternalProcessingCurrency  string `json:"OpExternalProcessingCurrency"`
	OperationScrapPercent         string `json:"OperationScrapPercent"`
	ChangedDateTime               string `json:"ChangedDateTime"`
	PlainLongText                 string `json:"PlainLongText"`
	ToComponentAllocation         string `json:"to_OpCompAlloc"`
}

type ToComponentAllocation struct {
	ProductionRoutingGroup       string `json:"ProductionRoutingGroup"`
	ProductionRouting            string `json:"ProductionRouting"`
	ProductionRoutingSequence    string `json:"ProductionRoutingSequence"`
	ProductionRoutingOpIntID     string `json:"ProductionRoutingOpIntID"`
	ProdnRtgOpBOMItemInternalID  string `json:"ProdnRtgOpBOMItemInternalID"`
	ProdnRtgOpBOMItemIntVersion  string `json:"ProdnRtgOpBOMItemIntVersion"`
	BillOfMaterialCategory       string `json:"BillOfMaterialCategory"`
	BillOfMaterial               string `json:"BillOfMaterial"`
	BillOfMaterialVariant        string `json:"BillOfMaterialVariant"`
	BillOfMaterialItemNodeNumber string `json:"BillOfMaterialItemNodeNumber"`
	MatlCompIsMarkedForBackflush bool   `json:"MatlCompIsMarkedForBackflush"`
	CreationDate                 string `json:"CreationDate"`
	LastChangeDate               string `json:"LastChangeDate"`
	ValidityStartDate            string `json:"ValidityStartDate"`
	ValidityEndDate              string `json:"ValidityEndDate"`
	ChangeNumber                 string `json:"ChangeNumber"`
	ChangedDateTime              string `json:"ChangedDateTime"`
}
