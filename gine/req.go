package gine

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"utahw/model"
	"utahw/mymysql"
	"utahw/service"
)

var StrategyAddress string

func GinInit() {
	r := gin.Default()
	/*
			API 列表
			1. 存通信标识，多维标识，路由标识
			2. 根据通信标识获得路由表示
		    3. 通信表示获得多维标识
			4. rid 获得多维表示

	*/
	r.GET("/address_init", func(c *gin.Context) {
		mysqlPath := c.Query("mysql_address")
		StrategyAddress = c.Query("strategy_address")
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
		var serviceId string
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
				serviceId = datas[0].ServiceId
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"err_no":     errNo,
			"route_id":   routeId,
			"mac":        mac,
			"ifn":        ifn,
			"service_id": serviceId,
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

	r.POST("/get_cid_from_multi_id", func(c *gin.Context) {
		var MultiData model.MultiData
		var errNo string
		var cids []string
		if err := c.ShouldBindJSON(&MultiData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err_no": "error",
			})
			return
		}
		var data model.Data
		data.UserId = MultiData.UserId
		data.DeviceId = MultiData.DeviceId
		data.AddressId = MultiData.AddressId
		data.ServiceId = MultiData.ServiceId
		data.DataId = MultiData.DataId

		datas, err := mymysql.CIdSearchByMid(data)

		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				for _, d := range datas {
					cids = append(cids, d.CId)
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no": errNo,
			"cids":   cids,
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

		errno := service.ConnectServicePost(scidInfo[0], dcidInfo[0], "http://"+routerInfo[0].Router+":40002")

		//连接泽军的
		if len(StrategyAddress) != 0 {
			service.SendStrategyServe(scidInfo[0].CId, dcidInfo[0].CId, "http://"+StrategyAddress+"/auth/policy/publishACL")
		}

		c.JSON(http.StatusOK, gin.H{
			"err_no": errno,
		})

	})

	r.Run()
}
