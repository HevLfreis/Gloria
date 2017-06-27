package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"net/url"
	"encoding/json"
)

//type BusInfo struct {
//	StopDistance string `json:"stopdis"`
//	Time string `json:"time"`
//}


func tests() {
	client := &http.Client{}

	form := url.Values{}
	form.Add("stoptype", "0")
	form.Add("stopid", "31.")
	form.Add("sid", "f7546535b8400fe607f73675f625de80")


	req, err := http.NewRequest("POST",
		"http://shanghaicity.openservice.kankanews.com/public/bus/Getstop",
		strings.NewReader(form.Encode()))

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "Hm_1vt_6f69830ae7173059e935b61372431b35=eSgsNFk/PumVYX9PERT0Ag==; Hm_lvt_6f69830ae7173059e935b61372431b35=1496394505,1497317097,1497317617; Hm_lpvt_6f69830ae7173059e935b61372431b35=1497317647; HH=4c8cebc7b78a1342526e71d16ebda3763a6d6edc; HK=de612c61ac430e7a1551c611c55832587741acc1; HG=7053f8f0c650d7822920182fd723fef6a56c2116; HA=8688bb3745d762dc8a25d7eb9c226821ab391e3f; HB=ODY4OGJiMzc0NWQ3NjJkYzhhMjVkN2ViOWMyMjY4MjFhYjM5MWUzZg==; HC=7bcae0bf36c4791c2fdf273b6e80288c2f9de551; HD=MjAxNzA2MTM=; HY=MjAxNzA2MTM=de612c61ac430e7a1551c611c55832587741acc17053f8f0c650d7822920182fd723fef6a56c2116027dda1b6d0c64ac9728871d72d4efdc146d1b81; HO=TWpBeE56QTJNVE09MDlNVEUyTWpRMzQ2VFc5NmFXeHNZUzgxTGpBZ0tFeHBiblY0T3lCQmJtUnliMmxrSURRdU5DNDBPeUJJVFNCT1QxUkZJREZNVkVWWElFSjFhV3hrTDB0VVZUZzBVQ2tnUVhCd2JHVlhaV0pMYVhRdk5UTTNMak0ySUNoTFNGUk5UQ3dnYkdsclpTQkhaV05yYnlrZ1ZtVnljMmx2Ymk4MExqQWdRMmh5YjIxbEx6TXpMakF1TUM0d0lFMXZZbWxzWlNCVFlXWmhjbWt2TlRNM0xqTTJJRTFwWTNKdlRXVnpjMlZ1WjJWeUx6WXVNQzR3TGpVMFgzSTRORGt3TmpNdU5UQXhJRTVsZEZSNWNHVXZWMGxHU1E9PTAyN2RkYTFiNmQwYzY0YWM5NzI4ODcxZDcyZDRlZmRjMTQ2ZDFiODE=; Hm_p1vt_6f69830ae7173059e935b61372431b35=eSgsNFk/UAOVo39REMVoAg==; _ga=GA1.2.947300797.1496394508; _gat=1")
	req.Header.Set("Host", "shanghaicity.openservice.kankanews.com")
	req.Header.Set("Origin", "http://shanghaicity.openservice.kankanews.com")
	req.Header.Set("Referer", "http://shanghaicity.openservice.kankanews.com/public/bus/mes/sid/f7546535b8400fe607f73675f625de80")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 4.4.4; HM NOTE 1LTEW Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36 MicroMessenger/6.0.0.54_r849063.501 NetType/WIFI")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
	var infos []BusInfo
	json.Unmarshal([]byte(body), &infos)
	fmt.Println(infos)
}
