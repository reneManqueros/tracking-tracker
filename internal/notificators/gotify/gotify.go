package gotify

import (
	"net/http"
	"net/url"
	"trackingtracker/internal/config"
)

type Gotify struct {
}

func (g Gotify) Send(title string, message string) {
	http.PostForm(config.Settings.GotifyURL,
		url.Values{
			"message": {
				message,
			}, "priority": {"9"},
			"title": {
				title,
			}},
	)
}
