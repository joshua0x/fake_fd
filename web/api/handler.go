package main

import (
	"../../utils/common"
	"../../utils/protocol"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"html/template"
	"net/http"
	"strconv"
)

var gradeList = []string{"一年级", "二年级", "三年级", "四年级", "五年级", "六年级", "初一", "初二", "初三", "高一", "高二", "高三"}

const sysHtml = `  <li class="course-card"><a data-modid="sys_course_card" data-tdw="{&quot;position&quot;:1,&quot;course_id&quot;:136045}" href="/pc/course.html?course_id=136045" target="_blank" rel="noopener noreferrer" has-expose="1">
                            <div class="course-title--wrapper">
                                <h5 class="course-title"><span>%s</span></h5>
                                <p class="course-times">%s</p>
                            </div>
                            <div class="course-teacher">
                                %s
                            </div>
                            <div class="course-sell">
                                <div class="course-hint">
                                    <span>剩%v个名额</span>
                                </div>
                                <div class="course-price">
                                    <span class="course-price--yen">&yen;</span>
                                    <span class="course-price--cost">%v</span>
                                </div>
                            </div></a></li>
`
const teacherHtml string = `<div class="teacher-wrapper">
                                                <img alt="老师头像" class="teacher-avatar" src=%s />
                                                <div class="teacher-info">
                                                    <p class="teacher-name">%s</p>
                                                    <p class="teacher-type">授课</p>
                                                </div>
                                            </div>`

func geneSysDetails(item *protocol.SysPackage) template.HTML {
	ret := ""
	for _, v := range item.CourseDetail {
		//
		left := v.StudentTotal - v.ApplyNum
		if left <= 0 {
			left = 1
		}
		ret += fmt.Sprintf(sysHtml, v.Name, v.TimePlan, tinfo(v.TeList), left, v.AfAmount/100)
	}
	return template.HTML(ret)
}

func tinfo(item []protocol.Teacher) string {
	ret := ""
	for _, v := range item {
		ret += fmt.Sprintf(teacherHtml, v.CoverUrl, v.Name)
	}
	return ret
}

func geneTeacherInfo(item *protocol.SpcCourse) template.HTML {
	//ret := ""
	//for _,v := range item.TeList{
	//	ret += fmt.Sprintf(teacherHtml,v.CoverUrl,v.Name)
	//}
	return template.HTML(tinfo(item.TeList))
}

func initRouter(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{"gteacher": geneTeacherInfo, "gsys": geneSysDetails})
	//r.SetFuncMap(template.FuncMap{"gsys":geneSysDetails})
	r.LoadHTMLFiles("../static/fd.html")
	r.GET("/", query)
}

func filterPrice(item *protocol.CourseInfo) {
	//price
	for i, _ := range item.SpeCourse {
		item.SpeCourse[i].PreAmonut /= 100
	}

	for i, _ := range item.HotCourse {
		item.HotCourse[i].PreAmonut /= 100
	}
}

func addGrade(grade int, item *protocol.CourseInfo) {
	if grade < 1 || grade > len(gradeList) {
		grade = 1
	}
	item.Grade = gradeList[grade-1]
}

//maps seqs

func query(c *gin.Context) {
	//query_srvs
	//get_query_str  redis_stroed
	grade := c.DefaultQuery("grade", "2")
	gradeInt, err := strconv.Atoi(grade)
	if err != nil || !(gradeInt >= 1 && gradeInt <= 12) {
		log.Info("invalid_arg")
		gradeInt = 1
		grade = "2"
	}

	bs := common.QueryByGrade(grade)
	retObj := protocol.CourseInfo{}
	if bs != nil {
		err := json.Unmarshal(bs, &retObj)
		if err != nil {
			log.Error(err)
		} else {
			retObj.LenHot = len(retObj.HotCourse)
			retObj.LenSys = len(retObj.SysPackage)
			retObj.LenSpe = len(retObj.SpeCourse)
			filterPrice(&retObj)
		}
	}
	//maps datas grade
	addGrade(gradeInt, &retObj)
	log.Debugf("qres %v\n", retObj)
	c.HTML(http.StatusOK, "fd.html", retObj)
	//page_loads floats
}
