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

  const formFunc = (row)=>{
    setFormData(row)
    setFormShow(true);
  }

  return (
    <>
      <WsTable
        store={store}
        table = {tableRef}
        searchs={
          [
            {type:'selectInput',listData:{'name':'节点名称','page_perms':'操作权限标识'}},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{formFunc({});}}
          ]
        }
        th={[
          {name:"name",title:'节点名称',width:150,render:v=>{return v||'-'}},
          {name:"page_perms",title:'操作权限标识',width:200,align:'left',render:v=>{return v||'-'}},
          {name:"data_perms",title:'数据权限标识',width:200,align:'left',render:v=>{return v||'-'}},
          {name:"created_at",title:'创建时间',width:140,align:'center',render:v=>{return v||'-'}},
          {title:'操作',name:'id',width:150,align:'center',render:function(v,row){
            return (<Space>
              <WsButton title="编辑" onClick={()=>{
                formFunc(row);
              }}/>
              <WsButton title="删除" pop={true} onClick={()=>{
                getData('perms/delete',{id:row.id},()=>{
                  tableRef.reload();
                },true);
              }}/>
            </Space>);
          }},
        ]}
        api="perms/list/get"
        rowBtnsClick={
          (act,rowData)=>{
           
          }
        }
      />
    {formShow&&<WsForm
        form={formRef}
        width={500}
        title="节点"
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"name",col:24,label:'节点名称',tooltip:"描述功能的名称 例：列表、更新等",compoType:'input',required:true},
          {name:"page_perms",col:24,label:'操作权限标识',tooltip:"用于前端按钮的权限控制，用大写英文描述 例：LIST、UPDATE等",compoType:'input',required:true},
          {name:"data_perms",col:24,label:'数据权限标识',tooltip:"用于后端接口的尾缀，区分大小写 例 /list/get、/update",compoType:'input',required:true},
        ]}    
        api="perms/add"
        updateApi = "perms/update"
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