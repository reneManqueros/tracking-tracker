package fedex

type trackingResponse struct {
	TrackPackagesResponse struct {
		Successful          bool `json:"successful"`
		PassedLoggedInCheck bool `json:"passedLoggedInCheck"`
		ErrorList           []struct {
			Code      string      `json:"code"`
			Message   string      `json:"message"`
			Source    interface{} `json:"source"`
			RootCause interface{} `json:"rootCause"`
		} `json:"errorList"`
		PackageList []struct {
			ShipperAccountNumber         interface{} `json:"shipperAccountNumber"`
			DelayDetail                  interface{} `json:"delayDetail"`
			TrackingNbr                  string      `json:"trackingNbr"`
			TrackingQualifier            string      `json:"trackingQualifier"`
			TrackingCarrierCd            string      `json:"trackingCarrierCd"`
			TrackingCarrierDesc          string      `json:"trackingCarrierDesc"`
			DisplayTrackingNbr           string      `json:"displayTrackingNbr"`
			ShipperCmpnyName             string      `json:"shipperCmpnyName"`
			ShipperName                  string      `json:"shipperName"`
			ShipperAddr1                 string      `json:"shipperAddr1"`
			ShipperAddr2                 string      `json:"shipperAddr2"`
			ShipperAddrLines             interface{} `json:"shipperAddrLines"`
			ShipperCity                  string      `json:"shipperCity"`
			ShipperStateCD               string      `json:"shipperStateCD"`
			ShipperZip                   string      `json:"shipperZip"`
			ShipperCntryCD               string      `json:"shipperCntryCD"`
			ShipperPhoneNbr              string      `json:"shipperPhoneNbr"`
			ShippedBy                    string      `json:"shippedBy"`
			RecipientCmpnyName           string      `json:"recipientCmpnyName"`
			RecipientName                string      `json:"recipientName"`
			RecipientAddr1               string      `json:"recipientAddr1"`
			RecipientAddr2               string      `json:"recipientAddr2"`
			RecipientAddrLines           interface{} `json:"recipientAddrLines"`
			RecipientCity                string      `json:"recipientCity"`
			RecipientStateCD             string      `json:"recipientStateCD"`
			RecipientZip                 string      `json:"recipientZip"`
			RecipientCntryCD             string      `json:"recipientCntryCD"`
			RecipientPhoneNbr            string      `json:"recipientPhoneNbr"`
			ShippedTo                    string      `json:"shippedTo"`
			KeyStatus                    string      `json:"keyStatus"`
			KeyStatusCD                  string      `json:"keyStatusCD"`
			LastScanStatus               string      `json:"lastScanStatus"`
			LastScanDateTime             string      `json:"lastScanDateTime"`
			ReceivedByNm                 string      `json:"receivedByNm"`
			SubStatus                    string      `json:"subStatus"`
			MainStatus                   string      `json:"mainStatus"`
			StatusBarCD                  string      `json:"statusBarCD"`
			ShortStatus                  string      `json:"shortStatus"`
			ShortStatusCD                string      `json:"shortStatusCD"`
			StatusLocationAddr1          string      `json:"statusLocationAddr1"`
			StatusLocationAddr2          string      `json:"statusLocationAddr2"`
			StatusLocationAddrLines      interface{} `json:"statusLocationAddrLines"`
			StatusLocationCity           string      `json:"statusLocationCity"`
			StatusLocationStateCD        string      `json:"statusLocationStateCD"`
			StatusLocationZip            string      `json:"statusLocationZip"`
			StatusLocationCntryCD        string      `json:"statusLocationCntryCD"`
			StatusWithDetails            string      `json:"statusWithDetails"`
			HalType                      string      `json:"halType"`
			HalCmpnyName                 string      `json:"halCmpnyName"`
			IsHALAddress                 bool        `json:"isHALAddress"`
			ShipDt                       string      `json:"shipDt"`
			DisplayShipDt                string      `json:"displayShipDt"`
			DisplayShipTm                string      `json:"displayShipTm"`
			DisplayShipDateTime          string      `json:"displayShipDateTime"`
			PickupDt                     string      `json:"pickupDt"`
			DisplayPickupDt              string      `json:"displayPickupDt"`
			DisplayPickupTm              string      `json:"displayPickupTm"`
			DisplayPickupDateTime        string      `json:"displayPickupDateTime"`
			EstDeliveryDt                string      `json:"estDeliveryDt"`
			EstDeliveryTm                string      `json:"estDeliveryTm"`
			DisplayEstDeliveryDt         string      `json:"displayEstDeliveryDt"`
			DisplayEstDeliveryTm         string      `json:"displayEstDeliveryTm"`
			DisplayEstDeliveryDateTime   string      `json:"displayEstDeliveryDateTime"`
			ActDeliveryDt                string      `json:"actDeliveryDt"`
			DisplayActDeliveryDt         string      `json:"displayActDeliveryDt"`
			DisplayActDeliveryTm         string      `json:"displayActDeliveryTm"`
			DisplayActDeliveryDateTime   string      `json:"displayActDeliveryDateTime"`
			TenderedDt                   string      `json:"tenderedDt"`
			DisplayTenderedDt            string      `json:"displayTenderedDt"`
			DisplayTenderedTm            string      `json:"displayTenderedTm"`
			DisplayTenderedDateTime      string      `json:"displayTenderedDateTime"`
			ApptDeliveryDt               string      `json:"apptDeliveryDt"`
			DisplayApptDeliveryDt        string      `json:"displayApptDeliveryDt"`
			DisplayApptDeliveryTm        string      `json:"displayApptDeliveryTm"`
			DisplayApptDeliveryDateTime  string      `json:"displayApptDeliveryDateTime"`
			AttemptedDelivery            interface{} `json:"attemptedDelivery"`
			AvailableAtStation           interface{} `json:"availableAtStation"`
			CodDetail                    interface{} `json:"codDetail"`
			NickName                     string      `json:"nickName"`
			Note                         string      `json:"note"`
			MatchedAccountList           []string    `json:"matchedAccountList"`
			FxfAdvanceETA                string      `json:"fxfAdvanceETA"`
			FxfAdvanceReason             string      `json:"fxfAdvanceReason"`
			FxfAdvanceStatusCode         string      `json:"fxfAdvanceStatusCode"`
			FxfAdvanceStatusDesc         string      `json:"fxfAdvanceStatusDesc"`
			DestLink                     string      `json:"destLink"`
			OriginLink                   string      `json:"originLink"`
			HasBillOfLadingImage         bool        `json:"hasBillOfLadingImage"`
			HasBillPresentment           bool        `json:"hasBillPresentment"`
			SignatureRequired            int         `json:"signatureRequired"`
			TotalKgsWgt                  string      `json:"totalKgsWgt"`
			DisplayTotalKgsWgt           string      `json:"displayTotalKgsWgt"`
			TotalLbsWgt                  string      `json:"totalLbsWgt"`
			DisplayTotalLbsWgt           string      `json:"displayTotalLbsWgt"`
			DisplayTotalWgt              string      `json:"displayTotalWgt"`
			PkgKgsWgt                    string      `json:"pkgKgsWgt"`
			DisplayPkgKgsWgt             string      `json:"displayPkgKgsWgt"`
			PkgLbsWgt                    string      `json:"pkgLbsWgt"`
			DisplayPkgLbsWgt             string      `json:"displayPkgLbsWgt"`
			DisplayPkgWgt                string      `json:"displayPkgWgt"`
			TotalDIMLbsWgt               string      `json:"totalDIMLbsWgt"`
			DisplayTotalDIMLbsWgt        string      `json:"displayTotalDIMLbsWgt"`
			TotalDIMKgsWgt               string      `json:"totalDIMKgsWgt"`
			DisplayTotalDIMKgsWgt        string      `json:"displayTotalDIMKgsWgt"`
			DisplayTotalDIMWgt           string      `json:"displayTotalDIMWgt"`
			Dimensions                   string      `json:"dimensions"`
			MasterTrackingNbr            string      `json:"masterTrackingNbr"`
			MasterQualifier              string      `json:"masterQualifier"`
			MasterCarrierCD              string      `json:"masterCarrierCD"`
			OriginalOutboundTrackingNbr  interface{} `json:"originalOutboundTrackingNbr"`
			OriginalOutboundQualifier    string      `json:"originalOutboundQualifier"`
			OriginalOutboundCarrierCD    string      `json:"originalOutboundCarrierCD"`
			InvoiceNbrList               []string    `json:"invoiceNbrList"`
			ReferenceList                []string    `json:"referenceList"`
			DoorTagNbrList               []string    `json:"doorTagNbrList"`
			ReferenceDescList            []string    `json:"referenceDescList"`
			PurchaseOrderNbrList         []string    `json:"purchaseOrderNbrList"`
			BillofLadingNbrList          []string    `json:"billofLadingNbrList"`
			ShipperRefList               []string    `json:"shipperRefList"`
			RmaList                      []string    `json:"rmaList"`
			DeptNbrList                  []string    `json:"deptNbrList"`
			ShipmentIDList               []string    `json:"shipmentIdList"`
			TcnList                      []string    `json:"tcnList"`
			PartnerCarrierNbrList        []string    `json:"partnerCarrierNbrList"`
			HasAssociatedShipments       bool        `json:"hasAssociatedShipments"`
			HasAssociatedReturnShipments bool        `json:"hasAssociatedReturnShipments"`
			IsDistributedService         bool        `json:"isDistributedService"`
			AssocShpGrp                  int         `json:"assocShpGrp"`
			DrTgGrp                      []string    `json:"drTgGrp"`
			AssociationInfoList          []struct {
				TrackingNumberInfo struct {
					TrackingNumber       string      `json:"trackingNumber"`
					TrackingQualifier    string      `json:"trackingQualifier"`
					TrackingCarrier      string      `json:"trackingCarrier"`
					ProcessingParameters interface{} `json:"processingParameters"`
				} `json:"trackingNumberInfo"`
				AssociatedType string `json:"associatedType"`
			} `json:"associationInfoList"`
			ReturnReason                string      `json:"returnReason"`
			ReturnRelationship          interface{} `json:"returnRelationship"`
			SkuItemUpcCdList            []string    `json:"skuItemUpcCdList"`
			ReceiveQtyList              []string    `json:"receiveQtyList"`
			ItemDescList                []string    `json:"itemDescList"`
			PartNbrList                 []string    `json:"partNbrList"`
			ServiceCD                   string      `json:"serviceCD"`
			ServiceDesc                 string      `json:"serviceDesc"`
			ServiceShortDesc            string      `json:"serviceShortDesc"`
			PackageType                 string      `json:"packageType"`
			Packaging                   string      `json:"packaging"`
			ClearanceDetailLink         string      `json:"clearanceDetailLink"`
			ShowClearanceDetailLink     bool        `json:"showClearanceDetailLink"`
			ManufactureCountryCDList    []string    `json:"manufactureCountryCDList"`
			CommodityCDList             []string    `json:"commodityCDList"`
			CommodityDescList           []string    `json:"commodityDescList"`
			CerNbrList                  []string    `json:"cerNbrList"`
			CerComplaintCDList          []string    `json:"cerComplaintCDList"`
			CerComplaintDescList        []string    `json:"cerComplaintDescList"`
			CerEventDateList            []string    `json:"cerEventDateList"`
			DisplayCerEventDateList     []string    `json:"displayCerEventDateList"`
			TotalPieces                 string      `json:"totalPieces"`
			SpecialHandlingServicesList []string    `json:"specialHandlingServicesList"`
			ShipmentType                string      `json:"shipmentType"`
			PkgContentDesc1             string      `json:"pkgContentDesc1"`
			PkgContentDesc2             string      `json:"pkgContentDesc2"`
			DocAWBNbr                   string      `json:"docAWBNbr"`
			OriginalCharges             string      `json:"originalCharges"`
			TransportationCD            string      `json:"transportationCD"`
			TransportationDesc          string      `json:"transportationDesc"`
			DutiesAndTaxesCD            string      `json:"dutiesAndTaxesCD"`
			DutiesAndTaxesDesc          string      `json:"dutiesAndTaxesDesc"`
			OrigPieceCount              string      `json:"origPieceCount"`
			DestPieceCount              string      `json:"destPieceCount"`
			BillNoteMsg                 string      `json:"billNoteMsg"`
			GoodsClassificationCD       string      `json:"goodsClassificationCD"`
			ReceipientAddrQty           string      `json:"receipientAddrQty"`
			DeliveryAttempt             string      `json:"deliveryAttempt"`
			CodReturnTrackNbr           string      `json:"codReturnTrackNbr"`
			ReturnMovementStatus        interface{} `json:"returnMovementStatus"`
			ScanEventList               []struct {
				Date                string      `json:"date"`
				Time                string      `json:"time"`
				GmtOffset           string      `json:"gmtOffset"`
				Status              string      `json:"status"`
				StatusCD            string      `json:"statusCD"`
				ScanLocation        string      `json:"scanLocation"`
				ScanDetails         string      `json:"scanDetails"`
				ScanDetailsHTML     string      `json:"scanDetailsHtml"`
				RtrnShprTrkNbr      string      `json:"rtrnShprTrkNbr"`
				StatusExceptionCode string      `json:"statusExceptionCode"`
				DelayDetail         interface{} `json:"delayDetail"`
				IsException         bool        `json:"isException"`
				IsClearanceDelay    bool        `json:"isClearanceDelay"`
				IsDelivered         bool        `json:"isDelivered"`
				IsDelException      bool        `json:"isDelException"`
			} `json:"scanEventList"`
			OriginAddr1             string      `json:"originAddr1"`
			OriginAddr2             string      `json:"originAddr2"`
			OriginAddrLines         interface{} `json:"originAddrLines"`
			OriginCity              string      `json:"originCity"`
			OriginStateCD           string      `json:"originStateCD"`
			OriginZip               string      `json:"originZip"`
			OriginCntryCD           string      `json:"originCntryCD"`
			OriginLocationID        string      `json:"originLocationID"`
			OriginTermCity          string      `json:"originTermCity"`
			OriginTermStateCD       string      `json:"originTermStateCD"`
			DestLocationAddr1       string      `json:"destLocationAddr1"`
			DestLocationAddr2       string      `json:"destLocationAddr2"`
			DestLocationAddrLines   []string    `json:"destLocationAddrLines"`
			DestLocationCity        string      `json:"destLocationCity"`
			DestLocationStateCD     string      `json:"destLocationStateCD"`
			DestLocationZip         string      `json:"destLocationZip"`
			DestLocationCntryCD     string      `json:"destLocationCntryCD"`
			DestLocationID          string      `json:"destLocationID"`
			DestLocationTermCity    string      `json:"destLocationTermCity"`
			DestLocationTermStateCD string      `json:"destLocationTermStateCD"`
			DestAddr1               string      `json:"destAddr1"`
			DestAddr2               string      `json:"destAddr2"`
			DestAddrLines           interface{} `json:"destAddrLines"`
			DestCity                string      `json:"destCity"`
			DestStateCD             string      `json:"destStateCD"`
			DestZip                 string      `json:"destZip"`
			DestCntryCD             string      `json:"destCntryCD"`
			HalAddr1                string      `json:"halAddr1"`
			HalAddr2                string      `json:"halAddr2"`
			HalAddrLines            interface{} `json:"halAddrLines"`
			HalCity                 string      `json:"halCity"`
			HalStateCD              string      `json:"halStateCD"`
			HalZipCD                string      `json:"halZipCD"`
			HalCntryCD              string      `json:"halCntryCD"`
			ActualDelAddrCity       string      `json:"actualDelAddrCity"`
			ActualDelAddrStateCD    string      `json:"actualDelAddrStateCD"`
			ActualDelAddrZipCD      string      `json:"actualDelAddrZipCD"`
			ActualDelAddrCntryCD    string      `json:"actualDelAddrCntryCD"`
			TotalTransitMiles       string      `json:"totalTransitMiles"`
			ExcepReasonList         []string    `json:"excepReasonList"`
			ExcepActionList         []string    `json:"excepActionList"`
			ExceptionReason         string      `json:"exceptionReason"`
			ExceptionAction         string      `json:"exceptionAction"`
			StatusDetailsList       []string    `json:"statusDetailsList"`
			TrackErrCD              string      `json:"trackErrCD"`
			DestTZ                  string      `json:"destTZ"`
			OriginTZ                string      `json:"originTZ"`
			IsMultiStat             string      `json:"isMultiStat"`
			MultiStatList           []struct {
				MultiPiec   string `json:"multiPiec"`
				MultiTm     string `json:"multiTm"`
				MultiDispTm string `json:"multiDispTm"`
				MultiSta    string `json:"multiSta"`
			} `json:"multiStatList"`
			MaskMessage                  string `json:"maskMessage"`
			DeliveryService              string `json:"deliveryService"`
			MilestoDestination           string `json:"milestoDestination"`
			Terms                        string `json:"terms"`
			PayorAcctNbr                 string `json:"payorAcctNbr"`
			MeterNumber                  string `json:"meterNumber"`
			OriginUbanizationCode        string `json:"originUbanizationCode"`
			OriginCountryName            string `json:"originCountryName"`
			IsOriginResidential          bool   `json:"isOriginResidential"`
			HalUrbanizationCD            string `json:"halUrbanizationCD"`
			HalCountryName               string `json:"halCountryName"`
			ActualDelAddrUrbanizationCD  string `json:"actualDelAddrUrbanizationCD"`
			ActualDelAddrCountryName     string `json:"actualDelAddrCountryName"`
			DestUrbanizationCD           string `json:"destUrbanizationCD"`
			DestCountryName              string `json:"destCountryName"`
			DelToDesc                    string `json:"delToDesc"`
			RecpShareID                  string `json:"recpShareID"`
			ShprShareID                  string `json:"shprShareID"`
			RequestedAppointmentInfoList []struct {
				SpclInstructDesc string `json:"spclInstructDesc"`
				DelivOptn        string `json:"delivOptn"`
				DelivOptnStatus  string `json:"delivOptnStatus"`
				ReqApptWdw       string `json:"reqApptWdw"`
				ReqApptDesc      string `json:"reqApptDesc"`
				RerouteTRKNbr    string `json:"rerouteTRKNbr"`
				BeginTm          string `json:"beginTm"`
				EndTm            string `json:"endTm"`
			} `json:"requestedAppointmentInfoList"`
			DefaultCDOType              string `json:"defaultCDOType"`
			ReturnAuthorizationName     string `json:"returnAuthorizationName"`
			TotalCustomsValueAmount     string `json:"totalCustomsValueAmount"`
			TotalCustomsValueCurrency   string `json:"totalCustomsValueCurrency"`
			PackageInsuredValueAmount   string `json:"packageInsuredValueAmount"`
			PackageInsuredValueCurrency string `json:"packageInsuredValueCurrency"`
			EstDelTimeWindow            struct {
				EstDelTmWindowStart          string `json:"estDelTmWindowStart"`
				EstDelTmWindowEnd            string `json:"estDelTmWindowEnd"`
				DisplayEstDelTmWindowTmStart string `json:"displayEstDelTmWindowTmStart"`
				DisplayEstDelTmWindowTmEnd   string `json:"displayEstDelTmWindowTmEnd"`
			} `json:"estDelTimeWindow"`
			StandardTransitTimeWindow struct {
				StdTransitTimeStart        string `json:"stdTransitTimeStart"`
				DisplayStdTransitTimeStart string `json:"displayStdTransitTimeStart"`
				StdTransitTimeEnd          string `json:"stdTransitTimeEnd"`
				DisplayStdTransitTimeEnd   string `json:"displayStdTransitTimeEnd"`
			} `json:"standardTransitTimeWindow"`
			StandardTransitDate struct {
				StdTransitDate        string `json:"stdTransitDate"`
				DisplayStdTransitDate string `json:"displayStdTransitDate"`
			} `json:"standardTransitDate"`
			PkgDimIn                  string `json:"pkgDimIn"`
			PkgDimCm                  string `json:"pkgDimCm"`
			ReturnedToShipperTrackNbr string `json:"returnedToShipperTrackNbr"`
			CommodityInfoList         []struct {
				CountryOfManufacture string `json:"countryOfManufacture"`
				HarmonizedCode       string `json:"harmonizedCode"`
				Description          string `json:"description"`
			} `json:"commodityInfoList"`
			StatusActionDesc             string      `json:"statusActionDesc"`
			DestinationGeoCoordinate     interface{} `json:"destinationGeoCoordinate"`
			ServiceCommitMessage         string      `json:"serviceCommitMessage"`
			ServiceCommitMessageType     string      `json:"serviceCommitMessageType"`
			LastUpdateDestinationAddress struct {
				StreetLineList                  []interface{} `json:"streetLineList"`
				City                            string        `json:"city"`
				StateOrProvinceCode             string        `json:"stateOrProvinceCode"`
				PostalCode                      string        `json:"postalCode"`
				CountryCode                     string        `json:"countryCode"`
				Residential                     bool          `json:"residential"`
				AddressVerificationID           string        `json:"addressVerificationId"`
				ShareID                         interface{}   `json:"shareId"`
				AddressClassification           interface{}   `json:"addressClassification"`
				AddressClassificationConfidence interface{}   `json:"addressClassificationConfidence"`
				Classification                  interface{}   `json:"classification"`
				UrbanizationCode                interface{}   `json:"urbanizationCode"`
				CountryName                     string        `json:"countryName"`
				GeographicCoordinates           interface{}   `json:"geographicCoordinates"`
				ProcessingParameters            interface{}   `json:"processingParameters"`
			} `json:"lastUpdateDestinationAddress"`
			HalAddressLocationID               interface{} `json:"halAddressLocationId"`
			StreetGeoCoordinate                interface{} `json:"streetGeoCoordinate"`
			IsIndirectSignatureReleaseEligible bool        `json:"isIndirectSignatureReleaseEligible"`
			IsRerouteEligible                  bool        `json:"isRerouteEligible"`
			IsRescheduleEligible               bool        `json:"isRescheduleEligible"`
			CoBrandedLogoLocation              interface{} `json:"coBrandedLogoLocation"`
			CoBrandedLogoURL                   interface{} `json:"coBrandedLogoUrl"`
			CoBrandedCouponLocation            interface{} `json:"coBrandedCouponLocation"`
			CoBrandedCouponURL                 interface{} `json:"coBrandedCouponUrl"`
			CoBrandedLogoAltTxt                interface{} `json:"coBrandedLogoAltTxt"`
			CoBrandedCouponAltTxt              interface{} `json:"coBrandedCouponAltTxt"`
			BrokerName                         string      `json:"brokerName"`
			BrokerCompanyName                  string      `json:"brokerCompanyName"`
			BuyerSoldToPartyName               string      `json:"buyerSoldToPartyName"`
			BuyerSoldToPartyCompanyName        string      `json:"buyerSoldToPartyCompanyName"`
			ImporterOfRecordCompanyName        string      `json:"importerOfRecordCompanyName"`
			ImporterOfRecordName               string      `json:"importerOfRecordName"`
			ConsolidationDetails               interface{} `json:"consolidationDetails"`
			ExclusionReasonDetails             interface{} `json:"exclusionReasonDetails"`
			PiecesPerShipment                  string      `json:"piecesPerShipment"`
			TotalPiecesPerMPSShipment          string      `json:"totalPiecesPerMPSShipment"`
			Matched                            bool        `json:"matched"`
			Delayed                            bool        `json:"delayed"`
			FxfAdvanceNotice                   bool        `json:"fxfAdvanceNotice"`
			Codrequired                        bool        `json:"codrequired"`
			Mpstype                            string      `json:"mpstype"`
			RthavailableCD                     string      `json:"rthavailableCD"`
			ExcepReasonListNoInit              []string    `json:"excepReasonListNoInit"`
			ExcepActionListNoInit              []string    `json:"excepActionListNoInit"`
			StatusDetailsListNoInit            []string    `json:"statusDetailsListNoInit"`
			IsExpiring                         bool        `json:"isExpiring"`
			IsInvalid                          bool        `json:"isInvalid"`
			IsReturn                           bool        `json:"isReturn"`
			IsException                        bool        `json:"isException"`
			IsExpired                          bool        `json:"isExpired"`
			IsDuplicate                        bool        `json:"isDuplicate"`
			ErrorList                          []struct {
				Code      string      `json:"code"`
				Message   string      `json:"message"`
				Source    interface{} `json:"source"`
				RootCause interface{} `json:"rootCause"`
			} `json:"errorList"`
			IsSuccessful                              bool `json:"isSuccessful"`
			IsPending                                 bool `json:"isPending"`
			IsCanceled                                bool `json:"isCanceled"`
			IsMPS                                     bool `json:"isMPS"`
			IsGMPS                                    bool `json:"isGMPS"`
			IsResidential                             bool `json:"isResidential"`
			IsConsolidationDetail                     bool `json:"isConsolidationDetail"`
			IsInProgress                              bool `json:"isInProgress"`
			IsPrePickup                               bool `json:"isPrePickup"`
			IsPickup                                  bool `json:"isPickup"`
			IsShipmentException                       bool `json:"isShipmentException"`
			IsClearanceDelay                          bool `json:"isClearanceDelay"`
			IsDelivered                               bool `json:"isDelivered"`
			IsInFedExPossession                       bool `json:"isInFedExPossession"`
			IsDroppedOff                              bool `json:"isDroppedOff"`
			IsDeliveryToday                           bool `json:"isDeliveryToday"`
			IsDocumentAvailable                       bool `json:"isDocumentAvailable"`
			IsSave                                    bool `json:"isSave"`
			IsHistorical                              bool `json:"isHistorical"`
			IsTenderedNotification                    bool `json:"isTenderedNotification"`
			IsDeliveredNotification                   bool `json:"isDeliveredNotification"`
			IsExceptionNotification                   bool `json:"isExceptionNotification"`
			IsEstimatedDeliveryDateChangeNotification bool `json:"isEstimatedDeliveryDateChangeNotification"`
			IsCurrentStatusNotification               bool `json:"isCurrentStatusNotification"`
			IsAnticipatedShipDtLabel                  bool `json:"isAnticipatedShipDtLabel"`
			IsShipPickupDtLabel                       bool `json:"isShipPickupDtLabel"`
			IsActualPickupLabel                       bool `json:"isActualPickupLabel"`
			IsOrderReceivedLabel                      bool `json:"isOrderReceivedLabel"`
			IsEstimatedDeliveryDtLabel                bool `json:"isEstimatedDeliveryDtLabel"`
			IsDeliveryDtLabel                         bool `json:"isDeliveryDtLabel"`
			IsActualDeliveryDtLabel                   bool `json:"isActualDeliveryDtLabel"`
			IsOrderCompleteLabel                      bool `json:"isOrderCompleteLabel"`
			IsTenderedDtLabel                         bool `json:"isTenderedDtLabel"`
			IsShipDtLabel                             bool `json:"isShipDtLabel"`
			IsEstDelTmWindowLabel                     bool `json:"isEstDelTmWindowLabel"`
			IsUnknownDirection                        bool `json:"isUnknownDirection"`
			IsSpod                                    bool `json:"isSpod"`
			IsSignatureAvailable                      bool `json:"isSignatureAvailable"`
			IsSignatureThumbnailAvailable             bool `json:"isSignatureThumbnailAvailable"`
			IsOriginalOutBound                        bool `json:"isOriginalOutBound"`
			IsChildPackage                            bool `json:"isChildPackage"`
			IsParentPackage                           bool `json:"isParentPackage"`
			IsReclassifiedAsSingleShipment            bool `json:"isReclassifiedAsSingleShipment"`
			IsMaskShipper                             bool `json:"isMaskShipper"`
			IsHalEligible                             bool `json:"isHalEligible"`
			IsFedexOfficeOnlineOrders                 bool `json:"isFedexOfficeOnlineOrders"`
			IsFedexOfficeInStoreOrders                bool `json:"isFedexOfficeInStoreOrders"`
			IsMultipleStop                            bool `json:"isMultipleStop"`
			IsHALResidential                          bool `json:"isHALResidential"`
			IsActualDelAddrResidential                bool `json:"isActualDelAddrResidential"`
			IsReqEstDelDt                             bool `json:"isReqEstDelDt"`
			IsMtchdByRecShrID                         bool `json:"isMtchdByRecShrID"`
			IsMtchdByShiprShrID                       bool `json:"isMtchdByShiprShrID"`
			IsHistoricalEDTW                          bool `json:"isHistoricalEDTW"`
			IsNonHistoricalEDTW                       bool `json:"isNonHistoricalEDTW"`
			IsTargetedMsg                             bool `json:"isTargetedMsg"`
			IsCommodityInfoAvail                      bool `json:"isCommodityInfoAvail"`
			IsStreetMapEligible                       bool `json:"isStreetMapEligible"`
			CDOInfoList                               []struct {
				SpclInstructDesc string `json:"spclInstructDesc"`
				DelivOptn        string `json:"delivOptn"`
				DelivOptnStatus  string `json:"delivOptnStatus"`
				ReqApptWdw       string `json:"reqApptWdw"`
				ReqApptDesc      string `json:"reqApptDesc"`
				RerouteTRKNbr    string `json:"rerouteTRKNbr"`
				BeginTm          string `json:"beginTm"`
				EndTm            string `json:"endTm"`
			} `json:"CDOInfoList"`
			IsFSM                    bool `json:"isFSM"`
			IsFreight                bool `json:"isFreight"`
			IsCustomCritical         bool `json:"isCustomCritical"`
			IsCDOEligible            bool `json:"isCDOEligible"`
			CDOExists                bool `json:"CDOExists"`
			IsExclusionReason        bool `json:"isExclusionReason"`
			IsNotFound               bool `json:"isNotFound"`
			IsInTransit              bool `json:"isInTransit"`
			IsOnSchedule             bool `json:"isOnSchedule"`
			IsDelException           bool `json:"isDelException"`
			IsHAL                    bool `json:"isHAL"`
			IsBeforePossessionStatus bool `json:"isBeforePossessionStatus"`
			IsWatch                  bool `json:"isWatch"`
			IsOutboundDirection      bool `json:"isOutboundDirection"`
			IsInboundDirection       bool `json:"isInboundDirection"`
			IsThirdpartyDirection    bool `json:"isThirdpartyDirection"`
			IsDestResidential        bool `json:"isDestResidential"`
		} `json:"packageList"`
		Alerts    interface{} `json:"alerts"`
		Cxsalerts interface{} `json:"cxsalerts"`
	} `json:"TrackPackagesResponse"`
}