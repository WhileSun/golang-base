import React, { useState} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import WsButton from '@/components/WsButton'
import { Space} from 'antd';

var store = {};
const Index = (props) => {
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [formData,setFormData] = useState({});
  const [formShow,setFormShow] = useState(false);

  return (
    <>
      <WsTable
        store={store}
        table = {tableRef}
        searchs={
          [
            {type:'selectInput',listData:{'project_name':'项目名称'}},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{setFormData({});setFormShow(true);}}
          ]
        }
        th={[
          {name:"project_name",title:'项目名称',width:200,render:v=>{return v||'-'}},
          {name:"remark",title:'备注',width:250,render:v=>{return v||'-'}},
          {name:"created_at",title:'创建时间',width:120,align:'center',render:v=>{return v||'-'}},
          {title:'操作',name:'id',width:150,align:'center',align:'left',render:function(v,row){
            return (<Space>
              <WsButton title="编辑" onClick={()=>{setFormData(row);setFormShow(true);}}/>
            </Space>);
          }},
        ]}
        api="work/taskProject/list/get"
      />
    {formShow&&<WsForm
        form={formRef}
        width={500}
        title="项目"
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"project_name",col:24,label:'项目名称',compoType:'input',required:true},
          {name:"remark",col:24,label:'备注',compoType:'textarea',required:true},
        ]}    
        api="work/taskProject/add"
        updateApi = "work/taskProject/update"
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