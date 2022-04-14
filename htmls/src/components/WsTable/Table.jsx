import requests from '@/request/index';
import './index.less';
import $ from 'jquery';
import { Table, Form, message, Button, Tooltip, Divider } from 'antd';
import React, { useState, useEffect, useMemo, useRef, useImperativeHandle } from 'react';
import { SettingOutlined, ReloadOutlined, SearchOutlined ,MinusSquareOutlined,PlusSquareOutlined } from '@ant-design/icons';
import { paramIsset, funcIsset, getRandStr, toTree, parseSearchParams, setFormParamStore, arrayColumn,filterName} from './utils/tools';
import useTable from './hooks/useTable';
import { WsModal } from '@/components/WsPopup';
import SearchForm from './components/searchForm';
import SearchButton from './components/searchButton';
import ColumnShow from './components/columnShow';
import initColumnFunc from './func/initColumn';
import initShowTreeFunc from './func/initShowTree';

const WsTable = (props, ref) => {
  const [formRef] = Form.useForm();
  const initialize = () => {
    console.log('initialize');
    let objs = {};
    objs.th = paramIsset(props.th, []);
    objs.size = paramIsset(props.size, 'small');
    objs.data = paramIsset(props.data, []);
    objs.rowKey = paramIsset(props.rowKey, 'id');
    objs.display = paramIsset(props.display, 'fixed');
    objs.api = paramIsset(props.api, '');
    objs.type = paramIsset(props.type, 'normal');
    objs.checkboxField = paramIsset(props.checkbox, false);
    objs.treeTable = paramIsset(props.treeTable, false); //是不是tree table
    objs.footHeight = 43; //底部宽度
    objs.chooseIds = []; //check选中的数据
    objs.divId = getRandStr('table');
    return objs;
  };
  const [state, setState] = useState(initialize);
  const [apiresp, setApiresp] = useState({});
  const [apiData, setApiData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [modalShow, setModalShow] = useState(paramIsset(props.modalShow, true)); //弹出框是否显示
  //tree table
  const [expandedRowKeys,setExpandedRowKeys] = useState([]);
  const [rowKeys,setRowKeys] = useState([]);
  const [treeTableshow, setTreeTableshow] = useState(true);

  //列设置
  const initShowTree = useMemo(() => { return initShowTreeFunc(state.th) }, []);
  const [showColumns, setShowColumns] = useState(initShowTree.allKeys);
  const [searchFormShow, setSearchFormShow] = useState(true);
  //table column等配置信息
  const tableSetting = useMemo(() => {
    return initColumnFunc(state.th, state.display, showColumns)
  }, [showColumns]);

  const defaultFormParams = { 'pageSize': 50, 'page': 1, ...props.params }; //默认参数
  const [params, setParams] = useState(defaultFormParams);
  //设置api接口的参数
  const setNewFormParams = (newParam, type = "") => {
    let objs = {};
    if (type == 'changePage') {
      objs = { ...params, ...newParam };
    } else if (type == 'submitForm') {
      //每次传入提交的参数不一样，所以在需要在基础上叠加
      let newDefaultFormParams = { ...defaultFormParams, page: params.page };
      objs = { ...newDefaultFormParams, ...newParam };
    } else if (type == 'initForm') {
      objs = { ...defaultFormParams, ...newParam };
    }
    setParams(objs);
    return objs;
  }

  //初始化store保存的参数
  const setInitformParams = () => {
    let params = setNewFormParams((parseSearchParams(props.store)), 'initForm');
    getData(params);
  }

  //form查询重置
  const handleFormReset = () => {
    console.log('handleFormReset');
    setFormParamStore(props.store, {});
    formRef.resetFields();
    let params = setNewFormParams({}, 'initForm');
    getData(params);
  };

  //for查询提交
  const handleFormSubmit = (values) => {
    console.log('handleFormSubmit');
    //针对本地数据筛选，一般针对tree table使用
    if(props.onLocalFilter){
      props.onLocalFilter(parseSearchParams(values))
    }else{
      setFormParamStore(props.store, values);
      let params = setNewFormParams(parseSearchParams(values), 'submitForm');
      getData(params);
    }
  };

  useEffect(() => {
    setInitformParams();
    //初始化设置字段数据
    formRef.setFieldsValue(props.store);
  }, []);

  //换页
  const tableChangePage = (page) => {
    let params = setNewFormParams({ page: page }, 'changePage');
    getData(params);
  };

  const tableReload = () => {
    getData(params);
  }
  //获取table设置展示类型的数据
  const getTableData = (data)=>{
    if(state.treeTable){
      setRowKeys(arrayColumn(data,"id"));
      return toTree(data);
    }
    return data;
  }

  const getData = async (apiParams = {}) => {
    setLoading(true);
    console.log('getData', apiParams);
    requests.post(state.api, apiParams).then(function (resp) {
      console.log('resp', resp);
      setApiresp(resp);
      setApiData(getTableData(resp.data));
      setLoading(false);
      if (resp.code != 0) {
        message.error(resp.msg);
      }
    }).catch(function (error) {
      setLoading(false);
      message.error("列表获取异常，请联系管理员处理！");
      console.log('error', error);
    });

  };


  //初始话查询表单
  const headerForm = useMemo(() => {
    return (
      <>
        <SearchForm
          formRef={formRef}
          searchs={props.searchs}
          show={searchFormShow}
          handleFormSubmit={handleFormSubmit}
        />
      </>
    );
  }, [searchFormShow]);

  //表单按钮
  const headerButton = useMemo(() => {
    const icon = { fontSize: '16px', marginRight: '15px' };
    const treeFunc = ()=>{
      if(treeTableshow){
        setExpandedRowKeys(rowKeys);
      }else{
        setExpandedRowKeys([]);
      }
      setTreeTableshow(!treeTableshow)
    }
    return (
      <>
        <SearchButton
          formRef={formRef}
          searchs={props.searchs}
          btns={props.btns}
          loading={loading}
          handleFormReset={handleFormReset}
          soltRightBtn={(
            <>
              {state.treeTable?
              <Tooltip placement="top" title='Tree Table 展开/隐藏'>
                {treeTableshow ? <PlusSquareOutlined style={icon} onClick={treeFunc} />:
                <MinusSquareOutlined  style={icon} onClick={treeFunc} />}
              </Tooltip>
              :""}
              <Tooltip placement="top" title='搜索框显示/隐藏'>
                <SearchOutlined style={icon} onClick={() => {
                  setSearchFormShow(!searchFormShow);
                }} />
              </Tooltip>
              <Tooltip placement="top" title='刷新'>
                <ReloadOutlined style={icon} onClick={tableReload} />
              </Tooltip>
              <ColumnShow
                data={initShowTree}
                showColumns={showColumns}
                setShowColumns={(keys) => { setShowColumns(keys) }}
                solt={(<Tooltip placement="top" title='列设置'><SettingOutlined style={{ fontSize: '16px' }} /></Tooltip>)}
              />
            </>
          )}
        />
      </>
    );
  }, [loading, searchFormShow, showColumns,treeTableshow]);

  useEffect(() => {
    if (state.display == 'fixed') {
      normalResize(state.divId,state.footHeight);
    }
  },[searchFormShow]);

  //界面自适应
  useEffect(() => {
    if (state.display == 'fixed') {
      if (state.type == 'normal') {
        function resize() {
          normalResize(state.divId, state.footHeight)
        }
        $(window).on('resize', resize);
        // normalResize(state.divId, state.footHeight);
        return () => {
          console.log('clear');
          $(window).off('resize', resize);
        }
      } else if (state.type == 'modal') {
        setTimeout(() => {
          modalResize(state.divId, state.footHeight);
        }, 10);
      }
    }
  }, []);

  var tableInstance = useTable(props.table);
  tableInstance.reload = () => { tableReload(); };
  tableInstance.getDataList = () => { return apiresp.data};
  tableInstance.filterName = (dataIndex,val)=>{
    if(val){
      setApiData(filterName(getTableData(apiresp.data),dataIndex,val))
    }else{
      setApiData(getTableData(apiresp.data))
    }
  };
  tableInstance.getChooseIds = () => { return state.chooseIds; };
  //支持原生ref
  useImperativeHandle(ref, () => { return tableInstance });

  const tableBody = (
    <>
      <div className={state.divId + " ws-table"}>
        <div className="ws-table-header">
          {headerForm}
          {headerButton}
        </div>
        <div className="ws-table-container">
          <Table
            ref={props.ref}
            rowSelection={state.checkboxField ? {
              type: 'checkbox',
              onChange: (selectedRowKeys) => {
                setState({ ...state, 'chooseIds': selectedRowKeys });
              }
            } : ""}
            columns={tableSetting.columns}
            dataSource={apiData}
            bordered={true}
            size={state.size}
            rowKey={state.rowKey}
            scroll={tableSetting.tableScroll}
            loading={loading}
            showSizeChanger={false}
            // expandRowByClick={true}
            onExpandedRowsChange={(expandedRows)=>{
              setExpandedRowKeys(expandedRows);
              console.log('onExpandedRowsChange',expandedRows);
            }}
            expandedRowKeys={expandedRowKeys}
            pagination={{
              position: ['bottomRight'],
              pageSize: params.pageSize,
              total: apiresp.totalSize,
              pageSizeOptions: [],
              current: params.page,
              size: 'small',
              showTotal: (total) => {
                return `共 ${total} 条`;
              },
            }}
            onChange={(paginate, filters, sorter, extra) => {
              if (extra.action === 'paginate') {
                tableChangePage(paginate.current);
              }
            }}
            // onRow={(record, index) => {
            //   return {
            //     onClick: event => {
            //       // console.log(event);
            //       if (event.target.dataset.act !== undefined && props.rowBtnsClick !== undefined) {
            //         props.rowBtnsClick(event.target.dataset.act, record);
            //       }
            //     },
            //   };
            // }}
          />
        </div>
      </div>
    </>
  );

  if (state.type == 'modal') {
    return (
      <WsModal
        content={tableBody}
        show={modalShow}
        width={paramIsset(props.width, 800)}
        title={paramIsset(props.title, '列表')}
        cancel={() => { setModalShow(false); if (props.cancel()) { props.cancel() }; }}
      />
    )
  } else {
    return tableBody;
  }
};

//常规fiexd
const normalResize = (divId, footHeight) => {
  if (checkTableShow(divId)) {
    let elem = $('.' + divId);
    let thtop = elem.find('.ant-table-body:first').offset().top;
    elem.find('.ant-table-body').height($(window).height() - thtop - footHeight);
  } else {
    setTimeout(() => {
      normalResize(divId, footHeight)
    }, 100);
  }
}

const checkTableShow = (divId) => {
  let elem = $('.' + divId);
  if (elem.find('.ant-table-body').offset() === undefined) {
    return false;
  } else {
    return true;
  }
}

//模态框的fiexd
const modalResize = (divId, footHeight) => {
  let elem = $('.' + divId);
  let thtop = elem.find('.ws-table-header').height();
  let thead = elem.find('.ant-table-header thead').height();
  elem.find('.ant-table-body').height($(window).height() - 180 - thtop - thead - 20 - footHeight);
}

export default WsTable;
