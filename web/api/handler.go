package main

import (
	"github.com/gin-gonic/gin"
	"../../utils/common"
	"../../utils/protocol"
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
	"html/template"
	"fmt"
)

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
func geneSysDetails(item *protocol.SysPackage) template.HTML{
	ret := ""
	for _,v := range item.CourseDetail{
		//
		left := v.StudentTotal-v.ApplyNum
		if left <= 0 {
			left = 1
		}
		ret += fmt.Sprintf(sysHtml,v.Name,v.TimePlan,tinfo(v.TeList),left,v.AfAmount/100)
	}
	return template.HTML(ret)
}


func tinfo(item []protocol.Teacher) string{
	ret := ""
	for _,v := range item{
		ret += fmt.Sprintf(teacherHtml,v.CoverUrl,v.Name)
	}
	return ret
}

func geneTeacherInfo(item *protocol.SpcCourse) template.HTML{
	//ret := ""
	//for _,v := range item.TeList{
	//	ret += fmt.Sprintf(teacherHtml,v.CoverUrl,v.Name)
	//}
	return template.HTML(tinfo(item.TeList))
}

func initRouter(r *gin.Engine){
	r.SetFuncMap(template.FuncMap{"gteacher":geneTeacherInfo,"gsys":geneSysDetails})
	//r.SetFuncMap(template.FuncMap{"gsys":geneSysDetails})
	r.LoadHTMLFiles("../static/fd.html")
	r.GET("/", query)
}

func filterPrice(item *protocol.CourseInfo) {
	//price
	for i,_ := range item.SpeCourse{
		item.SpeCourse[i].PreAmonut /= 100
	}

	for i,_ := range item.HotCourse{
		item.HotCourse[i].PreAmonut /= 100
	}

}

func query(c *gin.Context){
	//query_srvs
	//get_query_str  redis_stroed
	grade := c.DefaultQuery("grade","2")
	bs := common.QueryByGrade(grade)
	retObj := protocol.CourseInfo{}
	//log.Info("arg ",grade,bs)
	if bs !=nil {
		err := json.Unmarshal(bs,&retObj)
		if err != nil {
			log.Error(err)
		}else{
			retObj.LenHot = len(retObj.HotCourse)
			retObj.LenSys = len(retObj.SysPackage)
			retObj.LenSpe = len(retObj.SpeCourse)
			filterPrice(&retObj)
		}
	}
	log.Infof("qres %v\n",retObj)
	c.HTML(http.StatusOK,"fd.html",retObj)
	//page_loads floats
}

