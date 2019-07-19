package service

import (
	"errors"
	"fmt"

	"github.com/mumushuiding/util"
	"github.com/mumushuiding/workdiary/model"
)

// Project 项目
type Project struct {
	ProjectID int `json:"projectid"`
	// 一个中文长度为：3，英文为： 1
	ProjectContent string `gorm:"size:1000" json:"projectcontent"`
	UserID         int    `json:"userId"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	Progress       string `gorm:"size:1000" json:"progress"`
	PageIndex      int    `json:"pageIndex" gorm:"default:1"`
	PageSize       int    `json:"pageSize" gorm:"default:10"`
}

// Save Save
func (p *Project) Save() (int, error) {
	start, err := util.ParseDate(p.StartDate, util.YYYY_MM_DD)
	if err != nil {
		return 0, err
	}
	end, err := util.ParseDate(p.EndDate, util.YYYY_MM_DD)
	if err != nil {
		return 0, err
	}
	if start.Unix() > end.Unix() {
		return 0, errors.New("【开始日期】不能大于【结束日期】")
	}
	var entity = model.ResProject{
		ProjectContent: p.ProjectContent,
		UserID:         p.UserID,
		StartDate:      start,
		EndDate:        end,
		Progress:       p.Progress,
	}
	fmt.Println(entity.StartDate)
	err = entity.Save()
	if err != nil {
		return 0, err
	}
	return entity.ProjectID, nil
}

// FindAllPaged 分页查询
func (p *Project) FindAllPaged() (string, error) {
	datas, count, err := model.FindAllProjectPaged(p.PageIndex, p.PageSize, p.getSQL())
	if err != nil {
		return "", err
	}
	return util.ToPageJSON(datas, count, p.PageIndex, p.PageSize)
}
func (p *Project) getSQL() string {
	var sql string
	if len(p.EndDate) > 0 {
		// maps["endDate"] = p.EndDate
		sql += " and endDate <='" + p.EndDate + "'"
	}
	if len(p.StartDate) > 0 {
		// maps["startDate"] = p.StartDate
		sql += " and startDate>='" + p.StartDate + "'"
	}
	if len(p.Progress) > 0 {
		// maps["progress"] = p.Progress
		sql += " and progress like '%" + p.Progress + "%'"
	}
	if len(p.ProjectContent) > 0 {
		sql += " and projectContent like %'" + p.ProjectContent + "%'"
	}
	if p.ProjectID != 0 {
		// maps["projectId"] = p.ProjectId
		sql += " and projectId=" + fmt.Sprintf("%d", p.ProjectID)
	}
	if p.UserID != 0 {
		sql += " and userId=" + fmt.Sprintf("%d", p.UserID)
	}
	fmt.Println(sql[0:4])
	if sql[0:4] == " and" {
		return sql[5:]
	}
	return sql
}
