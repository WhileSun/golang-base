import {paramIsset} from '../utils/tools';
import React, { useState} from 'react';
import Draggable from 'react-draggable';
import {Modal,Button} from 'antd';
import $ from 'jquery';
import {ArrowsAltOutlined,ShrinkOutlined} from '@ant-design/icons';
import './index.less'

const WsModal = (props)=>{
    //model参数
  const stateObj = {
    modelMoveDisabled: true,
    modelMoveBounds: { left: 0, top:0, bottom: 0, right: 0 },
  };
  const [modelState, setModelState] = useState(stateObj);
  const [confirmLoading, setConfirmLoading] =useState(false);
  const [fullStatus,setFullStatus] = useState(paramIsset(props.fullStatus,false));

  const updateModelState = (newState) => {
    setModelState({...modelState,...newState})
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
  const handleCancel = (e) => {
    if(props.cancel){props.cancel()}
  };

  const handleSubmit = () => {
    console.log('handleSubmit');
    if(props.submit){props.submit()}
  };
  
  const footer = ()=>{
    let btns = [<Button key="back" onClick={handleCancel}>取消</Button>];
    if(props.submit){
      btns.push(<Button key="submit" type="primary" onClick={handleSubmit} loading={props.modelLoading}>提交</Button>);
    }
    return btns;
  }

  const toggleFullScreen = ()=>{
    setFullStatus(!fullStatus);
  }
  return (
    <>
      <Modal
        getContainer={false} 
        forceRender={paramIsset(props.forceRender, false)}
        wrapClassName = {fullStatus?"model-wrap-style modal-wrap-fullscreen":"model-wrap-style"}
        bodyStyle = {{maxHeight:($(window).height()-180)+'px',overflowY: "auto", padding:"5px 15px"}}
        style={{ top: 70 }}
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
            {paramIsset(props.title, '提示')}
            <button
              type="button"
              className="ant-modal-close"
              style={{ right: 42 }}
              onClick={toggleFullScreen}
            >
              <span className="ant-modal-close-x">
                {!fullStatus?<ArrowsAltOutlined />:<ShrinkOutlined/>}
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
        visible={paramIsset(props.show, false)}
        width={paramIsset(props.width, 600)}
        onOk={handleSubmit}
        onCancel={handleCancel}
        keyboard={false}
        maskClosable={false}
        destroyOnClose={true}
        footer={footer()}
        // centered
      >
        {props.content}
      </Modal>
    </>)
}

export default WsModal;