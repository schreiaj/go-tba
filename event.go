package tba

import (
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
