package gaode

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const GAODE_KEY = ""

func Test_SearchText(t *testing.T) {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	api := NewApi(GAODE_KEY)
	resp, e := api.SearchText(&SearchTextReq{Keywords: "北京大学"})
	if e != nil {
		t.Error("Test_SearchText", e, resp)
		return
	}
	log.Println(resp)

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
}

func Test_SearchAround(t *testing.T) {

	api := NewApi(GAODE_KEY)
	resp, e := api.SearchAround(&SearchAroundReq{Location: "116.456299,39.960767"})
	if e != nil {
		t.Error("Test_SearchAround", e, resp)
		return
	}
	log.Println(resp)

	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp)
}
