package dinner

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Person struct {
	UserName string
	Passwrd  string
	Cookie   string
}

func (p *Person) HaveDinner(user, pwd string) {
	if CheckDay() {
		p.UserName = user
		p.Passwrd = pwd
		p.Cookie = p.getCookie()
		//登录
		p.login()
		p.post()
	}
}
func (p *Person) getCookie() string {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", "http://124.127.188.104:1001/", nil)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	response, err := client.Do(reqest) //提交
	defer response.Body.Close()
	cookies := response.Cookies() //遍历cookies
	for _, cookie := range cookies {
		fmt.Println("cookie1:", cookie)
		if cookie.Name == "webpy_session_id" {
			return cookie.Value
		}
	}
	fmt.Println("state code:", response.StatusCode)
	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
	}
	fmt.Println(string(body)) //网页源码
	return ""

}

func (p *Person) login() {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", "http://124.127.188.104:1001/login", strings.NewReader(fmt.Sprintf("username=%s&passwd=%s", p.UserName, p.Passwrd)))
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.8")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Upgrade-Insecure-Requests", "1")
	reqest.Header.Add("Content-Length", "45")
	reqest.Header.Add("Referer", "http://124.127.188.104:1001/login")
	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Cookie", fmt.Sprintf("webpy_session_id=%s", p.Cookie))
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:57.0) Gecko/20100101 Firefox/57.0")
	resp, err := client.Do(reqest)
	//打印cookies
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}
	//打印状态码
	fmt.Println("state code:", resp.StatusCode)
	defer resp.Body.Close()

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("errors:", err1.Error())
	}

	fmt.Println(string(body))
}

func (p *Person) post() {
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", "http://124.127.188.104:1001/", strings.NewReader("key="))
	if err != nil {
		fmt.Println(err)
		// handle error
	}
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.8")
	reqest.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	reqest.Header.Add("Accept-Encoding", "gzip, deflate")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Upgrade-Insecure-Requests", "1")
	reqest.Header.Add("Content-Length", "45")
	reqest.Header.Add("Referer", "http://124.127.188.104:1001/")
	reqest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	reqest.Header.Add("Cookie", fmt.Sprintf("webpy_session_id=%s", p.Cookie))
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:57.0) Gecko/20100101 Firefox/57.0")
	resp, err := client.Do(reqest)
	defer resp.Body.Close()

	cookies := resp.Cookies()
	for _, cookie := range cookies {
		fmt.Println("cookie:", cookie)
	}
	fmt.Println("state code:", resp.StatusCode)

	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("errors:", err1.Error())
	}
	fmt.Println(string(body))
}
