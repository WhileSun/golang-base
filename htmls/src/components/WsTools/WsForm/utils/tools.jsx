import moment from 'moment';

//参数是否存在
const paramIsset = (param, defaultValue = '') => {
  return param === undefined ? defaultValue : param;
};

const getFieldToType = (fields) => {
  let row = {};
  for (let field of fields) {
    row[field['name']] = field['compoType'];
  }
  return row;
};

const getFieldToObj = (fields) => {
  let row = {};
  for (let field of fields) {
    row[field['name']] = field;
  }
  return row;
};


//state更新不是替换
const setUpdateState = (oldState, newState, func) => {
  const stateObj = Object.assign({}, oldState, newState);
  func(stateObj);
};

const validateMessages = {
  required: '[${label}]是必填字段',
  whitespace: '[${label}]不能是空字段',
};

const createFormRules = (field) => {
  let rules = [];
  if (field['compoType'] == 'input') {
    field['required'] === true ? rules.push({ required: true }, { whitespace: true }) : '';
  } else {
    field['required'] === true ? rules.push({ required: true }) : '';
  }
  field['rules'] === undefined ? '' : rules.push.apply(field['rules']);
  return rules;
};

//对form提交的数据进行清洗
const formValueClean = (fieldToObj, values) => {
  console.log(values);
  for (let field in fieldToObj) {
    let obj = fieldToObj[field];
    if(values[field] === undefined || values[field] ==="" ||  values[field] === null){
      values[field] = ""
    }else{
      switch (obj['compoType']) {
        case 'date':
          values[field] = moment(values[field]).format('YYYY-MM-DD');
          break;
        case 'datetime':
          values[field] = moment(values[field]).format('YYYY-MM-DD HH:mm:ss');
          break;
      }
    }
  }
  return values;
};

//表单初始化转化
const formFieldsTrans = (fieldToObj, row) => {
  let values = Object.assign({}, row);
  for (let field in values) {
    if(fieldToObj[field] !== undefined){
      values[field] = _fieldTrans(fieldToObj[field],values[field]);
    }
  }
  return values;
};

//表单单行转化
const formFieldTrans = (field) => {
  if(field.defaultValue === undefined){
    return;
  }else{
    return _fieldTrans(field,field.defaultValue);
  }
};

const _fieldTrans =(field, val)=>{
  let newVal = "";
  switch (field.compoType) {
    case 'date':
    case 'datetime':
      if(val === null){
        newVal = "";
      }else{
        newVal = moment(val);
      }
      break;
    case 'select':
    case 'input':
    case 'radio':
      //默认value需要转为字符，不然显示会有问题
      newVal = val.toString();
      break;
    case 'searchSelect':
      if(Array.isArray(val)){
        newVal = val.filter(v=>v!="")
      }else{
        newVal = val.toString();
      }
      break;
    default:
      newVal = val;
      break;
  }
  return newVal;
}

const delArrVal = (arr, val)=>{
  let index = arr.indexOf(val);
  if (index > -1) {
    arr.splice(index, 1);
  }
  return arr;
}

const loadPromise = (api,params,callback,onError)=>{
  return api.call(this,params).then((resp)=>{
    if(resp.code ==0){
        if(callback){
            return callback(resp.data);
        }
    }else{
      if(onError){
        onError(resp.msg)
      }
    }
}).catch(function(error){
  if(onError){
    onError(error)
  }
});
}

export {
  paramIsset,
  createFormRules,
  setUpdateState,
  formValueClean,
  getFieldToType,
  formFieldsTrans,
  getFieldToObj,
  formFieldTrans,
  validateMessages,
  loadPromise,
  delArrVal,
};
