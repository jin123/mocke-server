package commonfunc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

//动态调用结构体的方法
func DynamicInvoke(object interface{}, methodName string, args ...interface{}) string {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	r := reflect.ValueOf(object).MethodByName(methodName).Call(inputs)
	value := r[0]
	v2 := value.Interface().(string)
	return v2
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func JsonToMap(jsonStr string) interface{} {
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("json to map err: ", err)
		return ""
	}
	return mapResult
}

func MapToJson(maps interface{}) string {

	jsonStr, err := json.Marshal(maps)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	//fmt.Println(string(jsonStr))
	jsons := string(jsonStr)
	return jsons
}
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func FormatMap(field string, maps interface{}) {
	fmt.Println("maps=", maps.([]interface{}))

	//newMap = make(map[string]interface{})

}

//结构体转字典
func StructToMap(obj []interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data

}

func StructToJson(obj interface{}) string {
	fmt.Println("obj=", obj)

	return ""
}
