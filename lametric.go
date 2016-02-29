/*
Send a LaMetric push notification using go.

Check (LaMetric developer site](https://developer.lametric.com) for more information.

	url := "https://developer.lametric.com/api/V1/dev/widget/update/com.lametric.APP_ID"
	token := "APP_TOKEN"

	lametric.Push(url, token, &App{
		Frames: []lametric.Frame{
			lametric.Frame{
				Index: 0,
				Text:  "some text",
				Icon:  "666",
			},
			lametric.Frame{
				Index: 1,
				Icon:  "667",
				GoalData: &lametric.Goal{
					End:     100,
					Current: 66,
					Unit:    "%",
				},
			},
		},
	})

*/
package lametric

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// App is composed of Frames
type App struct {
	Frames []Frame `json:"frames"`
}

// Frame is composed of optional Text, Icon or GoalData
type Frame struct {
	Index    int    `json:"index"`
	Text     string `json:"text,omitempty"`
	Icon     string `json:"icon,omitempty"`
	GoalData *Goal  `json:"goalData,omitempty"`
}

// Goal contains data for numeric Frames
type Goal struct {
	Start   int    `json:"start"`
	Current int    `json:"current"`
	End     int    `json:"end"`
	Unit    string `json:"unit"`
}

// Push updates the LaMetric App by sending a push notifiation
func Push(url, accessToken string, app *App) error {
	payload, err := json.Marshal(app)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	r.Header.Set("Accept", "application/json")
	r.Header.Set("X-Access-Token", accessToken)
	r.Header.Set("Cache-Control", "no-cache")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
