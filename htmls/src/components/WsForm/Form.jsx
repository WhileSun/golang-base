import React, { useState, useRef,useMemo, useImperativeHandle, useEffect} from 'react';
import requests from '@/request/index';
import './index.less';
import useForm from './hooks/useForm';
import ItemField from './components/itemField'
import {
  paramIsset,
  validateMessages,
  formValueClean,
  formFieldsTrans,
  getFieldToObj,
} from './utils/tools';
import {
  Form,
  Row,
  Col,
  Button,
  message,
} from 'antd';
import {WsModal} from '@/components/WsPopup';


const WsForm = (props, ref) => {
  //使用Form ref
  const [formRef] = Form.useForm();

  const initialize = () => {
    let objs = {};
    objs.type = paramIsset(props.type, 'Modal');
    objs.fieldToObj = getFieldToObj(paramIsset(props.fields,{}));
    objs.initData = paramIsset(props.data,{});
    objs.api = paramIsset(props.api, '');
    objs.updateApi = paramIsset(props.updateApi, objs.api);
    objs.idKey = paramIsset(props.idKey, 'id');
    objs.updateForm = objs.initData[objs.idKey] !== undefined;
    return objs;
  }

  const [state, setState] = useState(initialize);
  const [modelLoading, setModelLoading] = useState(false);
  const [modelShow, setModelShow] = useState(paramIsset(props.modelShow, true));//弹出框是否显示

  const closePopp = ()=> {
    setModelShow(false);
    if(props.cancel){props.cancel()};
  }

  const getData = async (apiParams={}) => {
    setModelLoading(true);
    let api = state.updateForm ? state.updateApi : state.api;
    requests.post(api, apiParams).then(function (resp) {
      console.log('resp',resp);
      setModelLoading(false);
      if(resp.code ==0){
        closePopp();
        if(props.onSucc){
          props.onSucc()
        }
        message.success(resp.msg);
      }else{
        message.error(resp.msg);
      }
    }).catch(function(error){
      setModelLoading(false);
      message.error("表单提交异常，请联系管理员处理！");
      console.log('error',error);
    });
  };

  //提交form
  const onFormFinish = (values) => {
    const params = formValueClean(state.fieldToObj, values);
    if(state.updateForm){
      params[state.idKey] = state.initData[state.idKey];
    }
    if(props.onBeforeSubmit){
      props.onBeforeSubmit(params, () => {
        getData(params);
      });
    }else{
      getData(params);
    }
  };

  //form数据打入
  useEffect(()=>{
    const initialValues = formFieldsTrans(state.fieldToObj, state.initData);
    formRef.setFieldsValue(initialValues);
  },[]);

  var formInstance = useForm(props.form);
  formInstance.setModelShow=()=>{}
  //支持原生ref
  useImperativeHandle(ref,()=>{return formInstance});

  const formBody = useMemo(() => {
    console.log('initForm');
    const fields = props.fields;
    return (
      <Form
        form={formRef}
        layout="horizontal"
        size="default"
        validateMessages={validateMessages}
        onFinish={onFormFinish}
        // labelCol={{ span: 6 }}
        // wrapperCol={{ span: 18 }}
        // onFinishFailed={onFormFinishFailed}
      >
         <Row gutter={24}>
           {fields.map((field, index) =>(<ItemField field={field} index={index} key={index}
          formRef={formRef} initData={state.initData}/>))}
          {state.type=='Form'&&
          <Col span={24}>
          <Form.Item style={{marginLeft:"110px"}}>
            <Button type="primary" htmlType="submit">
              提交
            </Button>
          </Form.Item>
          </Col>}
          </Row>
      </Form>
    );
  }, [props.fields]);

  if (state.type == 'Modal') {
    return (
      <WsModal
      content={formBody}
      show = {modelShow}
      width = {props.width}
      fullStatus = {props.modelFull}
      title = {state.updateForm ? "编辑"+props.title:"添加"+props.title}
      modelLoading = {modelLoading}
      cancel = {()=>{formRef.resetFields();closePopp();}}
      submit = {()=>{formRef.submit();}}
      />
    )
  } else if(state.type == 'Form'){
    return (
      <Row>
        <Col span={paramIsset(props.widthCol,12)}>
          {formBody}
        </Col>
      </Row>
    )
  }else{
    return '';
  }
};
export default WsForm;
