package myapi

import "fmt"

type myApi struct {
}

func (api *myApi) productGetHotelsDetailindex() {

	fmt.Println("获取酒店详情接口")
}

func productGetAllHotelindex() {
	fmt.Println("获取所有酒店ID")
}
