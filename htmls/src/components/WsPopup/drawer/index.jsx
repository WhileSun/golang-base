import { paramIsset } from '../utils/tools';
import React, { useState } from 'react';
import { Drawer } from 'antd';
const WsDrawer = (props) => {
  const handleCancel = (e) => {
    if (props.cancel) {props.cancel()}
  };
  const handleSubmit = () => {
    if (props.submit) {props.submit();}
  };
  return (
    <>
      <Drawer
        title={paramIsset(props.title, '提示')}
        width={paramIsset(props.width, 600)}
        onClose={handleCancel}
        visible={paramIsset(props.show, false)}
        destroyOnClose={true}
        keyboard={false}
        maskClosable={false}
        placement={paramIsset(props.position,'right')}
        bodyStyle={{ paddingBottom: 80 }}
        extra={
          <Space>
            <Button>取消</Button>
            <Button type="primary">提交</Button>
          </Space>
        }
      >
        {props.content}
      </Drawer>
    </>
  );
};
