package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/jin123/mocke-server/src/logger"
)

var methodMap map[string]string

const (
	charsetUTF8 = "charset=UTF-8"
)

const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)

const (
	HeaderAccept                        = "Accept"
	HeaderAcceptEncoding                = "Accept-Encoding"
	HeaderAllow                         = "Allow"
	HeaderAuthorization                 = "Authorization"
	HeaderContentDisposition            = "Content-Disposition"
	HeaderContentEncoding               = "Content-Encoding"
	HeaderContentLength                 = "Content-Length"
	HeaderContentType                   = "Content-Type"
	HeaderCookie                        = "Cookie"
	HeaderSetCookie                     = "Set-Cookie"
	HeaderIfModifiedSince               = "If-Modified-Since"
	HeaderLastModified                  = "Last-Modified"
	HeaderLocation                      = "Location"
	HeaderUpgrade                       = "Upgrade"
	HeaderVary                          = "Vary"
	HeaderWWWAuthenticate               = "WWW-Authenticate"
	HeaderXForwardedFor                 = "X-Forwarded-For"
	HeaderXForwardedProto               = "X-Forwarded-Proto"
	HeaderXForwardedProtocol            = "X-Forwarded-Protocol"
	HeaderXForwardedSsl                 = "X-Forwarded-Ssl"
	HeaderXUrlScheme                    = "X-Url-Scheme"
	HeaderXHTTPMethodOverride           = "X-HTTP-Method-Override"
	HeaderXRealIP                       = "X-Real-IP"
	HeaderXRequestID                    = "X-Request-ID"
	HeaderServer                        = "Server"
	HeaderOrigin                        = "Origin"
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"
	HeaderStrictTransportSecurity       = "Strict-Transport-Security"
	HeaderXContentTypeOptions           = "X-Content-Type-Options"
	HeaderXXSSProtection                = "X-XSS-Protection"
	HeaderXFrameOptions                 = "X-Frame-Options"
	HeaderContentSecurityPolicy         = "Content-Security-Policy"
	HeaderXCSRFToken                    = "X-CSRF-Token"
)

type Endpoint struct {
	Type     string `json:"type"`
	Method   string `json:"method"`
	Status   int    `json:"status"`
	Path     string `json:"path"`
	JsonPath string `json:"jsonPath"`
	Folder   string `json:"folder"`
}
type ApiResult struct {
	Errno  int           `json:"errno"`
	Errmsg string        `json:"errmsg"`
	Data   []interface{} `json:"data"`
}

type API struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Endpoints []Endpoint `json:"endpoints"`
}

var apiRes ApiResult

var api API

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}

type myconnector struct {
}

func (my *myconnector) getMethod() {

	methodMap = map[string]string{
		"productGetHotelsDetailindex":           "GetHotelsDetail",
		"productGetAllHotelindex":               "GetAllHotel",
		"productDoubleCheckbeforeCheckRoomType": "DoubleCheckBeforeCheckRoomType",
		"productGetRoomPriceOneHotelindex":      "GetRoomPriceOneHotel",
	}
}
func (my *myconnector) GetHotelsDetail(r *http.Request, jsonStr string) string {
	return jsonStr
}

func (my *myconnector) DoubleCheckBeforeCheckRoomType(r *http.Request, jsonStr string) string {
	return jsonStr
}

func (my *myconnector) GetRoomPriceOneHotel(r *http.Request, jsonStr string) string {

	dates := GetBetweenDates(r.PostFormValue("start_time"), r.PostFormValue("end_time"))
	responseData := make([]interface{}, len(dates))
	//yourMap := []interface{}
	//fmt.Println("jsonStr=", jsonStr)
	err := json.Unmarshal([]byte(jsonStr), &apiRes)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	//fmt.Println("errno=", apiRes.Errno, "errormsg=", apiRes.Errmsg, "data=", apiRes.Data)
	apiData := apiRes.Data
	template := apiData[0]
	fmt.Println("type template=", reflect.TypeOf(template))
	for i, n := range dates {
		template.(map[string]interface{})["biz_date"] = n
		responseData[i] = template
		//responseData = append(responseData, template)
	}
	//fmt.Println("type apiData=", reflect.TypeOf(apiData))
	//fmt.Println("apiData=", apiData)
	//fmt.Println("responseData=", responseData)
	//apiRes.Data = yourMap

	/*newResponse, err := json.Marshal(apiRes)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}
	fmt.Println(newResponse)*/
	//fmt.Println("yourMap=", yourMap)
	//var newData map[string]interface{}

	//fmt.Println("biz date=", newData["biz_date"])

	//fmt.Println("mapResult=", mapResult)
	//fmt.Println("data type=", reflect.TypeOf(apiRes.Data))

	return jsonStr
}

//返回所有酒店id集合数据暂时不动
func (my *myconnector) GetAllHotel(r *http.Request, jsonStr1 string) string {
	return jsonStr1

}
func main() {

	//init_redis()
	raw, err := ioutil.ReadFile("./api.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &api)
	fmt.Println(api)
	if err != nil {
		log.Fatal(" ", err)
	}
	fmt.Println(api.Endpoints)
	for _, ep := range api.Endpoints {
		log.Print(ep)
		if len(ep.Folder) > 0 {
			http.Handle(ep.Path+"/", http.StripPrefix(ep.Path+"/", http.FileServer(http.Dir(ep.Folder))))
		} else {
			http.HandleFunc(ep.Path, response)
		}
	}
	httpPort := flag.Int("port", 4000, "this is http port")
	flag.Parse() //这个函数一定要放在这个位子
	fmt.Println("当前服务端口：", *httpPort)
	err = http.ListenAndServe(":"+strconv.Itoa(*httpPort), nil)

	if err != nil {
		log.Fatal(" ", err)
	}

}

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
func response(w http.ResponseWriter, r *http.Request) {

	appLogger := logger.CreateLogger()
	fmt.Println("请求的参数")
	//defer r.Body.Close()
	//fmt.Println(r.PostFormValue("start_time"))
	//fmt.Println(r.PostFormValue("end_time"))
	//fmt.Println(GetBetweenDates(r.PostFormValue("start_time"), r.PostFormValue("end_time")))

	r.ParseForm()
	appLogger.AccessLog(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	for _, ep := range api.Endpoints {
		if r.URL.Path == ep.Path && r.Method == ep.Method {
			fmt.Println("method:", r.Method)
			fmt.Println("path:", r.URL.Path)

			myc := new(myconnector)
			myc.getMethod()
			requestMethod := strings.Replace(r.URL.Path, "/", "", -1)
			//fmt.Println("output222")
			apiFunc := methodMap[requestMethod]
			fmt.Println("调用结构体的方法")

			w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
			w.WriteHeader(ep.Status)
			s := path2Response(ep.JsonPath)

			apiData := DynamicInvoke(myc, apiFunc, r, s)

			b := []byte(apiData)

			w.Write(b)
		}
		continue
	}
}

func path2Response(path string) string {

	//myapi.$apiFunc()
	file, err := os.Open(path)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	fmt.Println("返回值")
	fmt.Println(buf.String())
	return buf.String()
}

func init_redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	fmt.Println(client)
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("ping result: %s\n", pong)

	fmt.Printf("----------------------------------------\n")

	value, err := client.Get("test").Result()
	if err != nil {
		fmt.Printf("try get key[foo] error[%s]\n", err.Error())
		// err_handler(err)
	}
	fmt.Println(value)
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
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