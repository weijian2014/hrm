package db

import (
	"time"
)

type Employee struct {
	Id              uint64    `json:"id" description:"用户ID"`
	Name            string    `json:"name" description:"姓名"`
	Gender          string    `json:"gender" description:"性别"`
	Birthday        time.Time `json:"birthday" description:"出生日期"`
	Height          uint64    `json:"height" description:"身高"`
	Weight          uint64    `json:"weight" description:"体重"`
	Degree          string    `json:"degree" description:"学历"`
	Identifier      string    `json:"identifier" description:"身份证号"`
	Phone           string    `json:"phone" description:"手机号码"`
	PoliticalStatus string    `json:"political_status" description:"政治面貌"`
	SocialSecurity  string    `json:"social_security" description:"社保"`
	CurrentAddress  string    `json:"current_address" description:"现住址"`
	FirstWorkTime   time.Time `json:"first_work_time" description:"首次工作"`
	FormerEmployer  string    `json:"former_employer" description:"原单位"`
	Post            string    `json:"post" description:"岗位"`
	Salary          uint64    `json:"salary" description:"工资"`
	SecurityCard    string    `json:"security_card" description:"保安证"`
	Comments        string    `json:"comments" description:"备注"`
	UpdatedAt       time.Time `json:"updated_at" description:"更新时间"`
}
