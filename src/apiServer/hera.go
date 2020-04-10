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
		"productDoubleCheckbeforeCreateOrder":   "DoubleCheckBeforeCreateOrder",
		"productGetRoomPriceOneHotelindex":      "GetRoomPriceOneHotel",
		"productGetMultiHotelRoomStockindex":    "GetMultiHotelRoomStock",
	}
}

//获取酒店详情
func (my *hera) GetHotelsDetail(r *http.Request, jsonStr string) string {
	return jsonStr
}

//标准价下单之前对房量和房价进行验证（酒店下所有房型），频率无限制
func (my *hera) DoubleCheckBeforeCheckRoomType(r *http.Request, jsonStr string) string {
	//dates := commonfunc.GetBetweenDates(r.PostFormValue("check_in"), r.PostFormValue("check_out"))
	str := my.FormatApiResponsV2(r.PostFormValue("check_in"), r.PostFormValue("check_out"), "room_count_detail", "biz_date", jsonStr)
	return str
}

//标准价下单之前对房量和房价进行验证 ，频率无限制
func (my *hera) DoubleCheckBeforeCreateOrder(r *http.Request, jsonStr string) string {
	dates := commonfunc.GetBetweenDates(r.PostFormValue("check_in"), r.PostFormValue("check_out"))
	responseData := make([]interface{}, len(dates))
	apiResultMap := commonfunc.JsonToMap(jsonStr)

	allData := apiResultMap.(map[string]interface{})["data"]
	detailArrs := allData.(map[string]interface{})["room_price"]
	detailArrs = detailArrs.([]interface{})[0]
	for i, selectDate := range dates {
		newMaps := make(map[string]interface{})
		for key, value := range detailArrs.(map[string]interface{}) {
			newMaps[key] = value

		}
		newMaps["biz_date"] = selectDate + " 00:00:00"
		responseData[i] = newMaps
	}
	allData.(map[string]interface{})["room_price"] = responseData
	apiResultMap.(map[string]interface{})["data"] = allData
	jsonStr = commonfunc.MapToJson(apiResultMap)
	//fmt.Println("apiResultMap=", apiResultMap)
	return jsonStr
}

//通用的格式化时间的数据V1
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

//通用的格式化时间的数据V2 field 要格式化的字段，secondField是date下的二维map字段
func (my *hera) FormatApiResponsV2(startTime string, endTime string, secondField string, fileld string, jsonStr string) string {
	dates := commonfunc.GetBetweenDates(startTime, endTime)
	err := json.Unmarshal([]byte(jsonStr), &apiRes)
	if err != nil {
		fmt.Println(jsonStr)
		fmt.Println("json数据解析失败v2: ", err)
		return ""
	}
	apiData := apiRes.Data
	dataMap := apiData[0]
	fmt.Println("dataMap=", dataMap)
	roomDetailMap := dataMap.(map[string]interface{})[secondField]
	eachRoomDetailMap := roomDetailMap.([]interface{})[0]
	roomCountDetail := make([]interface{}, len(dates))
	for i, selectDate := range dates {
		newMaps := make(map[string]interface{})
		for key, value := range eachRoomDetailMap.(map[string]interface{}) {
			newMaps[key] = value
		}
		newMaps[fileld] = selectDate + " 00:00:00"
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
