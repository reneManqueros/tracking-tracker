package usps

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"trackingtracker/chrome"
	"trackingtracker/internal/models"
)

func Crawl(tl models.TrackingLog) {
	tl.TrackingDetails = make(map[string]models.TrackingDetail)

	resp, _ := chrome.Process(chrome.Request{
		URL:      "https://tools.usps.com/go/TrackConfirmAction?tLabels=" + tl.Code,
		LoadWait: 10,
		Domain:   "",
	})

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.Text))

	var statusList []string
	doc.Find(`#trackingHistory_1 span`).Each(func(i int, selection *goquery.Selection) {
		statusLine := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(selection.Text()), "\n", " "), "\t", "")
		statusList = append(statusList, statusLine)
	})
	td := models.TrackingDetail{}
	for index, statusMessage := range statusList {
		switch index % 3 {
		case 0:
			td.Time = statusMessage
		case 1:
			td.Status = statusMessage
		case 2:
			td.Description = statusMessage
			sEnc := b64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`%s-%s-%s`, td.Time, td.Status, td.Description)))
			if _, ok := tl.TrackingDetails[sEnc]; !ok {
				tl.TrackingDetails[sEnc] = td

				title := fmt.Sprintf("New status for %s: %s", tl.Code, td.Status)
				message := fmt.Sprintf(`%s - %s - %s`, td.Time, td.Status, td.Description)
				models.NotifierService.Send(title, message)
			}

			td = models.TrackingDetail{}
		}
	}
	tl.Save()
}
