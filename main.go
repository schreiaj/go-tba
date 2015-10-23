package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)
import "github.com/davecgh/go-spew/spew"
import "github.com/influxdb/influxdb/client/v2"

type Event struct{
  Key string `json:"key"`
  Name string `json:"name"`
  Short_Name string `json:"short_name"`
  Event_Code string `json:"event_code"`
  Official bool `json:"official"`
  Year int `json:"year"`
  Start_Date string `json:"start_date"`
}

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

func main(){

  client := &http.Client{}
  req, _ := http.NewRequest("GET", "http://www.thebluealliance.com/api/v2/event/2014cmp/matches", nil)
  req.Header.Set("X-TBA-App-Id", `W/"frc125:dynasty:v02"`)
  res, _ := client.Do(req)
  defer res.Body.Close()
  body, _ := ioutil.ReadAll(res.Body)

  e := []Match{}
  json.Unmarshal(body, &e)
  spew.Dump(e)
}
