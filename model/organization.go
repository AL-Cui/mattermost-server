package model

import (
	"encoding/json"
	"io"
)

type Org struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

func (o *Org) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

func OrgFromJson(data io.Reader) *Org {
	var o *Org
	json.NewDecoder(data).Decode(&o)
	return o
}
