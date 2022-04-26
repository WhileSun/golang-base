import React, { useState} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import WsButton from '@/components/WsButton'
import {getData,arrTransName,arrTransObj} from '@/utils/tools';
import { Space} from 'antd';
import {workTaskRecordLevelFunc,workTaskRecordStatusFunc} from "@/module/colorfunc";

var store = {};
const Index = (props) => {
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [formData,setFormData] = useState({});
  const [formShow,setFormShow] = useState(false);
  const TaskTypeName = {1:"新增功能",2:"修复BUG",3:"优化代码",4:"功能微调"};
  const performStatusName = {1:"未开始",2:"已完成",3:"进行中",4:"挂起",5:"测试中"}
  const taskLevelName = workTaskRecordLevelFunc({1:"普通",2:"紧急",3:"非常紧急"})
  const [projectSerachSelect,seProjectSerachSelect] = useState([]);

  const formFunc = (row)=>{
    setFormData(row);
    getData('load/work/taskProject/list/get',{},(data)=>{
      seProjectSerachSelect(arrTransName(data,{project_name:'label',id:'value'}));
      setFormShow(true);
    });
  }

  return (
    <>
      <WsTable
        store={store}
        table = {tableRef}
        searchs={
          [
            {type:'selectInput',listData:{'xm*project_name':'项目名称'}},
            {type:'selectInput',listData:{'rw*title':'任务标题',"rw*content":"任务内容","rw*remark":"备注"}},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{formFunc({})}}
          ]
        }
        th={[
          {title:'操作',name:'id',width:60,align:'center',render:function(v,row){
            return (<Space>
              <WsButton title="编辑" onClick={()=>{formFunc(row)}}/>
            </Space>);
          }},
          {name:"project_name",title:'项目名称',width:120,render:v=>{return v||'-'}},
          {name:"title",title:'任务标题',width:200,render:v=>{return v||'-'}},
          {name:"launch_time",title:'发起时间',align:'center',width:120,render:v=>{return v||'-'}},
          {name:"task_level",title:'优先级',align:'center',width:80,render:v=>{return taskLevelName[v]||'-'}},
          {name:"task_type",title:'任务类型',align:'center',width:100,render:v=>{return TaskTypeName[v]||'-'}},
          {name:"perform_status",title:'执行状态',align:'center',width:80,render:v=>{return workTaskRecordStatusFunc(v,performStatusName)||'-'}},
          {name:"start_time",title:'开始时间',align:'center',width:120,render:v=>{return v||'-'}},
          {name:"end_time",title:'结束时间',align:'center',width:120,render:v=>{return v||'-'}},
          {name:"remark",title:'备注',width:250,render:v=>{return v||'-'}},
          {name:"created_at",title:'创建时间',width:120,align:'center',render:v=>{return v||'-'}},
          {title:'操作',name:'id',width:60,align:'center',align:'left',render:function(v,row){
            return (<Space>
              <WsButton title="删除" pop={true} onClick={()=>{
                getData('work/taskRecord/delete',{id:row.id},()=>{
                  tableRef.reload();
                },true);
              }}/>
            </Space>);
          }}
        ]}
        api="work/taskRecord/list/get"
      />
    {formShow&&<WsForm
        form={formRef}
        width={800}
        modelFull = {true}
        title="项目任务"
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"task_project_id",col:24,label:'项目',compoType:'searchSelect',listData:projectSerachSelect,required:true},
          {name:"launch_time",col:12,label:'发起时间',compoType:'datetime',required:true},
          {name:"task_level",col:12,label:'优先级',tooltip:"数值越小优先级越高",compoType:'select',listData:taskLevelName,required:true},
          {name:"task_type",col:24,label:'任务类型',compoType:'radio',defaultValue:1,listData:TaskTypeName,required:true},
          {name:"perform_status",col:24,label:'执行状态',compoType:'radio',defaultValue:1,listData:performStatusName,required:true},
          {name:"title",col:24,label:'任务标题',compoType:'textarea',required:true},
          {name:"content",col:24,label:'任务内容',compoType:'mdEditor',api:"work/taskRecord/uploadPics",required:true},
          {name:"start_time",col:12,label:'开始时间',compoType:'datetime',required:false},
          {name:"end_time",col:12,label:'结束时间',compoType:'datetime',required:false},
          {name:"remark",col:24,label:'备注',compoType:'textarea',rows:4,required:false},
        ]} 
        api="work/taskRecord/add"
        updateApi = "work/taskRecord/update"
        onBeforeSubmit={(params, cb) => {
          cb();
        }}
        onSucc={()=>{
          tableRef.reload();
        }}
      />}
    </> 
  );
};

export default Index;