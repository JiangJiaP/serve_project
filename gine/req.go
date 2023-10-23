package gine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utahw/model"
	"utahw/mymysql"
)

func GinInit() {
	r := gin.Default()
	/*
			API 列表
			1. 存通信标识，多维标识，路由标识
			2. 根据通信标识获得路由表示
		    3. 通信表示获得多维标识
			4. rid 获得多维表示

	*/
	r.GET("/mysql_init", func(c *gin.Context) {
		mysqlPath := c.Query("address")

		err := mymysql.My_init(mysqlPath)
		state := "ok"
		if err != nil {
			state = "fail"
		}
		c.JSON(http.StatusOK, gin.H{
			"state": state,
		})

	})

	r.GET("/store", func(c *gin.Context) {

		var data model.Data
		data.MultiId = c.Query("multi_id")
		data.RouteId = c.Query("route_id")
		data.CId = c.Query("cid")
		data.Ifn = c.Query("ifn")
		data.MacId = c.Query("mac_id")

		mymysql.IdCreate(data)

		c.JSON(http.StatusOK, gin.H{
			"state": "ok",
		})

	})

	r.GET("/get_route_id_mac_ifn_from_cid", func(c *gin.Context) {
		cId := c.Query("cid")
		var routeId string
		var mac string
		var ifn string
		errNo := "0"
		datas, err := mymysql.CIdSearch(cId)
		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				routeId = datas[0].RouteId
				mac = datas[0].MacId
				ifn = datas[0].Ifn
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"err_no":   errNo,
			"route_id": routeId,
			"mac":      mac,
			"ifn":      ifn,
		})

	})
	r.GET("/get_cid_from_route_id", func(c *gin.Context) {
		rId := c.Query("route_id")
		datas, err := mymysql.RouteIdSearch(rId)
		errNo := "0"
		var cId string
		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				cId = datas[0].CId
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no": errNo,
			"cid":    cId,
		})

	})

	r.GET("/get_multi_id_from_cid", func(c *gin.Context) {
		cId := c.Query("cid")
		var errNo string
		datas, err := mymysql.CIdSearch(cId)
		var multiId string

		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				multiId = datas[0].MultiId
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no":   errNo,
			"multi_id": multiId,
		})
	})

	r.GET("/get_multi_id_from_route_id", func(c *gin.Context) {
		rId := c.Query("route_id")
		var errNo string
		datas, err := mymysql.RouteIdSearch(rId)
		var multiId string

		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				multiId = datas[0].MultiId
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no":   errNo,
			"multi_id": multiId,
		})

	})

	r.GET("/get_cid_from_multi_id",func(c *gin.Context){
		var data model.Data
		var errNo string
		var cid string
		data.UserId = c.Query("user_id")
		data.DeviceId = c.Query("device_id")
		data.AddressId = c.Query("address_id")
		data.ServiceId = c.Query("service_id")
		data.DataId = c.Query("data_id")

		datas, err := mymysql.CIdSearchByMid(data)
		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				cid = datas[0].CId
			}
		}

		c.JSON(http.StatusOK,gin.H{
			"err_no":   errNo,
			"cid": cid,
		})
	})


	r.GET("/get_sonic_ip", func(c *gin.Context) {
		var errNo string
//		datas ,err := mymysql.IpAddrSearchAllFromRouter()
//		if err != nil {
//			errNo = "1"
//		}
		datas := []model.RouterData{
			{"Router1", "192.168.1.1"},
			{"Router2", "192.168.1.2"},
		}



		c.JSON(http.StatusOK,gin.H{
			"err_no" : errNo,
			"datas" : datas,
		})
	})

	r.Run()
}
