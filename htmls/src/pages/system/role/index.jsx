import React, { useState, useRef,useMemo, useEffect} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import WsButton from '@/components/WsButton'
import {breakWords as bw,inArray,getData,toTree} from '@/utils/tools';
import { Space,message,Tag } from 'antd';
import {statusFunc} from '@/module/colorfunc';
import { FileOutlined,FolderOutlined} from '@ant-design/icons';
import {roleSuperName} from "@/config"

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
    getData('load/menu/list/get',{},(data)=>{
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
            {type:'selectInput',listData:{'role_name':'角色名称','role_identity':'权限标识'}},
            {label:'状态',type:'select',name:'status',listData:StatusList},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{formFunc({});}}
          ]
        }
        th={[
          {name:"role_name",title:'角色名称',width:120,align:'center',render:v=>{return v||'-'}},
          {name:"role_identity",title:'角色权限标识',width:100,align:'center',render:v=>{return v||'-'}},
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
        api="role/list/get"
        rowBtnsClick={
          (act,rowData)=>{
           
          }
        }
      />
    {formShow&&<WsForm
        form={formRef}
        width={500}
        title="角色"
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"role_name",col:24,label:'角色名称',compoType:'input',required:true},
          {name:"role_identity",col:24,label:'角色权限标识',compoType:'input',disabled:!!formData.id,tooltip:"角色唯一标识，创建后不能再修改",required:true},
          {name:"sort",col:24,label:'显示排序',compoType:'inputNumber',required:true},
          {name:"status",col:24,label:'角色状态',compoType:'radio',defaultValue:'true',listData:StatusList,required:true},
          {name:"perms",col:24,label:'菜单权限',compoType:'menuTree',listData:permsTree,required:true},
        ]}    
        api="role/add"
        updateApi = "role/update"
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