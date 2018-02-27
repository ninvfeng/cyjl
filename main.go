package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//用户ID 每次启动随机生成
var userid = strconv.Itoa(rand.New(rand.NewSource(time.Now().Unix())).Intn(100))

//图灵key 用户输入
var key = ""

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("-请输入图灵机器人key开始使用  -输入 exit 退出")

	for true {
		data, _, _ := reader.ReadLine()
		info := string(data)
		if info == "exit" {
			os.Exit(0)
		}
		if strings.Count(info, "")-1 == 32 {
			key = info
			tuling("成语接龙")
		} else if strings.Count(key, "")-1 != 32 {
			fmt.Println("请输入图灵机器人key")
		} else {
			tuling(info)
		}
	}
}

//图灵机器人
func tuling(info string) {

	//接口返回数据结构
	type tulingRes struct {
		Code int
		Text string
	}

	//请求接口
	data, err := post("http://www.tuling123.com/openapi/api", "key="+key+"&userid="+userid+"&info="+info)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//处理返回
	Res := new(tulingRes)
	json.Unmarshal([]byte(data), &Res)
	fmt.Println(Res.Text)

}

//发送post请求
func post(url, data string) (string, error) {
	res, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
