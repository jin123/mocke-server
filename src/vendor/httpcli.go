package vendor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/jin123/mocke-server/src/apiServer"
	commonfunc "github.com/jin123/mocke-server/src/common"
)

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

type API struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Endpoints []Endpoint `json:"endpoints"`
}

var api API

//生成一个htt服务
func CreateHttpConnect(port int) {
	dir := commonfunc.GetCurrentPath()
	raw, err := ioutil.ReadFile(dir + "/json/api.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//apiMap := make(map[string]interface{}, 0)
	json.Unmarshal(raw, &api)
	if err != nil {
		log.Fatal(" ", err)
	}
	for _, ep := range api.Endpoints {
		log.Print(ep)
		if len(ep.Folder) > 0 {
			http.Handle(ep.Path+"/", http.StripPrefix(ep.Path+"/", http.FileServer(http.Dir(ep.Folder))))
		} else {
			http.HandleFunc(ep.Path, response)
		}
	}
	err = http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(" ", err)
	}
}

func response(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	requestPath := r.URL.Path
	//allPath := api.Endpoints
	/*if urlPath, ok := allPath[requestPath]; ok {
		if urlPath["Method"] == r.Method {
			w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
			w.WriteHeader(urlPath.Status)
			myc := apiServer.CreateInstance() //以后改
			myc.SetMethod()
			requestMethod := strings.Replace(requestPath, "/", "", -1)
			allMethod := myc.MethodMap
			s := path2Response(urlPath.JsonPath)
			if apiFunc, ok := allMethod[requestMethod]; ok {
				c := commonfunc.DynamicInvoke(myc, apiFunc, r, s)
				b := []byte(c)
				w.Write(b)
			} else {
				fmt.Println("currect path is no method")
				b := []byte(s)
				w.Write(b)
			}
		}

	}
	fmt.Println(" currect path error !!!!!")
	ret := make(map[string]interface{})
	ret["errno"] = 500
	ret["errmsg"] = "error"
	jsonStr := commonfunc.MapToJson(ret)
	b := []byte(jsonStr)
	w.Write(b)
	*/
	for _, ep := range api.Endpoints {
		if requestPath == ep.Path && r.Method == ep.Method {
			w.Header().Set(HeaderContentType, MIMETextPlainCharsetUTF8)
			w.WriteHeader(ep.Status)
			myc := apiServer.CreateInstance() //以后改
			myc.SetMethod()
			requestMethod := strings.Replace(requestPath, "/", "", -1)
			allMethod := myc.MethodMap
			s := path2Response(ep.JsonPath)
			//apiFunc := allMethod[requestMethod]
			if apiFunc, ok := allMethod[requestMethod]; ok {
				c := commonfunc.DynamicInvoke(myc, apiFunc, r, s)
				b := []byte(c)
				w.Write(b)
			} else {
				b := []byte(s)
				w.Write(b)
				fmt.Println("currect path is no method")
			}

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
	return buf.String()
}

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}
