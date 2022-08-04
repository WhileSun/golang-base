import React, { useState, useEffect, useMemo } from 'react';
import ReactDOM from 'react-dom';
import $ from 'jquery';
import MenuTree from './components/menutree';
import { updateMdDocumentText, getMdDocumentText,uploadDocumentNameFile} from '@/services/api';
import { loadApi, getDefualtValue } from '@/utils/tools';
import { history } from 'umi';
import {WsMdEditor} from '@/components/WsTools';

const editorId = 'my-editor';
var html_text = '';
const Index = (props) => {
  const book_id = props.location.query.book_id;
  const [mdText, setMdText] = useState('');
  const [menuStatus, setMenuStatus] = useState(true);
  const [documentId, setDocumentId] = useState(getDefualtValue(localStorage.getItem('book::last:' + book_id), 0));
  useEffect(() => {
    history.listen((location, action) => {
      window.location.reload();
    });
  }, [])

  useEffect(() => {
    if (documentId == 0) {
      return;
    }
    localStorage.setItem('book::last:' + book_id, documentId);
    loadApi(getMdDocumentText, { document_id: documentId }, (data) => {
      setMdText(data.md_text)
    })
  }, [documentId])

  useEffect(() => {
    if ($(`#${editorId} .md-menu-wrapper`).length == 0) {
      $(`#${editorId} .md-content`).prepend(`<div class="md-menu-wrapper"></div>`);
    }
    ReactDOM.render(<MenuTree show={menuStatus} bookId={book_id} setDocumentId={(id) => { setDocumentId(id) }} documentId={documentId} />, document.getElementsByClassName('md-menu-wrapper')[0]);
  }, [menuStatus, documentId]);

  const onSave = () => {
    loadApi(updateMdDocumentText, { document_id: documentId, 'md_text': mdText, 'html_text': html_text }, () => {
    }, '保存成功')
  }

  return (
    <>
      <WsMdEditor
        modelValue={mdText}
        editorId={editorId}
        onHtmlChanged={(val) => {
          html_text=val;
        }}
        onChange={(val) => {
          setMdText(val)
        }}
        onSave={onSave}
        style={{ height: '100vh' }}
        onUploadApi={uploadDocumentNameFile}
        onMenuStatus={()=>{
          setMenuStatus(!menuStatus);
        }}
      />
    </>
  );
};

export default Index;