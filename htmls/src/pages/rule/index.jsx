import React, { useState} from 'react';
import WsForm from '@/components/WsForm';
import WsTable from '@/components/WsTable';
import {breakWords as bw,showHtml} from '@/utils/tools';


var store = {},store1 = {};
const fcode = function(s,n=5){
  return 'GY'+('0000000000'.slice(0,n-(s+'').length))+s;
};

const Index = () => {
  console.log('index');
  const supplierCoOp = {1:'推广',2:'采购'};
  const IsHasOrNot = {'f':'无','t':'有'};
  const IsDefault = {'f':'否','t':'是'};

  const formRef = WsForm.useForm();
  const tableRef = WsTable.useTable();
  const tableRef1 = WsTable.useTable();
  const [data,setData] = useState({});
  const [tableShow,setTableShow] = useState(false);
  const [formShow,setFormShow] = useState(false);
  return (
    <>
      <WsTable
        store={WsTable.useStore(store)}
        table = {tableRef}
        btns = {
          [
            {text:'添加',callback:()=>{setData({});setFormShow(true);}}
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
          {title:'操作',name:'id',width:60,align:'center',render:function(v,row){
            return showHtml(`<div>${Icon('修改','rowEditBtn')}${Icon('成本','rowUpdateCoOpBtn','修改合作模式')}</div>`);
          }},
          {name:"supplier_co_op",title:bw('合作|模式'),scrollFixed:true,align:'center',width:50,render:v=>{return supplierCoOp[v]||'-'}},
          {name:"supplier_id",tips:true,title:'供应商ID',scrollFixed:true,width:70,render:v=>{return fcode(v)}},
          {name:"supplier_name",tips:true,title:'供应商名称',width:160,render:v=>{return v||'-'}},
          {name:"supplier_contract_name",title:'供应商合同名称',width:160,render:v=>{return v||'-'}},
          {name:"supplier_address",tips:true,title:'供应商地址',width:160,render:v=>{return v||'-'}},
          {name:"supplier_refund_address",title:'供应商退货地址',width:160,render:v=>{return v||'-'}},
          {name:"supplier_contact_person",align:'center',title:bw('供应商|联系人'),width:60,render:v=>{return v||'-'}},
          {name:"supplier_phone",tips:true,title:'供应商电话',width:100,render:v=>{return v||'-'}},
          {name:"supplier_pay_account",tips:true,title:'供应商打款帐号',width:160,render:v=>{return v||'-'}},
          {name:"supplier_link",title:bw('供应商|链接'),width:60,render:v=>{return v?showHtml('<a target="_blank" title="'+v+'" href="'+v+'">打开链接</a> '):'-'}},
          {name:"has_contract",title:bw('有无|合同'),width:40,render:v=>{ return IsHasOrNot[v]||'-'}},
          {name:"support_one_item_shipping",title:bw('是否一|件代发'),width:45,render:v=>{return IsDefault[v]||'-'}},
          {name:"chosen_person",tips:true,title:'对接人',width:60},
          {name:"remark",tips:true,title:'备注',width:160,render:v=>{return v||'-'}},
          {name:"create_time",title:'创建时间',width:145,render:(v,row)=>{return showHtml(`<span title="${row.update_time?'更新于'+row.update_time:v}">${v}<span>`)}},
        ]}
        api="get_supplier_list"
        rowBtnsClick={
          (act,rowData)=>{
            if(act=='rowEditBtn'){
              setData(rowData);
              setFormShow(true);
            }else if(act=='rowUpdateCoOpBtn'){
              setData(rowData);
              setTableShow(true);
            }
          }
        }
      />
    {formShow&&<WsForm
        form={formRef}
        width={600}
        cancel = {()=>{
          setFormShow(false);
        }}
        data = {data}
        fields={[
          {name:"supplier_co_op",label:'合作模式',type:'string',disabled:(!data.id?false:true),compoType:'select',listData:supplierCoOp,defaultOption:'请选择',required:true},
          {name:"supplier_contract_name",col:24,label:'供应商合同名称',compoType:'input',required:false},
          {name:"supplier_pay_account",col:24,label:'供应商打款帐户',compoType:'input',required:false},
          {name:"supplier_name",col:24,label:'供应商名称',compoType:'input',required:true},
          {name:"supplier_address",col:24,label:'供应商地址',compoType:'input',required:true},
          {name:"supplier_refund_address",col:24,label:'供应商退货地址',compoType:'input',required:false},
          {name:"supplier_contact_person",label:'供应商联系人',compoType:'input',required:true},
          {name:"supplier_phone",label:'供应商电话',compoType:'input',required:true},
          {name:"supplier_link",placeholder:'请以http(s)://开头',col:24,label:'供应商链接',compoType:'input'},
          {name:"has_contract",label:'有无合同',compoType:'select',listData:IsHasOrNot,defaultOption:'请选择',required:true},
          {name:"support_one_item_shipping",label:'是否一件代发',compoType:'select',listData:IsDefault,defaultOption:'请选择',required:true},
          {name:"remark",col:24,label:'备注',type:'string',compoType:'textarea',required:false},
        ]}
        onBeforeSubmit={(params, cb) => {
          console.log(params);
          cb();
        }}
      />}
    {tableShow&&<WsTable
      store={WsTable.useStore(store1)}
      table = {tableRef1}
      type='modal'
      params = {{supplier_id:data.id}}
      cancel = {()=>{
        setTableShow(false);
      }}
      searchs={
        [
          {label:'人数',name:'number1',type:'input',placeholder:'人数'},
          {label:'日期',name:'date',type:'dateRange'},
        ]
      }
      th={[
        { title: '生效时间', name: 'start_time', width: 150 },
				{ title: '合作模式', name: 'supplier_co_op', width: 80, render: v => { return supplierCoOp[v] || '-' } },
				{ title: '备注', name: 'remark', width: 200, render: (v, row) => { return (v || '-') + '\n' + row.operator + '于' + row.create_time + '创建' } },
      ]}
      api="supplierCoOpLog_get_list"
      rowBtnsClick={
        (act,rowData)=>{
          if(act=='rowEditBtn'){
          }else if(act=='rowUpdateCoOpBtn'){
          }
        }
      }
    />}
    </> 
  );
};

export default Index;
