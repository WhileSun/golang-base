package service

import (
	"errors"
	"github.com/whilesun/go-admin/app/models"
	dto2 "github.com/whilesun/go-admin/app/types/dto"
	gsys2 "github.com/whilesun/go-admin/gctx"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtime"
)

type WorkTask struct {
}

func NewWorkTask() *WorkTask {
	return &WorkTask{}
}

func (s *WorkTask) checkTimeOrder(launchTimeStr string, startTimeStr string, endTimeStr string, workTaskModel *models.WorkTask) error {
	launchTime := gconvert.StrToDatetime(launchTimeStr)
	startTime := gconvert.StrToDatetime(startTimeStr)
	endTime := gconvert.StrToDatetime(endTimeStr)
	if !startTime.IsZero() && startTime.Before(launchTime) {
		return errors.New("开始时间不能发起时间之前")
	}
	if !endTime.IsZero() && endTime.Before(startTime) {
		return errors.New("结束时间不得大小于开启时间")
	}
	workTaskModel.LaunchTime = gtime.DateTime{Time: launchTime}
	if !startTime.IsZero() {
		workTaskModel.StartTime = gtime.DateTime{Time: startTime}
	}
	if !endTime.IsZero() {
		workTaskModel.EndTime = gtime.DateTime{Time: endTime}
	}
	return nil
}

func (s *WorkTask) Add(params dto2.AddWorkTask) error {
	workTaskModel := models.NewWorkTask()
	gconvert.StructCopy(params, workTaskModel)
	if err := NewWorkTask().checkTimeOrder(params.LaunchTimeStr, params.StartTimeStr, params.EndTimeStr, workTaskModel); err != nil {
		return err
	}
	if err := workTaskModel.Add(); err != nil {
		gsys2.Logger.Error("添加工作项目任务失败—>", err.Error())
		return errors.New("添加工作项目任务失败！")
	}
	return nil
}

func (s *WorkTask) Update(params dto2.UpdateWorkTask) error {
	workTaskModel := models.NewWorkTask()
	gconvert.StructCopy(params, workTaskModel)
	if err := NewWorkTask().checkTimeOrder(params.LaunchTimeStr, params.StartTimeStr, params.EndTimeStr, workTaskModel); err != nil {
		return err
	}
	if err := workTaskModel.Update(); err != nil {
		gsys2.Logger.Error("修改工作项目任务失败—>", err.Error())
		return errors.New("修改工作项目任务失败！")
	}
	return nil
}
