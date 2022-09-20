package responses

type ToMaterialAssignment struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			Product                        string      `json:"Product"`
			Plant                          string      `json:"Plant"`
			ProductionRoutingGroup         string      `json:"ProductionRoutingGroup"`
			ProductionRouting              string      `json:"ProductionRouting"`
			ProductionRoutingMatlAssgmt    string      `json:"ProductionRoutingMatlAssgmt"`
			ProductionRtgMatlAssgmtIntVers string      `json:"ProductionRtgMatlAssgmtIntVers"`
			CreationDate                   string      `json:"CreationDate"`
			LastChangeDate                 string      `json:"LastChangeDate"`
			ValidityStartDate              string      `json:"ValidityStartDate"`
			ValidityEndDate                string      `json:"ValidityEndDate"`
			ChangeNumber                   string      `json:"ChangeNumber"`
		} `json:"results"`
	} `json:"d"`
}
