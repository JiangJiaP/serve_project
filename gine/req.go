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

	//锻炼记录输入
	r.GET("/store", func(c *gin.Context) {

		var data model.Data
		data.MultiId = c.Query("multi_id")
		data.RouteId = c.Query("route_id")
		data.CId = c.Query("cid")

		c.JSON(http.StatusOK, gin.H{
			"state": "ok",
		})

	})

	r.GET("/get_route_id_from_cid", func(c *gin.Context) {
		cId := c.Query("cid")
		var routeId string
		errNo := "0"
		datas, err := mymysql.CIdSearch(cId)
		if err != nil {
			errNo = "1"
		} else {
			if len(datas) == 0 {
				errNo = "1"
			} else {
				routeId = datas[0].RouteId
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"err_no":   errNo,
			"route_id": routeId,
		})

	})
	r.GET("/get_cid_from_route_id", func(c *gin.Context) {
		rId := c.Query("route_id")
		datas := mymysql.RouteIdSearch(rId)
		cId := datas[0].CId

		c.JSON(http.StatusOK, gin.H{
			"cid": cId,
		})

	})

	r.GET("/get_multi_id_from_cid", func(c *gin.Context) {
		cId := c.Query("cid")
		datas, _ := mymysql.CIdSearch(cId)
		multiId := datas[0].RouteId

		c.JSON(http.StatusOK, gin.H{
			"multi_id": multiId,
		})
	})

	r.GET("/get_multi_id_from_route_id", func(c *gin.Context) {
		rId := c.Query("route_id")
		datas, _ := mymysql.CIdSearch(rId)
		multiId := datas[0].RouteId

		c.JSON(http.StatusOK, gin.H{
			"multi_id": multiId,
		})
	})

	r.Run()
}
