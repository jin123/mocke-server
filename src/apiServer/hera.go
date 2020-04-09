package apiServer

import (
	"encoding/json"
	"fmt"
	"net/http"

	commonfunc "github.com/jin123/mocke-server/src/common"
)

type ApiResult struct {
	Errno  int           `json:"errno"`
	Errmsg string        `json:"errmsg"`
	Data   []interface{} `json:"data"`
}

var apiRes ApiResult

type hera struct {
	MethodMap map[string]string
}

func CreateInstance() *hera {
	return new(hera)
}
func (my *hera) SetMethod() {

	my.MethodMap = map[string]string{
		"productGetHotelsDetailindex":           "GetHotelsDetail",
		"productGetAllHotelindex":               "GetAllHotel",
		"productDoubleCheckbeforeCheckRoomType": "DoubleCheckBeforeCheckRoomType",
		"productGetRoomPriceOneHotelindex":      "GetRoomPriceOneHotel",
		"productGetMultiHotelRoomStockindex":    "GetMultiHotelRoomStock",
	}
}
func (my *hera) GetHotelsDetail(r *http.Request, jsonStr string) string {
	return jsonStr
}

func (my *hera) DoubleCheckBeforeCheckRoomType(r *http.Request, jsonStr string) string {
	dates := commonfunc.GetBetweenDates(r.PostFormValue("check_in"), r.PostFormValue("check_out"))
	err := json.Unmarshal([]byte(jsonStr), &apiRes)
	if err != nil {
		fmt.Println("json数据解析失败: ", err)
	}
	apiData := apiRes.Data
	dataMap := apiData[0]
	fmt.Println("dataMap=", dataMap)
	roomDetailMap := dataMap.(map[string]interface{})["room_count_detail"]
	eachRoomDetailMap := roomDetailMap.([]interface{})[0]
	roomCountDetail := make([]interface{}, len(dates))
	for i, selectDate := range dates {
		newMaps := make(map[string]interface{})
		for key, value := range eachRoomDetailMap.(map[string]interface{}) {
			newMaps[key] = value
		}
		newMaps["biz_date"] = selectDate + " 00:00:00"
		roomCountDetail[i] = newMaps
	}
	dataMap.(map[string]interface{})["room_count_detail"] = roomCountDetail
	apiData[0] = dataMap
	apiRes.Data = apiData
	newResponse, err := json.Marshal(apiRes)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	jsonStr = string(newResponse)
	fmt.Println(jsonStr)

	return jsonStr
}

//通用的格式化时间的数据
func (my *hera) FormatApiResponsV1(startTime string, endTime string, fileld string, jsonStr string) string {
	dates := commonfunc.GetBetweenDates(startTime, endTime)
	responseData := make([]interface{}, len(dates))
	err := json.Unmarshal([]byte(jsonStr), &apiRes)
	if err != nil {
		fmt.Println("json数据解析失败: ", err)
	}
	apiData := apiRes.Data
	template := apiData[0]
	for i, selectDate := range dates {
		maps := template.(map[string]interface{})
		newMaps := make(map[string]interface{})
		for key, value := range maps {
			newMaps[key] = value

		}
		newMaps[fileld] = selectDate + " 00:00:00"
		responseData[i] = newMaps
	}
	fmt.Println("responseData=", responseData)
	apiRes.Data = responseData
	newResponse, err := json.Marshal(apiRes)
	if err != nil {
		fmt.Println("JSON ERR:", err)
	}
	jsonStr = string(newResponse)
	fmt.Println(jsonStr)
	return jsonStr
}

//获取酒店房型价格
func (my *hera) GetRoomPriceOneHotel(r *http.Request, jsonStr string) string {
	return my.FormatApiResponsV1(r.PostFormValue("start_time"), r.PostFormValue("end_time"), "biz_date", jsonStr)

}

//获取酒店的房型的库存
func (my *hera) GetMultiHotelRoomStock(r *http.Request, jsonStr string) string {
	return my.FormatApiResponsV1(r.PostFormValue("start_time"), r.PostFormValue("end_time"), "biz_date", jsonStr)
}

//返回所有酒店id集合数据暂时不动
func (my *hera) GetAllHotel(r *http.Request, jsonStr1 string) string {
	return jsonStr1

}
