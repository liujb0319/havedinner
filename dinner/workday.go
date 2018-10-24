package dinner

//
import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/json-iterator/go"
)

func CheckDay() bool {
	today := time.Now().Format("20060102")
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true}
	client := &http.Client{Transport: tr}

	url := fmt.Sprintf("http://api.k780.com:88/?app=life.workday&date=%s&appkey=10003&sign=b59bc3ef6191eb9f747dd4e83c99f2a4&format=json", today)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	resp, err := client.Do(req)
	//defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	workmk := jsoniter.Get(bytes, "result", "workmk").ToInt()
	if workmk == 1 {
		fmt.Printf("工作日：%s\n", string(bytes))
		return true
	} else {
		fmt.Printf("节假日：%s\n", string(bytes))
		return false
	}
}
