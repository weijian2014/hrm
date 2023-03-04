package api

import (
	"github.com/xuri/excelize/v2"
)

const (
	excelSheetName      = "Sheet1"
	nameCell            = "B2" // 姓名
	genderCell          = "D2" // 性别
	ageCell             = "F2" // 年龄
	workTimeCell        = "B3" // 工作时间
	salaryCell          = "D3" // 工资
	postCell            = "F3" // 岗位
	socialSecurityCell  = "B4" // 社保
	phoneCell           = "D4" // 电话
	formerEmployerCell  = "F4" // 原单位
	heightCell          = "B5" // 身高
	weightCell          = "D5" // 体重
	diplomaCell         = "F5" // 文化
	politicalStatusCell = "B6" // 政治面貌
	idCell              = "D6" // 身份证
	securityCardCell    = "F6" // 保安证
	currentAddressCell  = "B7" // 现住址
	commentsCell        = "A8" // 需要了解的情况
)

func init() {

}

type EmployeeInfo struct {
	Name            string // 姓名
	Gender          string // 性别
	Age             string // 年龄
	WorkTime        string // 工作时间
	Salary          string // 工资
	Post            string // 岗位
	SocialSecurity  string // 社保
	Phone           string // 电话
	FormerEmployer  string // 原单位
	Height          string // 身高
	Weight          string // 体重
	Diploma         string // 文化
	PoliticalStatus string // 政治面貌
	Id              string // 身份证
	SecurityCard    string // 保安证
	CurrentAddress  string // 现住址
	Comments        string // 需要了解的情况
}

func ReadExcel(excelFile string) (*EmployeeInfo, error) {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	e := &EmployeeInfo{}

	// 姓名
	e.Name, err = f.GetCellValue(excelSheetName, nameCell)
	if err != nil {
		return nil, err
	}

	// 性别
	e.Gender, err = f.GetCellValue(excelSheetName, genderCell)
	if err != nil {
		return nil, err
	}

	// 年龄
	e.Age, err = f.GetCellValue(excelSheetName, ageCell)
	if err != nil {
		return nil, err
	}

	// 工作时间
	e.WorkTime, err = f.GetCellValue(excelSheetName, workTimeCell)
	if err != nil {
		return nil, err
	}

	// 工资
	e.Salary, err = f.GetCellValue(excelSheetName, salaryCell)
	if err != nil {
		return nil, err
	}

	// 岗位
	e.Post, err = f.GetCellValue(excelSheetName, postCell)
	if err != nil {
		return nil, err
	}

	// 社保
	e.SocialSecurity, err = f.GetCellValue(excelSheetName, socialSecurityCell)
	if err != nil {
		return nil, err
	}

	// 电话
	e.Phone, err = f.GetCellValue(excelSheetName, phoneCell)
	if err != nil {
		return nil, err
	}

	// 原单位
	e.FormerEmployer, err = f.GetCellValue(excelSheetName, formerEmployerCell)
	if err != nil {
		return nil, err
	}

	// 身高
	e.Height, err = f.GetCellValue(excelSheetName, heightCell)
	if err != nil {
		return nil, err
	}

	// 体重
	e.Weight, err = f.GetCellValue(excelSheetName, weightCell)
	if err != nil {
		return nil, err
	}

	// 文化
	e.Diploma, err = f.GetCellValue(excelSheetName, diplomaCell)
	if err != nil {
		return nil, err
	}

	// 政治面貌
	e.PoliticalStatus, err = f.GetCellValue(excelSheetName, politicalStatusCell)
	if err != nil {
		return nil, err
	}

	// 身份证
	e.Id, err = f.GetCellValue(excelSheetName, idCell)
	if err != nil {
		return nil, err
	}

	// 保安证
	e.SecurityCard, err = f.GetCellValue(excelSheetName, securityCardCell)
	if err != nil {
		return nil, err
	}

	// 现住址
	e.CurrentAddress, err = f.GetCellValue(excelSheetName, currentAddressCell)
	if err != nil {
		return nil, err
	}

	// 需要了解的情况
	e.Comments, err = f.GetCellValue(excelSheetName, commentsCell)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func WriteExcel(excelFile string, ei *EmployeeInfo) error {
	f, err := excelize.OpenFile(excelFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// 姓名
	err = f.SetCellValue(excelSheetName, nameCell, ei.Name)
	if err != nil {
		return err
	}

	// 性别
	err = f.SetCellValue(excelSheetName, genderCell, ei.Gender)
	if err != nil {
		return err
	}

	// 年龄
	err = f.SetCellValue(excelSheetName, ageCell, ei.Age)
	if err != nil {
		return err
	}

	// 工作时间
	err = f.SetCellValue(excelSheetName, workTimeCell, ei.WorkTime)
	if err != nil {
		return err
	}

	// 工资
	err = f.SetCellValue(excelSheetName, salaryCell, ei.Salary)
	if err != nil {
		return err
	}

	// 岗位
	err = f.SetCellValue(excelSheetName, postCell, ei.Post)
	if err != nil {
		return err
	}

	// 社保
	err = f.SetCellValue(excelSheetName, socialSecurityCell, ei.SocialSecurity)
	if err != nil {
		return err
	}

	// 电话
	err = f.SetCellValue(excelSheetName, phoneCell, ei.Phone)
	if err != nil {
		return err
	}

	// 原单位
	err = f.SetCellValue(excelSheetName, formerEmployerCell, ei.FormerEmployer)
	if err != nil {
		return err
	}

	// 身高
	err = f.SetCellValue(excelSheetName, heightCell, ei.Height)
	if err != nil {
		return err
	}

	// 体重
	err = f.SetCellValue(excelSheetName, weightCell, ei.Weight)
	if err != nil {
		return err
	}

	// 文化
	err = f.SetCellValue(excelSheetName, diplomaCell, ei.Diploma)
	if err != nil {
		return err
	}

	// 政治面貌
	err = f.SetCellValue(excelSheetName, politicalStatusCell, ei.PoliticalStatus)
	if err != nil {
		return err
	}

	// 身份证
	err = f.SetCellValue(excelSheetName, idCell, ei.Id)
	if err != nil {
		return err
	}

	// 保安证
	err = f.SetCellValue(excelSheetName, securityCardCell, ei.SecurityCard)
	if err != nil {
		return err
	}

	// 现住址
	err = f.SetCellValue(excelSheetName, currentAddressCell, ei.CurrentAddress)
	if err != nil {
		return err
	}

	// 需要了解的情况
	err = f.SetCellValue(excelSheetName, commentsCell, ei.Comments)
	if err != nil {
		return err
	}

	err = f.Save()
	if err != nil {
		return err
	}

	return nil
}
