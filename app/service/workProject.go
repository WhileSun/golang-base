package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
)

type WorkProject struct {

}

func NewWorkProject() *WorkProject {
	return &WorkProject{}
}

func (s *WorkProject) checkProjectNameExist(projectName string) error {
	if id := models.NewWorkProject().CheckProjectNameExist(projectName); id > 0 {
		return errors.New(fmt.Sprintf("项目名称[%s]已经存在，请更换！", projectName))
	}
	return nil
}

func (s *WorkProject) Add(params dto.AddWorkProject,req *gin.Context) error{
	workProjectModel := models.NewWorkProject()
	gconvert.StructCopy(params, workProjectModel)
	workProjectModel.CreaterId = req.GetInt("userId")
	if err := NewWorkProject().checkProjectNameExist(workProjectModel.ProjectName); err != nil {
		return err
	}
	if err := workProjectModel.Add();err !=nil{
		gsys.Logger.Error("添加工作项目失败—>", err.Error())
		return errors.New("添加工作项目失败！")
	}
	return nil
}

func (s *WorkProject) Update(params dto.UpdateWorkProject) error{
	oldWorkProjectModel := models.NewWorkProject()
	oldWorkProjectModel.GetInfo(params.Id)

	workProjectModel := models.NewWorkProject()
	gconvert.StructCopy(params, workProjectModel)
	if oldWorkProjectModel.ProjectName != workProjectModel.ProjectName{
		if err := NewWorkProject().checkProjectNameExist(workProjectModel.ProjectName); err != nil {
			return err
		}
	}
	if err := workProjectModel.Update();err !=nil{
		gsys.Logger.Error("修改工作项目失败—>", err.Error())
		return errors.New("修改工作项目失败！")
	}
	return nil
}
