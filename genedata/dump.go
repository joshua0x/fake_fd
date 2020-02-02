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

/*
   "cid": 182981,
   "name": "人工智能第一课——北大博士带孩子走进生活中的科学",
   "pinyin": "",
   "cover_url": "//pub.idqqimg.com/pc/misc/files/20200106/0c173830638842c588046a61ae0adc02.png",
   "showid": 6,
   "payid": 2,
   "grade": 1984,
   "subject": 7058,
   "status": 1,
   "recordtime": 1578312730,
   "cover_url_color": "//pub.idqqimg.com/pc/misc/files/20200106/0c173830638842c588046a61ae0adc02.png",
   "time_plan": "2月3日-2月7日 19:00 5节",
   "hint_logo": "",
   "class_type": 0,
   "student_total": 0,
   "level": 0,
   "course_labels": [
     {
       "label_option_id": 4,
       "label_option_name": "20",
       "label_option_type": 1
     },
     {
       "label_option_id": 11,
       "label_option_name": "寒",
       "label_option_type": 2
     },
     {
       "label_option_id": 7003,
       "label_option_name": "三年级",
       "label_option_type": 5
     },
     {
       "label_option_id": 7058,
       "label_option_name": "科学课",
       "label_option_type": 4
     },
     {
       "label_option_id": 195262,
       "label_option_name": "",
       "label_option_type": 8
     },
     {
       "label_option_id": 195263,
       "label_option_name": "",
       "label_option_type": 9
     }
   ],
   "has_discount": 0,
   "pre_amount": 990,
   "af_amount": 990,
   "dis_bgtime": 0,
   "dis_edtime": 0,
   "sign_max": 1000000,
   "apply_num": 555,
   "sign_bgtime": 1577980800,
   "sign_end_time": 1581091199,
   "first_sub_bgtime": 1580727600,
   "first_sub_endtime": 1580731200,
   "last_sub_bgtime": 1581073200,
   "last_sub_endtime": 1581078600,
*/

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
	//	user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36
	//	x-request-id: c8f13ae3-33fe-4c3b-9cb6-1661fb0e97e2
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
