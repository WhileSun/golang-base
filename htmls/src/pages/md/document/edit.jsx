import React, { useState, useEffect, useMemo} from 'react';
import ReactDOM from 'react-dom';
import MdEditor from "md-editor-rt";
import "md-editor-rt/lib/style.css";
import $ from 'jquery';
import MenuTree from './components/menutree';
import { updateMdDocumentText , getMdDocumentText} from '@/services/api';
import { loadApi,getDefualtValue } from '@/utils/tools';
import { history } from 'umi';
import './index.less';
import '@static/js/iconfont-md-editor.js';
import '@static/js/iconfont-md-editor-other.js';

const editorId = 'my-editor';
const NormalToolbar = MdEditor.NormalToolbar;
var html_text = '';
const Index = (props) => {
  const book_id = props.location.query.book_id;
  const [mdText, setMdText] = useState('');
  const [htmlText, setHtmlText] = useState('');
  const [menuStatus, setMenuStatus] = useState(true);
  const [documentId, setDocumentId] = useState(getDefualtValue(localStorage.getItem('book::last:'+book_id),0));
  const [codeTheme] = useState('github')
  useEffect(()=>{
    history.listen((location, action) => {
      // window.location.reload();
    });
  },[])
  // const [catalogList, setList] = useState([]);

  const onUploadImg = async (files, callback) => {
    if (props.api !== undefined) {
      const res = await Promise.all(
        Array.from(files).map((file) => {
          return new Promise((rev, rej) => {
            const form = new FormData();
            form.append('image', file);
            requests.postFile(props.api, form)
              .then((res) => rev(res))
              .catch((error) => rej(error));
          });
        })
      );
      callback(res.map((item) => item.data.url));
    }
  }
  useEffect(()=>{
    if(documentId == 0){
      return;
    }
    localStorage.setItem('book::last:'+book_id,documentId);
    loadApi(getMdDocumentText,{document_id:documentId},(data)=>{
      setHtmlText(data.html_text);
      setMdText(data.md_text)
    })
  },[documentId])

  useEffect(() => {
    if($(`#${editorId} .md-menu-wrapper`).length ==0){
      $(`#${editorId} .md-content`).prepend(`<div class="md-menu-wrapper"></div>`);
    }
    ReactDOM.render(<MenuTree show={menuStatus} bookId={book_id} setDocumentId={(id)=>{setDocumentId(id)}} documentId={documentId}/>, document.getElementsByClassName('md-menu-wrapper')[0]);
  }, [menuStatus,documentId]);

  const onSave = ()=>{
      console.log(mdText,html_text);
      loadApi(updateMdDocumentText,{document_id:documentId,'md_text':mdText,'html_text':htmlText},()=>{
      },'保存成功')
  }

  return (
    <>
      <MdEditor
        className='md-editor-edit'
        modelValue={mdText}
        editorId={editorId}
        codeTheme={codeTheme}
        noIconfont
        onHtmlChanged = {(val)=>{
          console.log(val);
          html_text = val;
        }}
        onChange={(val) => {
          console.log(val);
          setMdText(val)
        }}
        onSave = {onSave}
        style={{ height: '100vh'}}
        onUploadImg={onUploadImg}
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
            trigger={
              <svg className="md-icon" aria-hidden="true">
                <use xlinkHref="#icon-zuocebianlan"></use>
            </svg>
            }
            onClick={()=>{
              setMenuStatus(!menuStatus);
            }}
          />
        ]}
      />
    </>
  );
};

export default Index;