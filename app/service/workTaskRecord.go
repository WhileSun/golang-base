package service

import (
	"errors"
	"github.com/whilesun/go-admin/app/dto"
	"github.com/whilesun/go-admin/app/models"
	"github.com/whilesun/go-admin/pkg/gsys"
	"github.com/whilesun/go-admin/pkg/utils/gconvert"
	"github.com/whilesun/go-admin/pkg/utils/gtime"
)

type WorkTaskRecordService struct {

}

func NewWorkTaskRecord() *WorkTaskRecordService {
	return &WorkTaskRecordService{}
}

func checkTimeOrder(launchTimeStr string,startTimeStr string,endTimeStr string,workTaskRecordModel *models.WorkTaskRecord) error{
	launchTime := gconvert.StrToDatetime(launchTimeStr)
	startTime := gconvert.StrToDatetime(startTimeStr)
	endTime := gconvert.StrToDatetime(endTimeStr)
	if !startTime.IsZero() && startTime.Before(launchTime){
		return errors.New("开始时间不能发起时间之前")
	}
	if !endTime.IsZero() && endTime.Before(startTime){
		return errors.New("结束时间不得大小于开启时间")
	}
	workTaskRecordModel.LaunchTime = gtime.DateTime{Time: launchTime}
	if !startTime.IsZero(){
		workTaskRecordModel.StartTime = gtime.DateTime{Time: startTime}
	}
	if !endTime.IsZero(){
		workTaskRecordModel.EndTime = gtime.DateTime{Time: endTime}
	}
	return nil
}

func (s *WorkTaskRecordService) Add(params dto.AddWorkTaskRecord) error{
	workTaskRecordModel := models.NewWorkTaskRecord()
	gconvert.StructCopy(params, workTaskRecordModel)
	if err := checkTimeOrder(params.LaunchTimeStr,params.StartTimeStr,params.EndTimeStr,workTaskRecordModel);err!=nil{
		return err
	}
	if err := workTaskRecordModel.Add();err !=nil{
		gsys.Logger.Error("添加工作项目任务失败—>", err.Error())
		return errors.New("添加工作项目任务失败！")
	}
	return nil
}

func (s *WorkTaskRecordService) Update(params dto.UpdateWorkTaskRecord) error{
	workTaskRecordModel := models.NewWorkTaskRecord()
	gconvert.StructCopy(params, workTaskRecordModel)
	if err := checkTimeOrder(params.LaunchTimeStr,params.StartTimeStr,params.EndTimeStr,workTaskRecordModel);err!=nil{
		return err
	}
	if err := workTaskRecordModel.Update();err !=nil{
		gsys.Logger.Error("修改工作项目任务失败—>", err.Error())
		return errors.New("修改工作项目任务失败！")
	}
	return nil
}