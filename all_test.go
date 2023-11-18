package whatever

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestGetProjInfo(t *testing.T) {

	j, _ := os.ReadFile("service_account.json")

	c := make(map[string]json.RawMessage)

	// unmarschal JSON
	e := json.Unmarshal(j, &c)

	fmt.Println(string(c["project_id"]))

	t.Error(e)

	// byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	// var dat map[string]interface{}
}
