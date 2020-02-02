package protocol

type CourseInfo struct {
	//maps
	SysPackage []SysPackage `json:"sys_package"`
	SpeCourse  []SpcCourse  `json:"spe_course"`
	HotCourse  []SpcCourse  `json:"hot_course"`
	LenSys,LenSpe,LenHot int
}


type SysPackage struct {
	Title             string `json:"title"`
	CourseBeginTime   int    `json:"course_bgtime"`
	CourseEndTime     int    `json:"course_endtime"`
	SoldCount         int    `json:"sold_count"`
	CourseSignEndTime int    `json:"course_sign_end_time"`
	CourseDetail []Course `json:"course_info"`
}

type Course struct {
	Cid int `json:"cid"`
	Name string `json:"name"`
	TimePlan string `json:"time_plan"`
	ApplyNum int `json:"apply_num"`
	StudentTotal int `json:"student_total"`  // sub get remains
	ClassType int `json:"class_type"`  // 数学 物理
	TeList []Teacher `json:"te_list"`
	AfAmount int `json:"af_amount"`
}

type Teacher struct {
	Name string `json:"name"`
	CoverUrl string `json:"cover_url"`

}

type SpcCourse struct {
	Name      string `json:"name"`
	CoverUrl  string `json:"cover_url"`
	TimePlan  string `json:"time_plan"`
	PreAmonut float32 `json:"pre_amount"`
	ApplyNum  int    `json:"apply_num"`
	TeList []Teacher `json:"te_list"`
	//CourseInfo []Course `json:"course_info"`
}

type HotCourse struct {
	Name      string `json:"name"`
	CoverUrl  string `json:"cover_url"`
	TimePlan  string `json:"time_plan"`
	PreAmonut int `json:"pre_amount"`
	//CourseInfo []Course `json:"course_info"`
	ApplyNum  int    `json:"apply_num"`
	TeList []Teacher `json:"te_list"`
}


