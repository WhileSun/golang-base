package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/whilesun/go-admin/pkg/e"
	"github.com/whilesun/go-admin/pkg/extend/dbtypes"
)

type IndexController struct {

}
type Test struct {
	EndTime   dbtypes.GTime  `form:"end_time"  binding:"required" label:"结束时间"`
}
type TestDate struct {
	Id        int              `gorm:"primary_key"  json:"id"`
	EndTime  dbtypes.GTime `json:"end_time"`
}

func (c *IndexController) Test(req *gin.Context){
	//var params Test
	//if err:= gvalidator.ReqValidate(req,&params);err!=nil{
	//	logger.Info("添加角色参数有误->",err.Error())
	//	e.New(req).MsgDetail(e.API_PARAMS_FAIL,err.Error())
	//	return
	//}
	//fmt.Println(carbon.Now().Format("Y-m-d H:i:s"))
	//fmt.Printf("%+v",params)
	//carbon.NewCarbon().SetTimezone(carbon.PRC)
	//testdata := &TestDate{}
	//testdata.EndTime = carbon.DateTime{carbon.Parse(params.EndTime)}
	////utils.StructCopy(params,testdata,"time.Time")
	////fmt.Printf("%+v",testdata)
	//gsys.Db.Create(testdata)
	//id := testdata.Id
	//id := 8
	//testdata1 := &TestDate{}
	//gsys.Db.Table("test_date").Where("id = ?",id).Scan(testdata1)
	//fmt.Printf("%+v",testdata1)
	//e.New(req).Data(e.SUCCESS,"",testdata1)
	e.New(req).Msg(e.SUCCESS)
}