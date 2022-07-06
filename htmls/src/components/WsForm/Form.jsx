import './index.less';
import React, { useState, useRef, useMemo, useImperativeHandle, useEffect } from 'react';
import useForm from './hooks/useForm';
import ItemField from './components/itemField';
import { paramIsset, validateMessages, formValueClean, formFieldsTrans, getFieldToObj } from './utils/tools';
import { Form, Row, Col, Button, message } from 'antd';
import { WsModal, WsDrawer } from '@/components/WsPopup';


const WsForm = (props, ref) => {
  //使用Form ref
  const [formRef] = Form.useForm();

  //基础配置
  const config = useMemo(() => {
    let param = {};
    param.mode = paramIsset(props.mode, 'Modal'); //form表单的展示模式
    param.fieldToObj = getFieldToObj(paramIsset(props.fields, {})); //name->th
    param.initData = paramIsset(props.data, {}); //初始化值
    param.api = paramIsset(props.api, ''); //添加api
    param.updateApi = paramIsset(props.updateApi, param.api); //修改api
    param.idKey = paramIsset(props.idKey, 'id'); //主键名称
    param.updateForm = param.initData[param.idKey] != undefined; //是否是更新
    param.fields = paramIsset(props.fields, []); // form字段

    param.onCancel = props.onCancel; //关闭执行
    param.onSucc = props.onSucc; //成功后回调
    param.onBeforeSubmit = props.onBeforeSubmit; //提交前回调

    param.title = paramIsset(props.title, ''); //弹出框标题
    param.width = paramIsset(props.width, 600); //弹出框宽度
    param.fullStatus = paramIsset(props.fullStatus, false); //弹出框是否全屏

    param.widthCol = paramIsset(props.widthCol, 12); //表单宽度
    return param;
  }, [props])

  const [loading, setLoading] = useState(false); //加载状态
  const [modelShow, setModelShow] = useState(paramIsset(props.modelShow, true));//弹出框是否显示

  //弹出框关闭事件
  const closePopFunc = () => {
    setModelShow(false);
    if (config.onCancel) { config.onCancel() };
  }

  const getData = async (apiParams = {}) => {
    setLoading(true);
    let api = config.updateForm ? config.updateApi : config.api;
    api.call(this, apiParams).then(function (resp) {
      console.log('resp', resp);
      setLoading(false);
      if (resp.code == 0) {
        closePopFunc();
        if (config.onSucc) {
          config.onSucc()
        }
        message.success(resp.msg);
      } else {
        message.error(resp.msg);
      }
    }).catch(function (error) {
      setLoading(false);
      message.error("表单提交异常，请联系管理员处理！");
      console.log('error', error);
    });
  };

  //提交form
  const onFormFinish = (values) => {
    const params = formValueClean(config.fieldToObj, values);
    if (config.updateForm) {
      params[config.idKey] = config.initData[config.idKey];
    }
    if (config.onBeforeSubmit) {
      config.onBeforeSubmit.call(this, params, () => {
        getData(params);
      });
    } else {
      getData(params);
    }
  };

  //form数据打入
  useEffect(() => {
    const initialValues = formFieldsTrans(config.fieldToObj, config.initData);
    formRef.setFieldsValue(initialValues);
  }, []);

  var formInstance = useForm(props.form);
  formInstance.setModelShow = () => { }
  //支持原生ref
  useImperativeHandle(ref, () => { return formInstance });

  const formHtml = useMemo(() => {
    console.log('initForm');
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
          {config.fields.map((field, index) => (<ItemField field={field} index={index} key={index}
            formRef={formRef} initData={config.initData} />))}
          {config.mode == 'Form' &&
            <Col span={24}>
              <Form.Item style={{ marginLeft: "110px" }}>
                <Button type="primary" htmlType="submit">
                  提交
                </Button>
              </Form.Item>
            </Col>}
        </Row>
      </Form>
    );
  }, [config.fields]);

  if (config.mode == 'Modal') {
    return (
      <WsModal
        content={formHtml}
        show={modelShow}
        width={config.width}
        fullStatus={config.fullStatus}
        title={config.updateForm ? "编辑" + config.title : "添加" + config.title}
        loading={loading}
        onCancel={() => { formRef.resetFields(); closePopFunc(); }}
        onSubmit={() => { formRef.submit(); }}
      />
    )
  } else if (config.mode == 'Form') {
    return (
      <Row>
        <Col span={config.widthCol}>
          {formHtml}
        </Col>
      </Row>
    )
  } else if (config.mode == 'Drawer') {
    return (
      <WsDrawer
        content={formHtml}
        show={modelShow}
        width={config.width}
        fullStatus={config.fullStatus}
        title={config.updateForm ? "编辑" + config.title : "添加" + config.title}
        loading={loading}
        onCancel={() => { formRef.resetFields(); closePopFunc(); }}
        onSubmit={() => { formRef.submit(); }}
      />);
  } else {
    return '';
  }
};
export default WsForm;
