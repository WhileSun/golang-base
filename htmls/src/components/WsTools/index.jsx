import WsModal from './WsModal';
import WsDrawer from './WsDrawer';
import WsButton from './WsButton';
import WsTable  from './WsTable';
import WsForm   from './WsForm';
import WsIcon from './WsIcon';
import WsMdEditor from './WsMdEditor';
import { forwardRef, useState, useEffect } from 'react';
import $ from 'jquery';
import ReactDOM from 'react-dom';

const ModalConfirmHtml = ({ config }) => {
    const [visible, setVisible] = useState(true);
    const [loading, setLoading] = useState(false);
    if(config.mode == 'danger'){
        config.content = <div style={{color:'red'}}>{config.content}</div>
    }
    
    const removeModal = () => {
        setLoading(false)
        setVisible(false);
        $(`body .wsmodal-confirm`).remove();
    }

    const onOk = async () => {
        setLoading(true)
        if(typeof config.onOk === "function"){
            setTimeout(()=>{
                config.onOk.call(this);
                removeModal();
            },300)
        }else{
            if (config.onOk instanceof Promise) {
                await config.onOk.then();
            } 
            removeModal();
        }
    }
    
    const onCancel = () => {
        if (config.onCancel) {
            config.onCancel.call(this);
        }
        setVisible(false);
        removeModal();
    }
    return (
        <>
            <WsModal visible={visible} title={config.title} content={config.content}
                loading={loading} onOk={onOk} onCancel={onCancel} width={400} />
        </>
    );
}

const modalConfirm = (config) => {
    if ($(`body .wsmodal-confirm`).length == 0) {
        $('body').append(`<div><div class="wsmodal-confirm"></div></div>`);
    }
    ReactDOM.render(<ModalConfirmHtml config={config} />, document.getElementsByClassName('wsmodal-confirm')[0]);
}

WsModal.confirm = modalConfirm;

export { WsModal, WsDrawer, WsButton, WsTable, WsForm, WsIcon, WsMdEditor};
