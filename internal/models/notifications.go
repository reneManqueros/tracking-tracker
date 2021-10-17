package models

import (
	"trackingtracker/internal/config"
	"trackingtracker/internal/notificators/gotify"
)

type NotificationProvider interface {
	Send(string, string)
}

var NotifierService Notifier

type Notifier struct {
}

func (n *Notifier) Send(title string, message string) {
	availableNotifiers := map[string]NotificationProvider{
		"gotify": gotify.Gotify{},
	}
	enabledNotifiers := make(map[string]NotificationProvider)
	for _, notifier := range config.Settings.NotificationProviders {
		enabledNotifiers[notifier] = availableNotifiers[notifier]
	}

	for _, v := range enabledNotifiers {
		v.Send(title, message)
	}
}
