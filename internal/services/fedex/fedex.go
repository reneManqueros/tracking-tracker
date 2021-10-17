package fedex

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"trackingtracker/internal/models"
)

func Crawl(tl models.TrackingLog) {
	body := strings.NewReader("action=trackpackages&data=%7B%22TrackPackagesRequest%22:%7B%22appDeviceType%22:%22DESKTOP%22,%22appType%22:%22WTRK%22,%22processingParameters%22:%7B%7D,%22uniqueKey%22:%22%22,%22supportCurrentLocation%22:true,%22supportHTML%22:true,%22trackingInfoList%22:%5B%7B%22trackNumberInfo%22:%7B%22trackingNumber%22:%22" + tl.Code + "%22,%22trackingCarrier%22:null%7D%7D%5D%7D%7D&format=json&locale=en_US&version=1")
	req, err := http.NewRequest("POST", "https://www.fedex.com/trackingCal/track", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Adrum", "isAjax:true")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Set("Origin", "https://www.fedex.com")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,es;q=0.8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	var trackingResponse trackingResponse
	json.Unmarshal(b, &trackingResponse)
	for _, td := range trackingResponse.TrackPackagesResponse.PackageList[0].ScanEventList {
		title := fmt.Sprintf("New status for %s: %s", tl.Code, td.Status)
		message := fmt.Sprintf(`%s %s- %s - %s`, td.Date, td.Time, td.Status, td.ScanLocation)
		sEnc := b64.StdEncoding.EncodeToString([]byte(title))
		if _, ok := tl.TrackingDetails[sEnc]; !ok {
			tl.TrackingDetails[sEnc] = models.TrackingDetail{
				Description: td.ScanLocation,
				Time:        fmt.Sprintf(`%s %s`, td.Date, td.Time),
				Status:      td.Status,
			}
			models.NotifierService.Send(title, message)
		}
	}
	tl.Save()
}
