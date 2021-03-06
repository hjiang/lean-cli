package stats

import (
	"runtime"

	"github.com/levigross/grequests"
)

// LeanCloud app info
var (
	leanCloudAppID  string
	leanCloudAppKey string
)

// Client is current user's client info
var Client ClientType

var session SessionType

// Init LeanCloud app info
func Init(appID string, appKey string) error {
	leanCloudAppID = appID
	leanCloudAppKey = appKey

	deviceID, err := GetDeviceID()
	if err != nil {
		return err
	}

	Client.ID = deviceID
	Client.Platform = runtime.GOOS
	session.ID, err = newDeviceID()
	return err
}

// Event is collect payload's evnets field type
type Event struct {
	Event string `json:"event"`
}

// ClientType is collect payload's client field type
type ClientType struct {
	ID         string `json:"id"`
	Platform   string `json:"platform"`
	AppVersion string `json:"app_version"`
	AppChannel string `json:"app_channel"`
}

// SessionType is collect payload's session filed type
type SessionType struct {
	ID string `json:"id"`
}

// Payload is leancloud statics collect's playload
type Payload struct {
	Client  ClientType  `json:"client"`
	Events  []Event     `json:"events"`
	Session SessionType `json:"session"`
}

// Collect the user's stats
func Collect(events []Event) {
	payload := &Payload{
		Client:  Client,
		Events:  events,
		Session: session,
	}
	grequests.Post("https://api.leancloud.cn/1.1/stats/open/collect", &grequests.RequestOptions{
		Headers: map[string]string{
			"X-LC-Id":  leanCloudAppID,
			"X-LC-Key": leanCloudAppKey,
		},
		JSON: payload,
	})
}
