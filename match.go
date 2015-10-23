package tba

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "os"
)

type Match struct{
  Key string `json:"key"`
  Level string `json:"comp_level"`
  Participants Participants `json:"alliances"`
}


type Participants struct{
  BlueAlliance Alliance `json:"blue"`
  RedAlliance Alliance `json:"red"`
}

type Alliance struct{
  Score int `json:"score"`
  Teams []string `json:"teams"`
}

func LoadMatches(eventKey string) []Match{
  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://www.thebluealliance.com/api/v2/event/" + eventKey + "/matches", nil)
  req.Header.Set("X-TBA-App-Id", os.Getenv("TBA_KEY"))
  res, _ := client.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  matches := []Match{}
  json.Unmarshal(body, &matches)
  return matches
}
