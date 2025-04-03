package request

type FormQuestionItemOption struct {
	Key    int    `json:"key"`
	Value  string `json:"value"`
	Status int    `json:"status"`
}

type FormQuestionItem struct {
	QuestionId            int                      `json:"question_id"`
	Title                 string                   `json:"title"`
	Pos                   int                      `json:"pos"`
	Status                int                      `json:"status"`
	ReplyType             int                      `json:"reply_type"`
	MustReply             bool                     `json:"must_reply"`
	Note                  string                   `json:"note"`
	OptionItem            []FormQuestionItemOption `json:"option_item"`
	Placeholder           string                   `json:"placeholder"`
	QuestionExtendSetting interface{}              `json:"question_extend_setting"`
}

type FormQuestion struct {
	Items []FormQuestionItem `json:"items"`
}

type FormSettingFillRange struct {
	UserIds       []string `json:"userids"`
	DepartmentIds []int    `json:"departmentids"`
}

type FormSettingManagerRange struct {
	UserIds []string `json:"userids"`
}

type FormSettingTimedRepeatInfo struct {
	Enable         bool `json:"enable"`
	WeekFlag       int  `json:"week_flag"`
	RemindTime     int  `json:"remind_time"`
	RepeatType     int  `json:"repeat_type"`
	SkipHoliday    bool `json:"skip_holiday"`
	DayOfMonth     int  `json:"day_of_month"`
	ForkFinishType int  `json:"fork_finish_type"`
}

type FormSetting struct {
	FillOutAuth         int                        `json:"fill_out_auth"`
	FillInRange         FormSettingFillRange       `json:"fill_in_range"`
	SettingManagerRange FormSettingManagerRange    `json:"setting_manager_range"`
	TimedRepeatInfo     FormSettingTimedRepeatInfo `json:"timed_repeat_info"`
	AllowMultiFill      bool                       `json:"allow_multi_fill"`
	TimedFinish         int                        `json:"timed_finish"`
	CanAnonymous        bool                       `json:"can_anonymous"`
	CanNotifySubmit     bool                       `json:"can_notify_submit"`
}

type FormInfo struct {
	FormTitle    string       `json:"form_title"`
	FormDesc     string       `json:"form_desc"`
	FormHeader   string       `json:"form_header"`
	FormQuestion FormQuestion `json:"form_question"`
	FormSetting  FormSetting  `json:"form_setting"`
}

type RequestWeDocCreateForm struct {
	SpaceId  string   `json:"spaceid"`
	FatherId string   `json:"fatherid"`
	FormInfo FormInfo `json:"form_info"`
}
