import React, { useState} from 'react';
import {WsTable,WsForm,WsButton} from '@/components/WsTools';
import { Space } from 'antd';

var store = {};
const Index = () => {
  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const [data,setData] = useState({});
  const [formShow,setFormShow] = useState(false);
  let datas = [];
  for(let i=0;i<50;i++){
    datas.push(
      {
        id:i,
        supplier_name:'name'+i,
        supplier_contract_name:'contract'+i,
      }
    );
  }

  const formFunc = (data={})=>{
    setFormShow(true);
    setData(data);
  }

  return (
    <>
      <WsTable
        store={store}
        table = {tableRef}
        btns = {
          [
            {title:'添加',onClick:()=>{formFunc()}}
          ]
        }
        searchs={
          [
            {label:'人数',name:'number1',type:'input',placeholder:'人数'},
            {label:'运营人员',type:'select',name:'people',listData:{1:'male1_1',2:'male2_2'}},
            {label:'运营人员1',type:'selectInput',listData:{'id':'产品名称','name':'产品编码'}},
            {label:'日期',name:'date',type:'dateRange'},
          ]
        }
        th={[
          {title:'操作',name:'id',width:170,align:'center',align:'left',render:function(v,row){
            return (<>
              <Space>
              <WsButton title="添加" onClick={()=>{
                formFunc({});
                }}/>
              <WsButton title="编辑" onClick={()=>{formFunc(row);}}/>
              <WsButton title="删除" pop={true} onClick={()=>{
                loadApi(deleteMenu,{id:row.id},()=>{
                  tableRef.reload();
                },true);
              }}/>
              </Space>
              </>
            );
          }},
          {name:"supplier_name",tips:true,title:'供应商名称',width:320,render:v=>{return v||'-'}},
          {name:"supplier_contract_name",title:'供应商合同名称',width:200,render:v=>{return v||'-'}},
        ]}
        api={() => {
          return Promise.resolve({
            code: 0,
            data: datas,
            msg: '',
            total:100,
          });
        }}
      />
    {formShow&&<WsForm
        form={formRef}
        width={600}
        onCancel = {()=>{
          setFormShow(false);
        }}
        data = {data}
        fields={[
          {name:"supplier_co_op",label:'合作模式',type:'string',compoType:'select',listData:[],defaultOption:'请选择',required:true},
          {name:"supplier_contract_name",col:24,label:'供应商合同名称',compoType:'input',required:false},
          {name:"supplier_link",placeholder:'请以http(s)://开头',col:24,label:'供应商链接',compoType:'input'},
          {name:"remark",col:24,label:'备注',type:'string',compoType:'textarea',required:false},
        ]}
        onBeforeSubmit={(params, cb) => {
          console.log(params);
          cb();
        }}
      />}
    </> 
  );
};

export default Index;
