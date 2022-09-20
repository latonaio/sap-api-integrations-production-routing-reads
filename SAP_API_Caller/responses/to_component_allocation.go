package responses

type ToComponentAllocation struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			ProductionRoutingGroup       string      `json:"ProductionRoutingGroup"`
			ProductionRouting            string      `json:"ProductionRouting"`
			ProductionRoutingSequence    string      `json:"ProductionRoutingSequence"`
			ProductionRoutingOpIntID     string      `json:"ProductionRoutingOpIntID"`
			ProdnRtgOpBOMItemInternalID  string      `json:"ProdnRtgOpBOMItemInternalID"`
			ProdnRtgOpBOMItemIntVersion  string      `json:"ProdnRtgOpBOMItemIntVersion"`
			BillOfMaterialCategory       string      `json:"BillOfMaterialCategory"`
			BillOfMaterial               string      `json:"BillOfMaterial"`
			BillOfMaterialVariant        string      `json:"BillOfMaterialVariant"`
			BillOfMaterialItemNodeNumber string      `json:"BillOfMaterialItemNodeNumber"`
			MatlCompIsMarkedForBackflush bool        `json:"MatlCompIsMarkedForBackflush"`
			CreationDate                 string      `json:"CreationDate"`
			LastChangeDate               string      `json:"LastChangeDate"`
			ValidityStartDate            string      `json:"ValidityStartDate"`
			ValidityEndDate              string      `json:"ValidityEndDate"`
			ChangeNumber                 string      `json:"ChangeNumber"`
			ChangedDateTime              string      `json:"ChangedDateTime"`
		} `json:"results"`
	} `json:"d"`
}
