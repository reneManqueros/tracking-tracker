package cainiao

type trackingResponse struct {
	Data []struct {
		AllowRetry         bool          `json:"allowRetry"`
		BizType            string        `json:"bizType"`
		CachedTime         string        `json:"cachedTime"`
		DestCountry        string        `json:"destCountry"`
		DestCpList         []interface{} `json:"destCpList"`
		HasRefreshBtn      bool          `json:"hasRefreshBtn"`
		LatestTrackingInfo struct {
			Desc     string `json:"desc"`
			Status   string `json:"status"`
			Time     string `json:"time"`
			TimeZone string `json:"timeZone"`
		} `json:"latestTrackingInfo"`
		MailNo        string        `json:"mailNo"`
		OriginCountry string        `json:"originCountry"`
		OriginCpList  []interface{} `json:"originCpList"`
		Section1      struct {
			CountryName string        `json:"countryName"`
			DetailList  []interface{} `json:"detailList"`
		} `json:"section1"`
		Section2 struct {
			CompanyName  string `json:"companyName"`
			CompanyPhone string `json:"companyPhone"`
			CountryName  string `json:"countryName"`
			DetailList   []struct {
				Desc     string `json:"desc"`
				Status   string `json:"status"`
				Time     string `json:"time"`
				TimeZone string `json:"timeZone"`
			} `json:"detailList"`
			URL string `json:"url"`
		} `json:"section2"`
		ShippingTime     float64 `json:"shippingTime"`
		ShowEstimateTime bool    `json:"showEstimateTime"`
		Status           string  `json:"status"`
		StatusDesc       string  `json:"statusDesc"`
		Success          bool    `json:"success"`
		SyncQuery        bool    `json:"syncQuery"`
	} `json:"data"`
	Success     bool    `json:"success"`
	TimeSeconds float64 `json:"timeSeconds"`
}
