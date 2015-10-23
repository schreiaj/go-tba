package tba

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
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
