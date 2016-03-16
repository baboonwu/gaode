package gaode

import (
	"net/url"
)

const (
	SEARCH_TEXT   = "http://restapi.amap.com/v3/place/text?"
	SEARCH_AROUND = "http://restapi.amap.com/v3/place/around?"
)

type Api struct {
	Key string
}

func NewApi(key string) *Api {
	return &Api{
		Key: key,
	}
}

func (api *Api) clientParams() url.Values {
	return url.Values{
		"key": []string{api.Key},
	}
}
