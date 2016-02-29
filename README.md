Send a LaMetric push notification using go.

Check [LaMetric developer site](https://developer.lametric.com) for more information.

```go
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
```
