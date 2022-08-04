import MdEditor from "md-editor-rt";
import "md-editor-rt/lib/style.css";
import React,{ useState } from "react";
import WsForm from "../WsForm";
import './js/iconfont-md-editor.js';
import './js/iconfont-md-editor-other.js';
import './index.less';
const NormalToolbar = MdEditor.NormalToolbar;


const WsMdEditor = (props) => {
  const [fileFormShow, setFileFormShow] = useState(false);
  const onUploadImgFunc = async (files, callback) => {
    if(props.onUploadApi){
      const res = await Promise.all(
        Array.from(files).map((file) => {
          return new Promise((rev, rej) => {
            const form = new FormData();
            form.append('file', file);
            props.onUploadApi(form)
              .then((res) => rev(res))
              .catch((error) => rej(error));
          });
        })
      );
      callback(res.map((item) => item.data.url));
    }
  }

  MdEditor.config({
    markedRenderer(renderer) {
      renderer.link = (href, title, text) => {
        return `<a href="${href}" title="${title || ''}" target="_blank">${text}</a>`;
      };
      return renderer;
    }
  });
  return (
    <>
      {fileFormShow&&<WsForm
        onCancel = {()=>{
          setFileFormShow(false);
        }}
        fields={[
          {name:"files",label:'附件上传',type:'string',compoType:'uploadFile',multiple:true,
          customRequest:props.onUploadApi,required:true},
        ]}
        onSelfSubmit={(params,cb) => {
          if(props.onUploadApi){
            let fileList = params.files.fileList;
            let text = props.modelValue;
            for(let i in fileList){
              let file = fileList[i];
              if(file.status== 'done'){
                text = text+`[${file.name}](${file.response?.url})\r\n`;
              }
            }
            props.onChange(text);
          }
          cb();
        }}
      />}
      <MdEditor
        className="md-editor-edit"
        noIconfont
        onUploadImg={onUploadImgFunc}
        toolbars={[
          'bold',
          'underline',
          'italic',
          'title',
          '-',
          'strikeThrough',
          'sub',
          'sup',
          'quote',
          'unorderedList',
          'orderedList',
          '-',
          'codeRow',
          'code',
          'link',
          'image',
          1,
          'table',
          'mermaid',
          'katex',
          '-',
          'revoke',
          'next',
          'save',
          '=',
          0,
          'fullscreen',
          'preview',
          'htmlPreview',
          'catalog',
        ]}
        defToolbars={[
          <NormalToolbar
            title="隐藏边栏"
            key="zuocebianlan"
            trigger={
              <svg className="md-icon" aria-hidden="true">
                <use xlinkHref="#icon-zuocebianlan"></use>
              </svg>
            }
            onClick={() => {
              if(props.onMenuStatus){
                props.onMenuStatus();
              }
            }}
          />,
          <NormalToolbar
            title="上传文件"
            key="uploadFile"
            trigger={
              <svg className="md-icon" aria-hidden="true">
                <use xlinkHref="#icon-attachment"></use>
              </svg>
            }
            onClick={() => {
              setFileFormShow(true);
            }}
          />
        ]}
        {...props}
      />
    </>
  );
}

export default WsMdEditor;