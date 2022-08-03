import React, { useState, useRef,useMemo, useEffect} from 'react';
import {WsButton,WsForm,WsTable} from '@/components/WsTools';
import {breakWords as bw,inArray,getData,loadApi,toTree,arrTransName} from '@/utils/tools';
import {statusFunc,menuTypeFunc} from '@/module/colorfunc';
import { Space} from 'antd';
import {WsIcon} from '@/components/WsTools';
import {getMenuList,addMenu,updateMenu,deleteMenu,getLoadMenuList,getLoadPermsList} from '@/services/api';

var store = {};

const getMenuPagePerms= (data,id)=>{
  let page_perms = [];
  data.forEach(val => {
    if(val.parent_id == id){
      if(val?.menu_type ==2){
        page_perms.push(val.page_perms)
      }
    }
  });
  return page_perms;
}

const Index = (props) => {
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [formData,setFormData] = useState({});
  const [formShow,setFormShow] = useState(false);
  const [menuTreeSelect,setMenuTreeSelect] = useState([]);
  const [permsList,setPermsList] = useState([]);
  const [menuType,setMenuType] = useState(1);
  const MenuStatusList = {'true':'正常','false':'关闭'};
  const MenuShowList = {'true':'是','false':'否'};
  const MenuTypeList = {1:'菜单',2:'节点'};
  
  const getInitMenuList = (cb)=>{
    loadApi(getLoadMenuList,{menu_type:1},(data)=>{
      setMenuTreeSelect(toTree(data));
      cb;
    });
  }
  const getInitPermsList = (cb)=>{
    loadApi(getLoadPermsList,{},(data)=>{
      setPermsList(arrTransName(data,{name:'title',page_perms:'key'}));
      cb;
    });
  }

  const formFunc = (row)=>{
    if(row.id && row.menu_type == 1){
      row.target_keys = getMenuPagePerms(tableRef.getDataList(),row.id);
    }else{
      row.target_keys = [];
    }
    setFormData(row);
    getInitMenuList(
      getInitPermsList(setFormShow(true))
    );
  }
  return (
    <>
      <WsTable
        store={store}
        table = {tableRef}
        treeTable = {true}
        searchs={
          [
            {type:'selectInput',listData:{'menu_name':'菜单名称'}},
          ]
        }
        btns = {
          [
            {title:'添加',onClick:()=>{setMenuType(1);formFunc({});}}
          ]
        }
        th={[
          {name:"menu_name",title:'菜单名称',width:200,render:v=>{return v||'-'}},
          {name:"sort",title:'菜单排序',width:60,align:'center',render:v=>{return v||'-'}},
          {name:"icon",title:'图标',width:40,align:'center',render:v=>{return (<WsIcon type={v} style={{fontSize: '15px'}}/>)||'-'}},
          {name:"menu_type",title:'菜单类型',width:80,align:'center',render:v=>{return menuTypeFunc(v,MenuTypeList)||'-'}},
          {name:"url",title:'栏目路径',width:160,render:v=>{return v||'-'}},
          {name:"status",title:'菜单状态',width:60,align:'center',render:v=>{return statusFunc(v,MenuStatusList)||'-'}},
          {name:"page_perms",title:bw('操作|权限标识'),width:120,render:v=>{return v||'-'}},
          {name:"data_perms",title:bw('数据|权限标识'),width:160,render:v=>{return v||'-'}},
          {title:'操作',name:'id',width:170,align:'center',align:'left',render:function(v,row){
            if(row.is_sys){
              return '';
            }
            return (<Space>
              {row.menu_type == 1?<WsButton title="添加" onClick={()=>{
                setMenuType(1);
                let new_row = {};
                new_row.parent_ids = [...row.parent_ids,row.id];
                formFunc(new_row);
                }}/>:""}
              <WsButton title="编辑" onClick={()=>{setMenuType(row.menu_type);formFunc(row);}}/>
              <WsButton title="删除" pop={true} onClick={()=>{
                loadApi(deleteMenu,{id:row.id},()=>{
                  tableRef.reload();
                },true);
              }}/>
            </Space>);
          }},
        ]}
        api={getMenuList}
        onLocalFilter={(params) => {
          tableRef.filterName("menu_name",params?.menu_name);
        }}
    />
    {formShow&&<WsForm
        form={formRef}
        title="菜单栏目"
        onCancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"parent_ids",col:24,label:'上级菜单',compoType:'cascader',extra:"不选，默认为顶级菜单栏目",listData:menuTreeSelect,parentSelect:true,fieldNames:{label:'menu_name',value:'id'},required:false},
          {name:"menu_type",col:24,label:'菜单类型',compoType:'radio',defaultValue:1,listData:MenuTypeList,disabled:!!formData.id,onChange:(event)=>{  
            setMenuType(event.target.value);
          },required:true},
          {name:"menu_name",col:24,label:'菜单名称',compoType:'input',required:true},
          {name:"url",col:24,label:'请求路径',compoType:'input',required:true,remove:inArray(menuType,[2],'int')},
          {name:"page_perms",col:24,label:'操作权限标识',compoType:'input',tooltip:"前端操作按钮权限标识",required:false,hidden:inArray(menuType,[1],'int')},
          {name:"data_perms",col:24,label:'数据权限标识',compoType:'input',tooltip:"后端数据接口权限标识",required:false,hidden:inArray(menuType,[1],'int')},
          {name:"sort",col:24,label:'显示排序',compoType:'inputNumber',required:true},
          {name:"icon",col:24,label:'图标',compoType:'input',required:false,hidden:inArray(menuType,[2],'int')},
          {name:"status",col:24,label:'菜单状态',compoType:'radio',defaultValue:'true',listData:MenuStatusList,required:true},
          {name:"show",col:24,label:'是否显示',compoType:'radio',defaultValue:'true',listData:MenuShowList,required:true,hidden:inArray(menuType,[2],'int')},
          {name:"data_perms_header",col:24,label:'数组权限标识',tooltip:'节点数据权限标识前缀，例：user、menu等(只对新增的节点有效)',compoType:'input',required:false,hidden:inArray(menuType,[2],'int')},
          {name:"target_keys",col:24,label:'节点选择',compoType:'transfer',required:false,listData:permsList,topTitle:['节点列表','已赋值节点'],hidden:inArray(menuType,[2],'int')},
        ]}    
        api={addMenu}
        updateApi = {updateMenu}
        onBeforeSubmit={(params, cb) => { 
          if(params.parent_ids == "" || params.parent_ids == undefined){
            params.parent_id = 0;
          }else{
            let len = (params.parent_ids).length-1;
            params.parent_id = params.parent_ids[len];
          }
          delete params.parent_ids;
          params.old_target_keys = formData.target_keys;
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
