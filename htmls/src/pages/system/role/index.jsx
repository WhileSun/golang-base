import React, { useState, useRef,useMemo, useEffect} from 'react';
import {WsButton,WsForm,WsTable} from '@/components/WsTools';
import {breakWords as bw,inArray,getData,toTree,loadApi} from '@/utils/tools';
import { Space,message,Tag } from 'antd';
import {statusFunc} from '@/module/colorfunc';
import { FileOutlined,FolderOutlined} from '@ant-design/icons';
import {roleSuperName} from "@/config"
import {getRoleList,addRole,updateRole,getMenuNameList} from "@/services/api";

var store = {};
export default (props) => {
  console.log("props",props);
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [formData,setFormData] = useState({});
  const [formShow,setFormShow] = useState(false);
  const [permsTree,setPermsTree] = useState([]);
  const StatusList = {'true':'正常','false':'关闭'};

  const formFunc = (row)=>{
    setFormData(row)
    loadApi(getMenuNameList,{},(data)=>{
      const permsTrans = (rows)=>{
        return rows.map((row)=>{
          if(row['status']){
            let menuRow  ={};
            menuRow.title = row['menu_name'];
            menuRow.key = String(row['id']);
            menuRow.icon = row['menu_type'] === 2 ?<FileOutlined />:<FolderOutlined />;
            if(row['children'] !==undefined){
              menuRow.children =  permsTrans(row['children']);
            }
            return menuRow;
          }
        }).filter(result => result!==undefined)
      }
      setPermsTree(permsTrans(toTree(data)));
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
            {type:'selectInput',listData:{'role_name':'角色名称'}},
            {label:'状态',type:'select',name:'status',listData:StatusList},
          ]
        }
        btns = {
          [
            {title:'添加',onClick:()=>{formFunc({});}}
          ]
        }
        th={[
          {name:"role_name",title:'角色名称',width:120,align:'center',render:v=>{return v||'-'}},
          {name:"sort",title:'排序',width:100,align:'center',render:v=>{return v}},
          {name:"status",title:'状态',width:100,align:'center',render:v=>{return statusFunc(v,StatusList)||'-'}},
          {title:'操作',name:'id',width:200,align:'center',align:'left',render:function(v,row){
            if(row.role_identity == roleSuperName){
              return "";
            }
            return (<Space>
              <WsButton title="编辑" onClick={()=>{
                row.perms = row.perms_ids.split(',');
                formFunc(row);
              }}/>
            </Space>);
          }},
        ]}
        api={getRoleList}
        rowBtnsClick={
          (act,rowData)=>{
           
          }
        }
      />
    {formShow&&<WsForm
        form={formRef}
        title="角色"
        onCancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"role_name",col:24,label:'角色名称',compoType:'input',required:true},
          {name:"sort",col:24,label:'显示排序',compoType:'inputNumber',required:true},
          {name:"status",col:24,label:'角色状态',compoType:'radio',defaultValue:'true',listData:StatusList,required:true},
          {name:"perms",col:24,label:'菜单权限',compoType:'menuTree',listData:permsTree,required:true},
        ]}    
        api={addRole}
        updateApi = {updateRole}
        onBeforeSubmit={(params, cb) => {
          params.perms_ids = params.perms.join(',');
          delete params.perms;
          cb();
        }}
        onSucc={()=>{
          tableRef.reload();
        }}
      />}
    </> 
  );
};