package tba

import (
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Event struct{
  Key string `json:"key"`
  Name string `json:"name"`
  Short_Name string `json:"short_name"`
  Event_Code string `json:"event_code"`
  Official bool `json:"official"`
  Year int `json:"year"`
  Start_Date string `json:"start_date"`
}

func LoadEvents(year string) []Event{
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://www.thebluealliance.com/api/v2/events/" + year, nil)
  req.Header.Set("X-TBA-App-Id", os.Getenv("TBA_KEY"))
  res, _ := client.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  events := []Event{}
  json.Unmarshal(body, &events)
  return events
}
