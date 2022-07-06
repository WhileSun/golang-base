import { paramIsset } from '../utils/tools';
import React, { useState,useMemo} from 'react';
import { Drawer,Space,Button} from 'antd';
import './index.less'

const WsDrawer = (props) => {

  const config = useMemo(()=>{
    let param = {};
    param.content = paramIsset(props.content,"");
    param.title = paramIsset(props.title,"提示");
    param.width = paramIsset(props.width,600);
    param.loading = paramIsset(props.loading,false);
    param.show = paramIsset(props.show,false);
    param.position = paramIsset(props.position, 'right');

    param.onCancel = props.onCancel;
    param.onSubmit = props.onSubmit;
    return param;
  },[props]);

  const handleCancel = (e) => {
    if (config.onCancel) {config.onCancel()}
  };
  const handleSubmit = () => {
    if (config.onSubmit) {config.onSubmit();}
  };

  const buttonHtml = ()=>{
    let btns = [<Button key="back" onClick={handleCancel}>取消</Button>];
    if(config.onSubmit){
      btns.push(<Button key="submit" type="primary" onClick={handleSubmit} loading={config.loading}>提交</Button>);
    }
    return btns;
  }

  return (
    <>
      <Drawer
        className = "drawer-wrap-style"
        title={config.title}
        width={config.width}
        onClose={handleCancel}
        visible={config.show}
        destroyOnClose={true}
        keyboard={false}
        maskClosable={false}
        placement={config.position}
        bodyStyle={{ paddingBottom: 80 }}
        extra={
          <Space>
          {buttonHtml()}
          </Space>
        }
      >
        {props.content}
      </Drawer>
    </>
  );
};

export default WsDrawer;