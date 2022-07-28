import React, { useState, useEffect, useMemo } from 'react';
import { PlusOutlined, DeleteOutlined, FormOutlined, MenuOutlined} from '@ant-design/icons';
import { Tree, Button, Tabs,Tooltip} from 'antd';
import './index.less';
import WsForm from '@/components/WsForm';
import { toTree, loadApi, transTree, getTreeKeys } from '@/utils/tools';
import { getMdDocumentNameList, addMdDocumentName, updateMdDocumentName, deleteMdDocumentName, dragMdDocumentName } from '@/services/api';
import { WsModal } from '@/components/WsTools';

const { DirectoryTree } = Tree;
const { TabPane } = Tabs;

const MenuTree = (props) => {

  const [ulStyle, setUlStyle] = useState();
  const [ulShow, setUlShow] = useState(false);
  const formRef = WsForm.useForm();
  const [formData, setFormData] = useState({});
  const [formShow, setFormShow] = useState(false);
  const [treeData, setTreeData] = useState([]);
  const [treeNode, setTreeNode] = useState({});
  var book_id = props.bookId;

  const getInitDocumentList = (init = false) => {
    loadApi(getMdDocumentNameList, { book_id: book_id }, (data) => {
      if (init) {
        if (data.length > 0 && props.documentId == 0) {
          props.setDocumentId(data[0]['id']);
        }
      }
      setTreeData(transTree(toTree(data), 'document_name', 'id'));
    });
  }

  useEffect(() => {
    const closeClickrightMenu = () => {
      if (ulShow) {
        setUlShow(false);
      }
    }
    window.addEventListener('click', closeClickrightMenu);
    return () => {
      window.removeEventListener('click', closeClickrightMenu);
    }
  }, [ulShow])

  useEffect(() => {
    getInitDocumentList(true);
  }, []);

  const onSelect = (selectedKeys, info) => {
    props.setDocumentId(selectedKeys[0]);
  };

  const showForm = (data = {}) => {
    setFormData(data);
    setFormShow(true);
  }

  return (
    <>
      <div style={{ display: (props.show ? 'block' : 'none') }}>
        <Tabs type="card" tabBarGutter={0}
          style={{ background: '#FAFAFA',width:'250px',height:'100vh',borderRight:'1px solid #ddd'}}
          tabBarExtraContent={{
            right:(props.onlyShow?'':<Tooltip placement="right" title='创建文档'>
              <div style={{cursor: 'pointer',fontWeight:'bold',fontSize:'14px',paddingRight:'10px'}}  onClick={() => { showForm({ parent_id: 0 }) }}><PlusOutlined/></div>
              </Tooltip>)
          }}
        >
          <TabPane tab={<span style={{fontSize:'14px',color:'#666'}}>{props.onlyShow?'目录':'文档'}</span>} key="1">
            <div className="menu-tree">
              <DirectoryTree
                icon={false}
                expandAction={false}
                showIcon={false}
                blockNode
                draggable={{ icon: false }}
                onDrop={({ dragNode, dropPosition, dropToGap, node }) => {
                  console.log(dragNode, dropPosition, dropToGap, node);
                  const dropPos = node.pos.split('-');
                  const dragPosition = dropPosition - Number(dropPos[dropPos.length - 1]);
                  const dragNodeId = dragNode.key;
                  const nodeId = node.key;
                  console.log(nodeId);
                  loadApi(dragMdDocumentName, {
                    'drag_node_id': dragNodeId,
                    'node_id': nodeId,
                    'drag_position': dragPosition,
                    'drag_gap': dropToGap
                  }, () => {
                    getInitDocumentList();
                  }, '保存顺序成功')
                }}
                defaultExpandedKeys={[]}
                selectedKeys={[props.documentId.toString()]}
                onSelect={onSelect}
                treeData={treeData}
                titleRender={(dom) => { return <div style={{ height: '30px', lineHeight: '30px' }}>{dom.title}</div> }}
                onRightClick={({ event, node }) => {
                  setTreeNode(node);
                  setUlShow(true)
                  setUlStyle({
                    position: 'absolute',
                    zIndex: 1000,
                    width: '100px',
                    height: 'auto',
                    background: '#f6f6f6',
                    border: '1px solid #d6d6d6',
                    boxShadow: '0 0 8px rgb(99 99 99 / 30%)',
                    left: event.clientX + 'px',
                    top: event.clientY + 'px'
                  })
                }}
              >
              </DirectoryTree>
              {props.onlyShow ? '':<div style={{ display: (ulShow ? 'block' : 'none'), ...ulStyle }} className="menu-tree-clickright">
                <ul>
                  <li onClick={() => { showForm({ parent_id: treeNode.key }) }}><span><PlusOutlined /></span>添加</li>
                  <li onClick={() => { showForm({ id: treeNode.key, document_name: treeNode.title }) }}><span><FormOutlined /></span>编辑</li>
                  <li onClick={() => {
                    var keys = getTreeKeys([treeNode]);
                    WsModal.confirm(
                      {
                        mode: 'danger',
                        content: '确定要删除该文档吗?',
                        onOk: () => {
                          loadApi(deleteMdDocumentName, { ids: keys }, () => {
                            getInitDocumentList();
                          }, '删除成功')
                        }
                      }
                    );
                  }}><span><DeleteOutlined /></span>删除</li>
                </ul>
              </div>}
            </div>
          </TabPane>
          {/* <TabPane tab="搜索" key="2">
            Content of tab 2
          </TabPane> */}
        </Tabs>
      </div>
      {formShow && <WsForm
        form={formRef}
        width={500}
        onCancel={() => {
          setFormShow(false);
        }}
        data={formData}
        fields={[
          { name: "document_name", col: 24, label: '文档名称', compoType: 'input', required: true },
          { name: "document_ident", col: 24, label: '文档标识', compoType: 'input', required: false, remove: !!formData.id },
        ]}
        api={addMdDocumentName}
        updateApi={updateMdDocumentName}
        onBeforeSubmit={(params, cb) => {
          params.book_id = book_id;
          params.parent_id = formData.parent_id;
          cb();
        }}
        onSucc={() => {
          getInitDocumentList();
        }}
      />}
    </>
  );
};

export default MenuTree;