package cainiao

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"trackingtracker/internal/models"
)

func Crawl(tl models.TrackingLog) {
	trackingCode := tl.Code
	response, _ := http.Get("https://global.cainiao.com/detail.htm?mailNoList=" + trackingCode)
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	doc.Find(`#waybill_list_val_box`).Each(func(i int, selection *goquery.Selection) {
		var cainiaoJSON trackingResponse
		json.Unmarshal([]byte(selection.Text()), &cainiaoJSON)
		for _, v := range cainiaoJSON.Data[0].Section2.DetailList {
			sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`%s-%s-%s`, v.Desc, v.Time, v.Status)))
			if _, ok := tl.TrackingDetails[sEnc]; !ok {
				tl.TrackingDetails[sEnc] = models.TrackingDetail{
					Description: v.Desc,
					Time:        v.Time,
					Status:      v.Status,
				}

				title := fmt.Sprintf("New status for %s: %s", trackingCode, v.Status)
				message := fmt.Sprintf(`%s - %s - %s`, v.Time, v.Status, v.Desc)
				models.NotifierService.Send(title, message)
			}
		}
	})
	tl.Save()
}
