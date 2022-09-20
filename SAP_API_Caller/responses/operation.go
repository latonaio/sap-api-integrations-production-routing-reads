package responses

type Operation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ProductionRoutingGroup         string      `json:"ProductionRoutingGroup"`
			ProductionRouting              string      `json:"ProductionRouting"`
			ProductionRoutingSequence      string      `json:"ProductionRoutingSequence"`
			ProductionRoutingOpIntID       string      `json:"ProductionRoutingOpIntID"`
			ProductionRoutingOpIntVersion  string      `json:"ProductionRoutingOpIntVersion"`
			Operation                      string      `json:"Operation"`
			CreationDate                   string      `json:"CreationDate"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			ChangeNumber                   string      `json:"ChangeNumber"`
			ValidityStartDate              string      `json:"ValidityStartDate"`
			ValidityEndDate                string      `json:"ValidityEndDate"`
			OperationText                  string      `json:"OperationText"`
			LongTextLanguageCode           string      `json:"LongTextLanguageCode"`
			Plant                          string      `json:"Plant"`
			OperationControlProfile        string      `json:"OperationControlProfile"`
			OperationStandardTextCode      string      `json:"OperationStandardTextCode"`
			WorkCenterTypeCode             string      `json:"WorkCenterTypeCode"`
			WorkCenterInternalID           string      `json:"WorkCenterInternalID"`
			CapacityCategoryCode           string      `json:"CapacityCategoryCode"`
			OperationCostingRelevancyType  string      `json:"OperationCostingRelevancyType"`
			NumberOfTimeTickets            string      `json:"NumberOfTimeTickets"`
			NumberOfConfirmationSlips      string      `json:"NumberOfConfirmationSlips"`
			OperationSetupType             string      `json:"OperationSetupType"`
			OperationSetupGroupCategory    string      `json:"OperationSetupGroupCategory"`
			OperationSetupGroup            string      `json:"OperationSetupGroup"`
			OperationReferenceQuantity     string      `json:"OperationReferenceQuantity"`
			OperationUnit                  string      `json:"OperationUnit"`
			OpQtyToBaseQtyNmrtr            string      `json:"OpQtyToBaseQtyNmrtr"`
			OpQtyToBaseQtyDnmntr           string      `json:"OpQtyToBaseQtyDnmntr"`
			MaximumWaitDuration            string      `json:"MaximumWaitDuration"`
			MaximumWaitDurationUnit        string      `json:"MaximumWaitDurationUnit"`
			MinimumWaitDuration            string      `json:"MinimumWaitDuration"`
			MinimumWaitDurationUnit        string      `json:"MinimumWaitDurationUnit"`
			StandardQueueDuration          string      `json:"StandardQueueDuration"`
			StandardQueueDurationUnit      string      `json:"StandardQueueDurationUnit"`
			MinimumQueueDuration           string      `json:"MinimumQueueDuration"`
			MinimumQueueDurationUnit       string      `json:"MinimumQueueDurationUnit"`
			StandardMoveDuration           string      `json:"StandardMoveDuration"`
			StandardMoveDurationUnit       string      `json:"StandardMoveDurationUnit"`
			MinimumMoveDuration            string      `json:"MinimumMoveDuration"`
			MinimumMoveDurationUnit        string      `json:"MinimumMoveDurationUnit"`
			OpIsExtlyProcdWithSubcontrg    bool        `json:"OpIsExtlyProcdWithSubcontrg"`
			PurchasingInfoRecord           string      `json:"PurchasingInfoRecord"`
			PurchasingOrganization         string      `json:"PurchasingOrganization"`
			PlannedDeliveryDuration        string      `json:"PlannedDeliveryDuration"`
			MaterialGroup                  string      `json:"MaterialGroup"`
			PurchasingGroup                string      `json:"PurchasingGroup"`
			Supplier                       string      `json:"Supplier"`
			NumberOfOperationPriceUnits    string      `json:"NumberOfOperationPriceUnits"`
			CostElement                    string      `json:"CostElement"`
			OpExternalProcessingPrice      string      `json:"OpExternalProcessingPrice"`
			OpExternalProcessingCurrency   string      `json:"OpExternalProcessingCurrency"`
			OperationScrapPercent          string      `json:"OperationScrapPercent"`
			ChangedDateTime                string      `json:"ChangedDateTime"`
			PlainLongText                  string      `json:"PlainLongText"`
			ToComponentAllocation          struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_OpCompAlloc"`
		} `json:"results"`
	} `json:"d"`
}
