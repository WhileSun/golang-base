import React, { useState, useRef,useMemo, useEffect} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import WsButton from '@/components/WsButton'
import {breakWords as bw,inArray,getData,arrTransName} from '@/utils/tools';
import { Space} from 'antd';
import {statusFunc} from '@/module/colorfunc';
import {defaultPasswd,userSuperName} from '@/config';

var store = {};
export default (props) => {
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [formData,setFormData] = useState({});
  const [formShow,setFormShow] = useState(false);
  const [roleSerachSelect,setRoleSerachSelect] = useState([]);
  const StatusList = {'true':'正常','false':'关闭'};

  const formFunc = (row)=>{
    setFormData(row)
    getData('load/role/list/get',{},(data)=>{
      setRoleSerachSelect(arrTransName(data,{role_name:'label',id:'value'}));
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
            {type:'selectInput',listData:{'u*realname':'用户名称','u*username':'登录账号'}},
            {label:'状态',type:'select',name:'u*status',listData:StatusList},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{formFunc({});}}
          ]
        }
        th={[
          {name:"username",title:'登录账号',width:120,render:v=>{return v||'-'}},
          {name:"realname",title:'用户名称',width:150,align:'left',render:v=>{return v||'-'}},
          {name:"status",title:'用户状态',width:80,align:'center',render:v=>{return statusFunc(v,StatusList)||'-'}},
          {name:"role_names",title:'用户关联角色',width:200,align:'left',render:v=>{return v||'-'}},
          {name:"created_at",title:'创建时间',width:140,align:'center',render:v=>{return v||'-'}},
          {title:'操作',name:'id',width:200,align:'left',render:function(v,row){
            if(row.username == userSuperName){
              return "";
            }
            return (<Space>
              <WsButton title="编辑" onClick={()=>{
                row.role_ids = row.role_ids_str.split(",");
                formFunc(row);
              }}/>
            </Space>);
          }},
        ]}
        api="user/list/get"
        rowBtnsClick={
          (act,rowData)=>{
           
          }
        }
      />
    {formShow&&<WsForm
        form={formRef}
        width={500}
        title="用户"
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"username",col:24,label:'用户账号',tooltip:"用户在系统中登录的账号",extra:defaultPasswd,compoType:'input',required:true},
          {name:"realname",col:24,label:'用户名称',tooltip:"用户在系统中展示的名称",compoType:'input',required:true},
          {name:"status",col:24,label:'用户状态',compoType:'radio',defaultValue:'true',listData:StatusList,required:true},
          {name:"role_ids",mode:'multiple',col:24,label:'角色',compoType:'searchSelect',listData:roleSerachSelect,required:true},
        ]}    
        api="user/add"
        updateApi = "user/update"
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