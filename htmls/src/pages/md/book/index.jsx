import React, { useState, useRef,useMemo, useEffect} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import WsButton from '@/components/WsButton'
import {loadApi} from '@/utils/tools';
import { Space} from 'antd';
import { history, Link } from 'umi';
import {getMdBookList,addMdBook,updateMdBook} from '@/services/api';
import qs from 'qs';

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
            {type:'selectInput',listData:{'book_name':'书籍名称'}},
          ]
        }
        btns = {
          [
            {text:'添加',callback:()=>{formFunc({});}}
          ]
        }
        th={[
          {name:"book_name",title:'书籍名称',width:150,render:v=>{return v||'-'}},
          {name:"book_ident",title:'书籍标识',width:200,align:'left',render:v=>{return v||'-'}},
          {name:"created_at",title:'创建时间',width:140,align:'center',render:v=>{return v||'-'}},
          {name:"id",title:'文档内容',width:80,align:'center',render:(v,row)=>{
            return ( 
                <Link target = "_blank" to={{pathname:'/daily/md/document',search:qs.stringify({book_id:v})}}>
                  <WsButton title="前往" onClick={()=>{
                     }}/>
                </Link>
              )
          }},
          {title:'操作',name:'id',width:80,align:'center',render:function(v,row){
            return (<Space>
              <WsButton title="编辑" onClick={()=>{
                formFunc(row);
              }}/>
              {/* <WsButton title="删除" pop={true} onClick={()=>{
                loadApi(deletePerms,{id:row.id},()=>{
                  tableRef.reload();
                },true);
              }}/> */}
            </Space>);
          }},
        ]}
        api={getMdBookList}
      />
    {formShow&&<WsForm
        form={formRef}
        title="书籍"
        width={450}
        onCancel = {()=>{
          setFormShow(false);
        }}
        data = {formData}
        fields={[
          {name:"book_name",col:24,label:'书籍名称',compoType:'input',required:true},
          {name:"book_ident",col:24,label:'书籍标识',tooltip:"用于区分书籍",placeholder:'不填写,系统自动生成',compoType:'input',required:false},
        ]}    
        api={addMdBook}
        updateApi = {updateMdBook}
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