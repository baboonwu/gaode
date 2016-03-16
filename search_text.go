package gaode

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// 参见 http://restapi.amap.com/v3/place/text?parameters

type SearchTextReq struct {

	// keywords和types两者至少必选其一

	// 查询关键字. 规则：多个关键字用“|”分割
	Keywords string

	// 查询 POS 类型，多个类型用“|”分割；
	// 可选值：文本分类、分类代码（建议使用分类代码，避免文本分类输入错误操作的搜索失败）
	// 分类代码由六位数字组成，后四位为0代表大类名称，后两位为0代表小类名称，如需搜索大类下所有分类，输入去掉后尾0。
	// 例如：180000为道路附属服务， 全类别下搜索types=18;搜索下一分类警示信息，types=1801; 搜索再下级分类，types=180101
	Types string

	// 可选值：城市中文、中文全拼、citycode、adcode, 如：北京/beijing/010/110000
	City string

	// 仅返回指定城市数据. 可选，默认false
	Citylimit bool

	// 是否按照层级展示子POI数据。可选，children=1
	Children int

	// 最大每页记录数为50条。超出取值范围按最大值返回，可选，默认 20
	Offset int

	// 当前页数, 最大翻页数100
	Page int

	// 返回数据格式类型, 可选值：json,xml，默认 json
	Output string
}

type GDResp struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	POIs     []POI  `json:"pois"`
}

type POI struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Tag      string
	Type     string
	Typecode string
	BizType  string
	Address  string      `json:"address"`
	Location string      `json:"location"` // Lng,Lat
	Tel      interface{} `json:"tel"`
	Postcode interface{}
	Websit   interface{}
	Email    interface{}
	Pcode    string
	Pname    string
	Citycode string `json:"citycode"`
	Cityname string `json:"cityname"`
	Adcode   string
	Adname   string
}

// 关键字搜索
func (api *Api) SearchText(req *SearchTextReq) (*GDResp, error) {

	// set param
	params := api.clientParams()
	params.Set("keywords", req.Keywords)

	url := SEARCH_TEXT + params.Encode()
	log.Println(url)

	http.DefaultClient.Timeout = 5 * time.Second // 超时 5s
	r, e := http.Get(url)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	defer r.Body.Close()

	// result
	resp := &GDResp{}
	dec := json.NewDecoder(r.Body)
	if e := dec.Decode(resp); e != nil {
		log.Println(e)
		return nil, e
	}

	return resp, nil
}
