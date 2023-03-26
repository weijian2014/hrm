package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Employee struct {
	Id              uint64    `xml:"id" json:"id" description:"用户ID"`
	Name            string    `xml:"name" json:"name" description:"姓名"`
	Gender          string    `xml:"gender" json:"gender" description:"性别"`
	Birthday        time.Time `xml:"birthday" json:"birthday" description:"出生日期"`
	FirstWorkTime   time.Time `xml:"first_work_time" json:"first_work_time" description:"参加工作时间"`
	Salary          uint64    `xml:"salary" json:"salary" description:"工资"`
	Post            string    `xml:"post" json:"post" description:"岗位"`
	SocialSecurity  string    `xml:"social_security" json:"social_security" description:"社保"`
	Phone           string    `xml:"phone" json:"phone" description:"电话"`
	FormerEmployer  string    `xml:"former_employer" json:"former_employer" description:"原单位"`
	Height          uint64    `xml:"height" json:"height" description:"身高"`
	Weight          uint64    `xml:"weight" json:"weight" description:"体重"`
	Degree          string    `xml:"degree" json:"degree" description:"学历"`
	PoliticalStatus string    `xml:"political_status" json:"political_status" description:"政治面貌"`
	Identifier      string    `xml:"identifier" json:"identifier" description:"身份证"`
	SecurityCard    string    `xml:"security_card" json:"security_card" description:"保安证"`
	CurrentAddress  string    `xml:"current_address" json:"current_address" description:"现住址"`
	Comments        string    `xml:"comments" json:"comments" description:"备注"`
	UpdatedAt       time.Time `xml:"update_at" json:"update_at" description:"更新时间"`
}

func (e *Employee) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 表名为employees
	if err = db.AutoMigrate(e); err != nil {
		return err
	}

	return nil
}

func SelectAllEmployee() (*[]Employee, error) {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	employees := new([]Employee)
	result := db.Find(employees)
	if result.Error != nil {
		return nil, result.Error
	}

	return employees, nil
}

func (e *Employee) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(e, "name = ?", e.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (e *Employee) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 插入记录
	ret := db.Create(e)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (e *Employee) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(e).Update("name", e.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (e *Employee) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Unscoped().Where("id = ?", e.Id).Delete(e)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (e *Employee) BeforeDelete(tx *gorm.DB) error {
	return nil
}
