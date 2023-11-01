package gine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utahw/model"
	"utahw/mymysql"
	"utahw/service"
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

	r.GET("/get_cid_from_multi_id", func(c *gin.Context) {
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

		c.JSON(http.StatusOK, gin.H{
			"err_no": errNo,
			"cid":    cid,
		})
	})

	r.GET("/get_sonic_ip", func(c *gin.Context) {
		var errNo string
		datas, err := mymysql.IpAddrSearchAllFromRouter()
		if err != nil {
			errNo = "1"
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no": errNo,
			"datas":  datas,
		})
	})

	r.POST("/store_sonic_router", func(c *gin.Context) {
		var data model.RouterData

		// 解析请求正文中的JSON 数据
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
			return
		}

		mymysql.SonicRouterStore(data)
		c.JSON(http.StatusOK, gin.H{
			"state": "ok",
		})
	})

	r.POST("/connect", func(c *gin.Context) {
		var data model.ConnectData

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
			return
		}

		dcidInfo, err := mymysql.CIdSearch(data.Dcid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
		}

		routerInfo, err := mymysql.SonicRouterSearchFromIpaddress(dcidInfo[0].RouteId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
		}

		scidInfo, err := mymysql.CIdSearch(data.Scid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
		}

		service.ConnectServicePost(scidInfo[0], dcidInfo[0], routerInfo[0].IpAddr+":40002/inner_connect")

		//连接泽军的

	})

	r.Run()
}
