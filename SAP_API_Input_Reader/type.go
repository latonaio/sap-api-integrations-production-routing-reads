package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema     string `json:"api_schema"`
	MaterialCode  string `json:"material_code"`
	Plant         string `json:"plant/supplier"`
	Stock         string `json:"stock"`
	DocumentType  string `json:"document_type"`
	DocumentNo    string `json:"document_no"`
	PlannedDate   string `json:"planned_date"`
	ValidatedDate string `json:"validated_date"`
	Deleted       bool   `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string `json:"connection_key"`
	Result            bool   `json:"result"`
	RedisKey          string `json:"redis_key"`
	Filepath          string `json:"filepath"`
	ProductionRouting struct {
		ProductionRoutingGroup        string      `json:"ProductionRoutingGroup"`
		ProductionRouting             string      `json:"ProductionRouting"`
		ProductionRoutingInternalVers string      `json:"ProductionRoutingInternalVers"`
		IsMarkedForDeletion           bool        `json:"IsMarkedForDeletion"`
		BillOfOperationsDesc          string      `json:"BillOfOperationsDesc"`
		Plant                         string      `json:"Plant"`
		BillOfOperationsUsage         string      `json:"BillOfOperationsUsage"`
		BillOfOperationsStatus        string      `json:"BillOfOperationsStatus"`
		ResponsiblePlannerGroup       string      `json:"ResponsiblePlannerGroup"`
		MinimumLotSizeQuantity        string      `json:"MinimumLotSizeQuantity"`
		MaximumLotSizeQuantity        string      `json:"MaximumLotSizeQuantity"`
		BillOfOperationsUnit          string      `json:"BillOfOperationsUnit"`
		CreationDate                  string      `json:"CreationDate"`
		CreatedByUser                 string      `json:"CreatedByUser"`
		LastChangeDate                string      `json:"LastChangeDate"`
		ValidityStartDate             string      `json:"ValidityStartDate"`
		ValidityEndDate               string      `json:"ValidityEndDate"`
		ChangeNumber                  string      `json:"ChangeNumber"`
		PlainLongText                 string      `json:"PlainLongText"`
		MaterialAssignment            struct {
			Product                        string `json:"Product"`
			Plant                          string `json:"Plant"`
			ProductionRoutingMatlAssgmt    string `json:"ProductionRoutingMatlAssgmt"`
			ProductionRtgMatlAssgmtIntVers string `json:"ProductionRtgMatlAssgmtIntVers"`
			CreationDate                   string `json:"CreationDate"`
			LastChangeDate                 string `json:"LastChangeDate"`
			ValidityStartDate              string `json:"ValidityStartDate"`
			ValidityEndDate                string `json:"ValidityEndDate"`
			ChangeNumber                   string `json:"ChangeNumber"`
		} `json:"MaterialAssignment"`
		Sequence struct {
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
			Operation                    struct {
				ProductionRoutingGroup        string      `json:"ProductionRoutingGroup"`
				ProductionRouting             string      `json:"ProductionRouting"`
				ProductionRoutingSequence     string      `json:"ProductionRoutingSequence"`
				ProductionRoutingOpIntID      string      `json:"ProductionRoutingOpIntID"`
				ProductionRoutingOpIntVersion string      `json:"ProductionRoutingOpIntVersion"`
				Operation                     string      `json:"Operation"`
				CreationDate                  string      `json:"CreationDate"`
				LastChangeDate                string      `json:"LastChangeDate"`
				ChangeNumber                  string      `json:"ChangeNumber"`
				ValidityStartDate             string      `json:"ValidityStartDate"`
				ValidityEndDate               string      `json:"ValidityEndDate"`
				OperationText                 string      `json:"OperationText"`
				LongTextLanguageCode          string      `json:"LongTextLanguageCode"`
				Plant                         string      `json:"Plant"`
				OperationControlProfile       string      `json:"OperationControlProfile"`
				OperationStandardTextCode     string      `json:"OperationStandardTextCode"`
				WorkCenterTypeCode            string      `json:"WorkCenterTypeCode"`
				WorkCenterInternalID          string      `json:"WorkCenterInternalID"`
				CapacityCategoryCode          string      `json:"CapacityCategoryCode"`
				OperationCostingRelevancyType string      `json:"OperationCostingRelevancyType"`
				NumberOfTimeTickets           string      `json:"NumberOfTimeTickets"`
				NumberOfConfirmationSlips     string      `json:"NumberOfConfirmationSlips"`
				OperationSetupType            string      `json:"OperationSetupType"`
				OperationSetupGroupCategory   string      `json:"OperationSetupGroupCategory"`
				OperationSetupGroup           string      `json:"OperationSetupGroup"`
				OperationReferenceQuantity    string      `json:"OperationReferenceQuantity"`
				OperationUnit                 string      `json:"OperationUnit"`
				OpQtyToBaseQtyNmrtr           string      `json:"OpQtyToBaseQtyNmrtr"`
				OpQtyToBaseQtyDnmntr          string      `json:"OpQtyToBaseQtyDnmntr"`
				MaximumWaitDuration           string      `json:"MaximumWaitDuration"`
				MaximumWaitDurationUnit       string      `json:"MaximumWaitDurationUnit"`
				MinimumWaitDuration           string      `json:"MinimumWaitDuration"`
				MinimumWaitDurationUnit       string      `json:"MinimumWaitDurationUnit"`
				StandardQueueDuration         string      `json:"StandardQueueDuration"`
				StandardQueueDurationUnit     string      `json:"StandardQueueDurationUnit"`
				MinimumQueueDuration          string      `json:"MinimumQueueDuration"`
				MinimumQueueDurationUnit      string      `json:"MinimumQueueDurationUnit"`
				StandardMoveDuration          string      `json:"StandardMoveDuration"`
				StandardMoveDurationUnit      string      `json:"StandardMoveDurationUnit"`
				MinimumMoveDuration           string      `json:"MinimumMoveDuration"`
				MinimumMoveDurationUnit       string      `json:"MinimumMoveDurationUnit"`
				OpIsExtlyProcdWithSubcontrg   bool        `json:"OpIsExtlyProcdWithSubcontrg"`
				PurchasingInfoRecord          string      `json:"PurchasingInfoRecord"`
				PurchasingOrganization        string      `json:"PurchasingOrganization"`
				PlannedDeliveryDuration       string      `json:"PlannedDeliveryDuration"`
				MaterialGroup                 string      `json:"MaterialGroup"`
				PurchasingGroup               string      `json:"PurchasingGroup"`
				Supplier                      string      `json:"Supplier"`
				NumberOfOperationPriceUnits   string      `json:"NumberOfOperationPriceUnits"`
				CostElement                   string      `json:"CostElement"`
				OpExternalProcessingPrice     string      `json:"OpExternalProcessingPrice"`
				OpExternalProcessingCurrency  string      `json:"OpExternalProcessingCurrency"`
				OperationScrapPercent         string      `json:"OperationScrapPercent"`
				ChangedDateTime               string      `json:"ChangedDateTime"`
				PlainLongText                 string      `json:"PlainLongText"`
				ComponentAllocation           struct {
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
				} `json:"ComponentAllocation"`
			} `json:"Operation"`
		} `json:"Sequence"`
	} `json:"ProductionRouting"`
	APISchema    string   `json:"api_schema"`
	Accepter     []string `json:"accepter"`
	MaterialCode string   `json:"material_code"`
	Plant        string   `json:"plant"`
	Deleted      bool     `json:"deleted"`
}
