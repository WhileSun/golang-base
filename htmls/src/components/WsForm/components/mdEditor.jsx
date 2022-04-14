import Editor from "md-editor-rt";
import "md-editor-rt/lib/style.css";
import React,{ useState } from "react";
import requests from '@/request/index';

const MdEditor = (props) => {
  const [text, setText] = useState(props.initMdValue);

  const onUploadImg = async (files, callback)=>{
    if(props.api !==undefined){
      const res = await Promise.all(
        Array.from(files).map((file) => {
          return new Promise((rev, rej) => {
            const form = new FormData();
            form.append('image', file);
            requests.postFile(props.api,form)
            .then((res) => rev(res))
            .catch((error) => rej(error));
          });
        })
      );
      callback(res.map((item) => item.data.url));
    }
  }
  
  return (
    <>
    <Editor
      modelValue={text}
      onChange={(modelValue) => {
        setText(modelValue)
        props.setData(modelValue)
      }}
      onUploadImg ={onUploadImg}
      toolbars={[
        'revoke',
        'next',
        '-',
        'bold',
        'underline',
        'italic',
        // '-',
        'strikeThrough',
        'title',
        // 'sub',
        // 'sup',
        'quote',
        'unorderedList',
        'orderedList',
        '-',
        'codeRow',
        'code',
        'link',
        'image',
        'table',
        // 'save',
        '-',
        // 'fullscreen',
        'pageFullscreen',
        'preview',
      ]}
    />
    </>
  );
};

export default MdEditor;
