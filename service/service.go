package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"utahw/model"
)



func decodeMultiData(jsonData string) (model.MultiData, error) {
	var multiData model.MultiData
	err := json.Unmarshal([]byte(jsonData), &multiData)
	return multiData, err
}

func ConnectServicePost(scidInfo model.Data, dcidInfo model.Data, url string)(errno string){



	requestData := map[string]interface{}{
		"datas": []map[string]string{
			{
				"cid":      scidInfo.CId,
				"route_id": scidInfo.RouteId,
				"mac":      scidInfo.MacId,
				"ifn":      scidInfo.Ifn,
			},
			{
				"cid":      dcidInfo.CId,
				"route_id": dcidInfo.RouteId,
				"mac":      dcidInfo.MacId,
				"ifn":      dcidInfo.Ifn,
			},
		},
		"err_no": "",
	}

	// 将请求体转换为JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSON编码失败:", err)
		return
	}


	// 发送POST请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("POST请求失败:", err)
		errno = "400"
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		fmt.Println("响应解码失败:", err)
		errno = "400"
		return
	}
	errno = response["err_no"]

	fmt.Println("响应体:", response)
	return errno
}
