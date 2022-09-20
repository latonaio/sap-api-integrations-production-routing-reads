package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
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
			ToMaterialAssignment          struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_MatlAssgmt"`
			ToSequence struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Sequence"`
		} `json:"results"`
	} `json:"d"`
}
