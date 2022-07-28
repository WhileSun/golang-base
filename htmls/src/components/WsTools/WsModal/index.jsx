import { paramIsset } from '../utils/tools';
import React, { useState, useMemo } from 'react';
import Draggable from 'react-draggable';
import { Modal, Button } from 'antd';
import $ from 'jquery';
import { ArrowsAltOutlined, ShrinkOutlined } from '@ant-design/icons';
import './index.less'

const WsModal = (props) => {

  const config = useMemo(() => {
    let param = {};
    param.content = paramIsset(props.content, "");
    param.title = paramIsset(props.title, "提示");
    param.width = paramIsset(props.width, 600);
    param.fullVisible = paramIsset(props.fullVisible, false);
    param.fullStatus = paramIsset(props.fullStatus, false);
    param.loading = paramIsset(props.loading, false);
    param.visible = paramIsset(props.visible, false);
    param.forceRender = paramIsset(props.forceRender, false);
    param.okText = paramIsset(props.okText, '确定');
    param.footer = paramIsset(props.footer, []);

    param.onCancel = props.onCancel;
    param.onOk = props.onOk;
    return param;
  }, [props]);

  //model参数
  const stateObj = {
    modelMoveDisabled: true,
    modelMoveBounds: { left: 0, top: 0, bottom: 0, right: 0 },
  };
  const [modelState, setModelState] = useState(stateObj);
  const [fullStatus, setFullStatus] = useState(config.fullStatus);

  const updateModelState = (newState) => {
    setModelState({ ...modelState, ...newState })
  };
  //model拖拉
  const draggleRef = React.createRef();
  const handleModelMove = (event, uiData) => {
    const { clientWidth, clientHeight } = window?.document?.documentElement;
    const targetRect = draggleRef?.current?.getBoundingClientRect();
    updateModelState({
      modelMoveBounds: {
        left: -targetRect?.left + uiData?.x,
        right: clientWidth - (targetRect?.right - uiData?.x),
        top: -targetRect?.top + uiData?.y,
        bottom: clientHeight - (targetRect?.bottom - uiData?.y),
      },
    });
  };
  //关闭
  const handleCancel = (e) => {
    if (config.onCancel) { config.onCancel(e) }
  };
  //提交
  const handleSubmit = (e) => {
    if (config.onOk) { config.onOk(e) }
  };

  const footer = () => {
    if (config.footer.length > 0) {
      return config.footer;
    } else {
      let btns = [<Button key="back" onClick={handleCancel}>取消</Button>,
      <Button key="submit" type="primary" onClick={handleSubmit} loading={config.loading}>{config.okText}</Button>]; 
      return btns;
    }
  }

  const toggleFullScreen = () => {
    setFullStatus(!fullStatus);
  }

  return (
    <>
      <Modal
        getContainer={false}
        forceRender={config.forceRender}
        wrapClassName={config.fullVisible && fullStatus ? "wsmodel-wrap-style wsmodal-wrap-fullscreen" : "wsmodel-wrap-style"}
        bodyStyle={{ maxHeight: ($(window).height() - 180) + 'px', overflowY: "auto", padding: "5px 15px" }}
        // style={{ top: 70 }}
        title={
          <div
            style={{
              width: '100%',
              cursor: 'move',
            }}
            onMouseOver={() => {
              if (modelState.modelMoveDisabled) {
                updateModelState({ modelMoveDisabled: false });
              }
            }}
            onMouseOut={() => {
              updateModelState({ modelMoveDisabled: true });
            }}
          >
            {config.title}
            <button
              type="button"
              className="ant-modal-close"
              style={{ right: 42 }}
              onClick={toggleFullScreen}
            >
              <span className="ant-modal-close-x">
                {config.fullVisible ? (!fullStatus ? <ArrowsAltOutlined /> : <ShrinkOutlined />) : ''}
              </span>
            </button>
          </div>
        }
        modalRender={(modal) => (
          <Draggable
            disabled={modelState.modelMoveDisabled}
            bounds={modelState.modelMoveBounds}
            onStart={(event, uiData) => handleModelMove(event, uiData)}
          >
            <div ref={draggleRef}>{modal}</div>
          </Draggable>
        )}
        visible={config.visible}
        width={config.width}
        onOk={handleSubmit}
        onCancel={handleCancel}
        keyboard={false}
        maskClosable={false}
        destroyOnClose={true}
        footer={footer()}
      // centered
      >
        {config.content}
      </Modal>
    </>)
}

export default WsModal;