import React, { useState, useEffect, useMemo } from 'react';
import ReactDOM from 'react-dom';
import MdEditor from "md-editor-rt";
import "md-editor-rt/lib/style.css";
import MenuTree from './components/menutree';
import { getMdDocumentText } from '@/services/api';
import { loadApi,getDefualtValue} from '@/utils/tools';
import {WsButton} from '@/components/WsTools';
import { Space,Button } from 'antd';
import { history, Link } from 'umi';
import qs from 'qs';

const editorId = 'my-editor';
const Index = (props) => {
    const book_id = props.location.query.book_id;
    const [mdText, setMdText] = useState('');
    const [menuStatus, setMenuStatus] = useState(true);
    const [documentId, setDocumentId] = useState(getDefualtValue(localStorage.getItem('book::last:'+book_id),0));
    const [documentName, setDocumentName] = useState('');
    useEffect(() => {
        history.listen((location, action) => {
            window.location.reload();
        });
    }, [])
    // const [catalogList, setList] = useState([]);
    useEffect(() => {
        if (documentId == 0 ) {
            return;
        }
        localStorage.setItem('book::last:'+book_id,documentId);
        loadApi(getMdDocumentText, { document_id: documentId }, (data) => {
            setDocumentName(data.document_name)
            setMdText(data.md_text)
        })
    }, [documentId])

    return (
        <>
            <div style={{ display: 'flex',flexDirection:'column',height:'100vh'}}>
                <div style={{height:'50px',display:'flex',justifyContent:'space-between',alignItems:'center',padding:'5px 20px',borderBottom:'1px solid #ddd'}}>
                    <div style={{fontSize:'16px',fontWeight:600,color:'#333333'}}>项目名称</div>
                    <div>
                        <Space>
                            <Link to={{pathname:'/daily/md/document/edit',search:qs.stringify({book_id:book_id})}}>
                                <WsButton title='编辑' />
                                {/* <Button danger type="primary">编辑</Button> */}
                            </Link>
                        </Space>
                    </div>
                </div>
                <div style={{ display: 'flex',flex:1,height:0}}>
                    <MenuTree show={menuStatus} bookId={book_id}
                        setDocumentId={(id) => { setDocumentId(id) }}
                        documentId={documentId}
                        onlyShow={true}
                        setDocumentName={(name) => { setDocumentName(name) }}
                    />
                    <div style={{ flex: 1, flexDirection: 'column',height:'100%',overflow:'auto'}}>
                        <div style={{ padding: '0px 20px' }}>
                            <h1 style={{ fontSize: '30px', textAlign: 'center', fontWeight: 400, margin: '0px' }}>{documentName}</h1>
                        </div>
                        <MdEditor
                            modelValue={mdText}
                            editorId={editorId}
                            previewOnly
                            style={{ height: 'auto', flex: 1, padding: '0px 20px 100px' }}
                        />
                    </div>
                </div>
            </div>
        </>
    );
};

export default Index;