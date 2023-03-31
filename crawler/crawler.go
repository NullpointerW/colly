package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"
)

type T struct {
	Tr []struct {
		Td    []string `json:"td"`
		Align string   `json:"@align,omitempty"`
		Style string   `json:"@style,omitempty"`
	} `json:"tr"`
}

var client = &http.Client{
	Timeout: 2 * time.Second}

const (
	// 正则表达式，匹配出 XL 的账号密码
	reAccount = `(账号|迅雷账号)(；|：)[0-9:]+(| )密码：[0-9a-zA-Z]+`
)

// GetAccountAndPwd 获取网站 账号密码
func GetAccountAndPwd(url string) {
	// 获取网站数据
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	//resp, err := http.Get(url)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("http.Get error : ", err)
	}
	defer resp.Body.Close()

	// 去读数据内容为 bytes
	dataBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error : ", err)
	}

	// 字节数组 转换成 字符串
	str := string(dataBytes)

	log.Println(str)

	// 过滤 XL 账号和密码
	re := regexp.MustCompile(reAccount)

	// 匹配多少次， -1 默认是全部
	results := re.FindAllStringSubmatch(str, -1)

	log.Println(len(results))
	// 输出结果
	for _, result := range results {
		log.Println(result[0])
	}
}

// func main() {
// 	// 简单设置log 参数
// 	log.SetFlags(log.Lshortfile | log.LstdFlags)
// 	// 传入网站地址，爬取开始爬取数据
// 	GetAccountAndPwd("https://www.ucbug.com/jiaocheng/63149.html?_t=1582307696")
// }
