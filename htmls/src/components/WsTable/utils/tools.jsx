import moment from 'moment';
import $ from 'jquery';
//参数是否存在
const paramIsset = (param, defaultValue = '') => {
  return param === undefined ? defaultValue : param;
};

const funcIsset = (func, defaultValue = ()=>{}) => {
  return func === undefined ? defaultValue : func();
};

//state更新不是替换
const setUpdateState = (oldState, newState, func) => {
  const stateObj = Object.assign({}, oldState, newState);
  func(stateObj);
};

//对form提交的数据进行清洗
const momentDate = (momentObj, type = 'date') => {
  let datestr = '';
  switch (type) {
    case 'date':
      datestr = moment(momentObj).format('YYYY-MM-DD');
      break;
    case 'datetime':
      datestr = moment(momentObj).format('YYYY-MM-DD hh:mm:ss');
      break;
  }
  return datestr;
};

function S4() {
  return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1);
}

function getRandStr(prefix = '') {
  var uuid = prefix + S4() + S4() + S4();
  return uuid;
}

function toTree(rows,parent_id=0,parent_ids=[]) {
  if(rows ==undefined){
    return [];
  }
  let node =[];
  let filterRows = rows.filter(row => row.parent_id === parent_id)
  if(filterRows.length){
    filterRows.forEach(filterRow=>{
      if(parent_id==0){
        filterRow.parent_ids = parent_ids
      }else{
        filterRow.parent_ids = [...parent_ids,parent_id]
      }
      let n = toTree(rows, filterRow.id,filterRow.parent_ids)
      if(n.length){
        filterRow.children=n;
      }
      node.push(filterRow)
    })
  }
  return node
}

function filterName(rows,index,val,new_rows=[]){
  rows.forEach(row=>{
    if(row[index].toString().includes(val)){
      new_rows.push(row);
    }else{
      //存在子集的情况
      if(row?.children && row?.children?.length>0){
        filterName(row.children,index,val,new_rows)
      }
    }
  })
  return new_rows;
}

const parseFormParams=(params)=>{
  let newParams = {};
  for(let field in params){
    let val = params[field];
    if(field.indexOf('selectInput_')>-1){
      if(!!$.trim(val['val'])){
        newParams[val['key']] = val['val'];
      }   
    }else if(field.indexOf('dateRange_')>-1){
      if(val != ""){
        newParams[field.replace('dateRange_',"")] = momentDate(val[0])+'~'+momentDate(val[1]);
      }
    }else{
      if(!!$.trim(val)){
        newParams[field] = val;
      }
    }
  }
  return newParams;
}

const setFormParamStore=(store,params)=>{
  if(store){
    if(Object.keys(params).length==0){
      for(let key in store){
        delete store[key];
      }
    }else{
      for(let key in params){
        store[key] = params[key]
      }
    }
  }
}


const arrayColumn = (arr,v,k="")=>{
  let lists = [];
  let objs = {};
  arr.forEach(val=>{
    if(k ==""){
      lists.push(val[v]);
    }else{
      objs[val[k]] = val[v];
    }
  })
  if(k == ""){
    return lists;
  }else{
    return objs;
  }
}

export { paramIsset, setUpdateState, momentDate, funcIsset, getRandStr,toTree,parseFormParams,setFormParamStore,arrayColumn,filterName};
