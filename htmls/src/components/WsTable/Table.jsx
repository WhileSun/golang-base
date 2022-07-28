import './index.less';
import $ from 'jquery';
import { Table, Form, message, Button, Tooltip} from 'antd';
import React, { useState, useEffect, useMemo, useRef, useImperativeHandle } from 'react';
import { SettingOutlined, ReloadOutlined, SearchOutlined, MinusSquareOutlined, PlusSquareOutlined } from '@ant-design/icons';
import { paramIsset, getRandStr, toTree, parseFormParams, setFormParamStore, arrayColumn, filterName } from './utils/tools';
import useTable from './hooks/useTable';
import { WsModal } from '@/components/WsTools';
import HeaderSearchForm from './components/headerSearchForm';
import HeaderButtonLeft from './components/headerButtonLeft';
import ColumnShowButton from './components/columnShowButton';
import initColumnFunc from './func/initColumn';
import initShowColumnFunc from './func/initShowColumn';
import {normalResize,modalResize} from './func/initTableTool';
const WsTable = (props, ref) => {
  const [formRef] = Form.useForm();
  
  //父级设置配置
  const config = useMemo(() => {
    let param = {};
    param.th = paramIsset(props.th); //列表展示字段
    param.size =  paramIsset(props.size, 'small'); //尺寸
    param.rowKey =paramIsset(props.rowKey, 'id'); //主键ID，选中返回的参数
    param.display = paramIsset(props.display, 'fixed'); //展示方式fixed:固定高度;fluid:不固定
    param.api = paramIsset(props.api, new Promise((vals)=>{})); //promise api
    param.mode = paramIsset(props.mode, 'normal'); //模式: normal modal
    param.checkboxField = paramIsset(props.checkbox, false); //是否开启选择框
    param.treeTable = paramIsset(props.treeTable, false); //tree table开启
    param.footHeight = paramIsset(props.footHeight, 43); //fixed展示，底部的高度
    param.divId = paramIsset(props.divId, getRandStr('table'));  //table的ID
    param.btns = paramIsset(props.btns, []); // 按钮
    param.searchs = paramIsset(props.searchs, []); //搜索框

    param.data = paramIsset(props.data, []); //table直接赋值
    param.store = props.store;
    param.otherFormParams = props.params; //额外参数
    param.onLocalFilter = props.onLocalFilter; //对本地数据筛选

    param.page =  paramIsset(props.page, 1); //第几页
    param.pageSize =  paramIsset(props.pageSize, 50); //每页数
    return param;
  }, [props]);

  const [apiresp, setApiresp] = useState({}); //api原生数据
  const [apiData, setApiData] = useState([]); //api经过内部转化数据
  const [loading, setLoading] = useState(false);
  const [checkedIds,setCheckedIds] = useState([]); //表格选中的ID
  const [modalShow, setModalShow] = useState(paramIsset(props.modalShow, true)); //弹出框是否显示
  //tree table
  const [expandedRowKeys, setExpandedRowKeys] = useState([]); //展开的节点ID
  const [rowKeys, setRowKeys] = useState([]); //当前所有节点ID
  const [treeTableshow, setTreeTableshow] = useState(true);  //是否是树形表
  //列设置
  const initShowColumn = useMemo(() => { return initShowColumnFunc(config.th) }, []); //表格字段的字段和值
  const [showColumns, setShowColumns] = useState(initShowColumn.allKeys); //表格展示的字段,默认全部
  const [searchFormShow, setSearchFormShow] = useState(true); //是否显示搜索框
  //table column等配置信息
  const tableSetting = useMemo(() => {
    return initColumnFunc(config.th, config.display, showColumns)
  }, [showColumns]);

  /**查询参数整合 */
  const defaultFormParams = { 'pageSize': config.pageSize, 'page': config.page, ...config.otherFormParams}; //默认参数
  const [formParams, setFormParams] = useState(defaultFormParams);
  const setFormParamsFunc = (newParam, mode = "") => {
    let params = {};
    if (mode == 'changePage') {
      params = { ...formParams, ...newParam };
    } else if (mode == 'submitForm') {
      //每次传入提交的参数不一样，所以在需要在基础上叠加
      let newDefaultFormParams = { ...defaultFormParams, page: formParams.page };
      params = { ...newDefaultFormParams, ...newParam };
    } else if (mode == 'initForm') {
      params = { ...defaultFormParams, ...newParam };
    }
    setFormParams(params);
    return params;
  }

  //初始化store保存的参数
  const setInitformParams = () => {
    let params = setFormParamsFunc((parseFormParams(config.store)), 'initForm');
    getData(params);
  }

  //form查询重置
  const handleFormReset = () => {
    console.log('handleFormReset');
    setFormParamStore(config.store, {});
    formRef.resetFields();
    let params = setFormParamsFunc({}, 'initForm');
    getData(params);
  };

  //for查询提交
  const handleFormSubmit = (values) => {
    console.log('handleFormSubmit');
    //针对本地数据筛选，一般针对tree table使用
    if (config.onLocalFilter) {
      config.onLocalFilter(parseFormParams(values))
    } else {
      setFormParamStore(config.store, values);
      let params = setFormParamsFunc(parseFormParams(values), 'submitForm');
      getData(params);
    }
  };

  useEffect(() => {
    setInitformParams();
    //初始化设置字段数据
    formRef.setFieldsValue(config.store);
  }, []);

  //换页
  const tableChangePage = (page) => {
    let params = setFormParamsFunc({ page: page }, 'changePage');
    getData(params);
  };

  //刷新
  const tableReload = () => {
    getData(formParams);
  }

  //转化table的数据类型
  const transTableData = (data) => {
    if (config.treeTable) {
      setRowKeys(arrayColumn(data, "id"));
      return toTree(data);
    }
    return data;
  }

  const getData = async (apiParams = {}) => {
    setLoading(true);
    config.api.call(this,apiParams).then(function (resp) {
      console.log('resp', resp);
      setApiresp(resp);
      setApiData(transTableData(resp.data));
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

  //初始化查询表单
  const headerSearchForm = useMemo(() => {
    return (
      <>
        <HeaderSearchForm
          formRef={formRef}
          searchs={config.searchs}
          handleFormSubmit={handleFormSubmit}
        />
      </>
    );
  }, []);

  //头部左边按钮
  const headerButtonLeft = useMemo(() => {
    const fieldLen = Object.keys(paramIsset(config.searchs, [])).length;
    return (
      <>
        <HeaderButtonLeft btns={config.btns} />
        {fieldLen > 0 ? (
          <>
            <Button htmlType="button" onClick={handleFormReset} style={{ marginRight: '10px' }} loading={loading}>
              重置
            </Button>
            <Button type="primary" onClick={() => { formRef.submit(); }} loading={loading}>
              查询
            </Button>
          </>
        ) : (
          ''
        )}
      </>
    );
  }, [loading])

  //头部右边按钮
  const headerButtonRight = useMemo(() => {
    const iconStyle = { fontSize: '16px', marginRight: '15px' };
    const treeFunc = () => {
      if (treeTableshow) {
        setExpandedRowKeys(rowKeys);
      } else {
        setExpandedRowKeys([]);
      }
      setTreeTableshow(!treeTableshow)
    }
    return (
      <>
        {config.treeTable ?
          <Tooltip placement="top" title='Tree Table 展开/隐藏'>
            {treeTableshow ? <PlusSquareOutlined style={iconStyle} onClick={treeFunc} /> :
              <MinusSquareOutlined style={iconStyle} onClick={treeFunc} />}
          </Tooltip>
          : ""}
        <Tooltip placement="top" title='搜索框显示/隐藏'>
          <SearchOutlined style={iconStyle} onClick={() => {
            setSearchFormShow(!searchFormShow);
          }} />
        </Tooltip>
        <Tooltip placement="top" title='刷新'>
          <ReloadOutlined style={iconStyle} onClick={tableReload} />
        </Tooltip>
        <ColumnShowButton
          data={initShowColumn}
          showColumns={showColumns}
          setShowColumns={(keys) => { setShowColumns(keys) }}
          solt={(<Tooltip placement="top" title='列设置'><SettingOutlined style={{ fontSize: '16px' }} /></Tooltip>)}
        />
      </>
    )
  }, [searchFormShow, showColumns, treeTableshow, rowKeys])

  useEffect(() => {
    if (config.display == 'fixed') {
      normalResize(config.divId, config.footHeight);
    }
  }, [searchFormShow]);

  //界面自适应
  useEffect(() => {
    if (config.display == 'fixed') {
      if (config.mode == 'normal') {
        function resize() {
          normalResize(config.divId, config.footHeight)
        }
        $(window).on('resize', resize);
        return () => {
          console.log('clear');
          $(window).off('resize', resize);
        }
      } else if (config.mode == 'modal') {
        setTimeout(() => {
          modalResize(config.divId, config.footHeight);
        }, 10);
      }
    }
  }, []);

  //映射ref函数
  var tableInstance = useTable(props.table);
  tableInstance.reload = () => { tableReload(); };
  tableInstance.getDataList = () => { return apiresp.data };
  tableInstance.filterName = (dataIndex, val) => {
    if (val) {
      setApiData(filterName(transTableData(apiresp.data), dataIndex, val))
    } else {
      setApiData(transTableData(apiresp.data))
    }
  };
  tableInstance.getCheckedIds = () => { return checkedIds; };
  // //支持原生ref
  useImperativeHandle(ref, () => { return tableInstance });

  const tableHtml = (
    <>
      <div className={config.divId + " ws-table"}>
        <div className="ws-table-header">
          <div className="header-search-form" style={{ display: searchFormShow ? "block" : 'none' }}>
            {headerSearchForm}
          </div>
          <div className="header-button">
            <div>{headerButtonLeft}</div>
            <div>{headerButtonRight}</div>
          </div>
        </div>
        <div className="ws-table-container">
          <Table
            ref={props.ref}
            rowSelection={config.checkboxField ? {
              type: 'checkbox',
              onChange: (selectedRowKeys) => {
                setCheckedIds(selectedRowKeys);
              }
            } : ""}
            columns={tableSetting.columns}
            dataSource={apiData}
            bordered={true}
            size={config.size}
            rowKey={config.rowKey}
            scroll={tableSetting.tableScroll}
            loading={loading}
            showSizeChanger={false}
            // expandRowByClick={true}
            onExpandedRowsChange={(expandedRows) => {
              setExpandedRowKeys(expandedRows);
            }}
            expandedRowKeys={expandedRowKeys}
            pagination={{
              position: ['bottomRight'],
              pageSize: formParams.pageSize,
              total: apiresp.totalSize,
              pageSizeOptions: [],
              current: formParams.page,
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

  //输出样式
  if (config.mode == 'modal') {
    return (
      <WsModal
        content={tableHtml}
        show={modalShow}
        width={paramIsset(props.width, 800)}
        title={paramIsset(props.title, '列表')}
        cancel={() => { setModalShow(false); if (props.cancel()) { props.cancel() }; }}
      />
    )
  } else {
    return tableHtml;
  }
};

export default WsTable;
