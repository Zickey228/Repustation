package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sort"
	"strconv"
)

type dataJSON struct {
	ID         int    `json:"ID"`
	Sender     string `json:"sender"`
	Seller     string `json:"seller"`
	Comment    string `json:"comment"`
	Investment int    `json:"investment"`
	Polarity   int    `json:"polarity"`
	Ratings    int    `json:"ratings"`
}

//发送给全节点

func submitReview(sender string, seller string, comment string, investment string, polarity string) {
	conn, err := net.Dial("tcp", ip2+":9887")
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error())

	}
	defer conn.Close()

	invest, _ := strconv.Atoi(investment)
	pol, _ := strconv.Atoi(polarity)

	// JSON 格式的数据
	data := dataJSON{
		Sender:     sender,
		Seller:     seller,
		Comment:    comment,
		Investment: invest,
		Polarity:   pol,
		Ratings:    0,
	}

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON encode error:", err)
		return
	}

	_, err = conn.Write(json)
	if err != nil {
		fmt.Println("Send json error:", err)
		return
	}

	fmt.Println("json sent to server")
}

func checkReview(target string) bool {
	//从全节点取得数据
	//向全节点发送目标id
	conn, err := net.Dial("tcp", ip2+":9886")
	if err != nil {
		fmt.Println("Error connecting to server:@9886", err.Error())
	}
	defer conn.Close()

	_, err = conn.Write([]byte(target))
	if err != nil {
		fmt.Println("Error sending to server:@9886", err.Error())
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return false
	}
	response := string(buffer[:n])

	//fmt.Println("Server response:", response)

	//从服务端网页取得数据
	// 构建GET请求URL
	//url := fmt.Sprintf("http://repustation.000webhostapp.com/index.php?seller=%s", target)
	url := "http://repustation.000webhostapp.com/index.php?seller=" + target
	//fmt.Println(target, url)
	// 发送GET请求
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// 读取响应数据
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// 解析JSON响应
	var data []dataJSON
	if err := json.Unmarshal(responseBody, &data); err != nil {
		panic(err)
	}

	// 根据ID大小排序
	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	// 提取comment字段
	comments := make([]string, len(data))
	hash := ""
	for i, row := range data {
		comments[i] = row.Comment
		newHash := sha256.Sum256([]byte(hash + comments[i]))
		hash = hex.EncodeToString(newHash[:])
		//fmt.Println(hash)
	}
	if response == hash {
		return true
	} else {
		return false
	}
	// 打印结果
	//fmt.Println("Comments for seller:", target)
	//for i, comment := range comments {
	//	fmt.Printf("%d. %s\n", i+1, comment)
	//}

}
