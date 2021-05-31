package ch42

import (
	"encoding/json"
	"testing"
)

type BaseInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BaseInfo BaseInfo `json:"base_info"`
	JobInfo  JobInfo  `json:"job_info"`
}

var jsonStr = `{
"base_info":{"name":"jiawei","age":999},
"job_info":{"skills":["php", "python", "go", "js"]}
}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(e)
	if v, err := json.Marshal(e); err == nil {
		t.Log(string(v))
	} else {
		t.Error(err)
	}
}
