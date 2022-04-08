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

type WorkTaskProjectService struct {

}

func NewWorkTaskProject() *WorkTaskProjectService {
	return &WorkTaskProjectService{}
}

func checkProjectNameExist(workTaskProjectModel *models.WorkTaskProject) error {
	if id := workTaskProjectModel.CheckProjectNameExist(); id > 0 {
		return errors.New(fmt.Sprintf("项目名称[%s]已经存在，请更换！", workTaskProjectModel.ProjectName))
	}
	return nil
}

func (s *WorkTaskProjectService) Add(params dto.AddWorkTaskProject,req *gin.Context) error{
	workTaskProjectModel := models.NewWorkTaskProject()
	gconvert.StructCopy(params, workTaskProjectModel)
	workTaskProjectModel.CreaterId = req.GetInt("userId")
	if err := checkProjectNameExist(workTaskProjectModel); err != nil {
		return err
	}
	if err := workTaskProjectModel.Add();err !=nil{
		gsys.Logger.Error("添加工作项目失败—>", err.Error())
		return errors.New("添加工作项目失败！")
	}
	return nil
}

func (s *WorkTaskProjectService) Update(params dto.UpdateWorkTaskProject) error{
	oldWorkTaskProjectModel := models.NewWorkTaskProject()
	oldWorkTaskProjectModel.GetRowById(params.Id)

	workTaskProjectModel := models.NewWorkTaskProject()
	gconvert.StructCopy(params, workTaskProjectModel)
	if oldWorkTaskProjectModel.ProjectName != workTaskProjectModel.ProjectName{
		if err := checkProjectNameExist(workTaskProjectModel); err != nil {
			return err
		}
	}
	if err := workTaskProjectModel.Update();err !=nil{
		gsys.Logger.Error("修改工作项目失败—>", err.Error())
		return errors.New("修改工作项目失败！")
	}
	return nil
}
