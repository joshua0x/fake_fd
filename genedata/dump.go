package main

import (
	//"bytes"
	"io/ioutil"
	"log"
	"net/http"
	//"encoding/json"
	"../utils/common"
	"../utils/protocol"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"strconv"
)


//api

func main() {
	//xiaox
	//cz 7002
	//gz
	client := &http.Client{}
	//valueType:="application/json"
	//referer: https://fudao.qq.com/grade/7002/
	//	sec-fetch-mode: cors
	//	sec-fetch-site: same-origin

	referer := "https://fudao.qq.com/grade/%v/"
	headerMap := map[string]string{
		"referer": "https://fudao.qq.com/grade/%v/",
	}
	url := "https://fudao.qq.com/cgi-proxy/course/index_pc_discover?client=4&platform=3&version=30&grade=%v&t=0.07704148927558774"
	maxGrade := 6 + 3 + 3
	base := 0
	for i := 1; i <= maxGrade; i++ {
		if i <= 6 {
			base = 7000 + i
		} else if i <= 9 {
			base = 6000 + i - 6
		} else {
			base = 5000 + i - 9
		}
		eUrl := fmt.Sprintf(url, base)
		refe := fmt.Sprintf(referer, base)
		headerMap["referer"] = refe
		request, _ := http.NewRequest("POST", eUrl, nil)
		for k, v := range headerMap {
			request.Header.Set(k, v)
		}
		response, _ := client.Do(request)
		if response.StatusCode == 200 {
			body, _ := ioutil.ReadAll(response.Body)
			//json loads  json fmts
			//saved@mysql  __
			obj := protocol.CourseInfo{}
			va, _, _, _ := jsonparser.Get(body, "result")
			err := json.Unmarshal(va, &obj)
			if err != nil {
				log.Println(err)
			}
			log.Printf("%+v", obj.SysPackage[0])
			err = common.GeneData(strconv.Itoa(i), va)
			if err != nil {
				log.Println(err)
			}
		}
	}

}
