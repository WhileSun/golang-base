import {
  Form,
  Col,
  Input,
  Select,
  DatePicker,
  Space,
  Cascader,
  InputNumber,
  Radio,
  Tree,
  Checkbox,
  Transfer
} from 'antd';
import { LockOutlined } from '@ant-design/icons';
import MdEditor from './mdEditor';
const { Option } = Select;
const { TextArea } = Input;
import { delArrVal,paramIsset, createFormRules, formFieldTrans } from '../utils/tools';
import React, { useState} from 'react';

const ItemField = (props) => {
  const field = paramIsset(props.field, {});
  const index = props.index;
  const formRef = props.formRef;
  const initData = props.initData;

  if (field.remove === true) {
    return '';
  }
  const rules = createFormRules(field);
  const className = paramIsset(field['className']) + ' form-item-class';
  let htmls = '';
  const listData = field.listData;
  if (field.compoType === 'input') {
    htmls = (
      <Input
        placeholder={
          field.placeholder === undefined
            ? '请输入' + field.label
            : field.placeholder
        }
        type={field.type}
        disabled={field.disabled}
        onChange={field.onChange}
      />
    );
  } else if (field.compoType === 'inputPasswd') {
    htmls = (
      <Input.Password
        prefix={<LockOutlined />}
        type="password"
        placeholder="密码"
      />
    );
  } else if (field.compoType === 'text') {
    htmls = (
      <span>
        {initData[field.name]}
      </span>
    );
  } else if (field.compoType === 'inputNumber') {
    htmls = (
      <InputNumber
        placeholder={
          field.placeholder === undefined
            ? '请输入' + field.label
            : field.placeholder
        }
        onChange={field.onChange}
        style={{ width: '100%' }}
      />
    );
  } else if (field.compoType === 'textarea') {
    //maxLength 最大长度
    htmls = (
      <TextArea placeholder={field.placeholder} maxLength={field.maxLen} rows={field.rows} />
    );
  } else if (field.compoType === 'radio') {
    htmls = (
      <Radio.Group onChange={field.onChange}>
        {Object.keys(listData).map((radioKey, radioIndex) => {
          return (
            <Radio
              value={radioKey}
              key={radioIndex.toString()}
              disabled={field.disabled}
            >
              {listData[radioKey]}
            </Radio>
          );
        })}
      </Radio.Group>
    );
  } else if (field.compoType === 'select') {
    htmls = (
      <Select onChange={field.onChange} disabled={field.disabled}>
        {field.defaultOption !== undefined ? (
          <Option value="">{field.defaultOption}</Option>
        ) : (
          ''
        )}
        {Object.keys(listData).map((selectKey, selectIndex) => {
          return (
            <Option value={selectKey} key={selectIndex.toString()}>
              {listData[selectKey]}
            </Option>
          );
        })}
      </Select>
    );
  } else if (field.compoType === 'searchSelect') {
    // listData { label, value }[]
    htmls = (
      <Select
        showSearch
        showArrow
        mode={field.mode} //multiple | tags
        placeholder={field.placeholder}
        optionFilterProp="label"
        onChange={field.onChange}
        options={listData}
      />
    );
  } else if (field.compoType === 'date') {
    htmls = <DatePicker style={{ width: '100%' }} />;
  } else if (field.compoType === 'datetime') {
    htmls = <DatePicker showTime style={{ width: '100%' }} />;
  } else if (field.compoType === 'br') {
    //换行 col可以自定义 {compoType:'br',col:24}
  } else if (field.compoType === 'mdEditor') {
    const initMdValue =
      initData[field.name] !== undefined ? initData[field.name] : '';
    const [mdEditorValue, setMdEditorValue] = useState('');
    const setFormData = (value) => {
      setMdEditorValue(value);
      const vals = {};
      vals[field.name] = value;
      formRef.setFieldsValue(vals);
    };
    htmls = (
      <MdEditor
        initMdValue={initMdValue}
        setData={setFormData}
        api={field.api}
      />
    );
  } else if (field.compoType === 'cascader') {
    //options是个数组
    htmls = (
      <Cascader
        options={listData}
        onChange={field.onChange}
        fieldNames={field.fieldNames}
        changeOnSelect={field.parentSelect}
      />
    );
  } else if (field.compoType == 'menuTree') {
    //目录选择
    let defaultCheckedKeys =
      initData[field.name] !== undefined ? [...initData[field.name]] : [];
    let expandedKeys = [];
    let checkedKeys = [];
    //遍历tree提取数据
    const menuTreeFunc = (data) => {
      data.map((val) => {
        checkedKeys.push(val['key']);
        if (val['children'] !== undefined && val['children'].length > 0) {
          delArrVal(defaultCheckedKeys, val['key']);
          expandedKeys.push(val['key']);
          menuTreeFunc(val['children']);
        }
      });
    };
    menuTreeFunc(listData);
    const [menuTreeExpandedKeys, setMenuTreeExpandedKeys] = useState([]);
    const [menuTreeCheckedKeys, setMenuTreeCheckedKeys] =
      useState(defaultCheckedKeys);
    //check form表单数据绑定
    const menuTreeCheckFunc = (checkVal, parentCheckVal) => {
      const vals = {};
      setMenuTreeCheckedKeys(checkVal);
      vals[field.name] = [...checkVal, ...parentCheckVal];
      formRef.setFieldsValue(vals);
    };
    htmls = (
      <>
        <Space className="menu-tree-top">
          <Checkbox
            onChange={(event) => {
              if (event.target.checked) {
                setMenuTreeExpandedKeys(expandedKeys);
              } else {
                setMenuTreeExpandedKeys([]);
              }
            }}
          >
            展开/折叠
          </Checkbox>
          <Checkbox
            onChange={(event) => {
              if (event.target.checked) {
                menuTreeCheckFunc(checkedKeys, []);
              } else {
                menuTreeCheckFunc([], []);
              }
            }}
          >
            全选/全不选
          </Checkbox>
        </Space>
        <Tree
          checkable
          showLine={{ showLeafIcon: false }}
          showIcon={true}
          selectable={false}
          expandedKeys={menuTreeExpandedKeys}
          onExpand={(expandVal) => {
            setMenuTreeExpandedKeys(expandVal);
          }}
          onCheck={(checkVal, { halfCheckedKeys }) => {
            menuTreeCheckFunc(checkVal, halfCheckedKeys);
          }}
          checkedKeys={menuTreeCheckedKeys}
          treeData={listData}
        />
      </>
    );
  }else if(field.compoType == 'transfer'){
    //目录选择
    let defaultTargetKeys = initData[field.name] !== undefined ? [...initData[field.name]] : [];
    const [targetKeys, setTargetKeys] = useState(defaultTargetKeys);
    const [selectedKeys, setSelectedKeys] = useState([]);
    const onChange = (nextTargetKeys, direction, moveKeys) => {
      // console.log('targetKeys:', nextTargetKeys);
      // console.log('direction:', direction);
      // console.log('moveKeys:', moveKeys);
      setTargetKeys(nextTargetKeys);
    };

    const onSelectChange = (sourceSelectedKeys, targetSelectedKeys) => {
      // console.log('sourceSelectedKeys:', sourceSelectedKeys);
      // console.log('targetSelectedKeys:', targetSelectedKeys);
      setSelectedKeys([...sourceSelectedKeys, ...targetSelectedKeys]);
    };
    htmls = (<Transfer
      dataSource={listData}
      // showSearch
      titles={paramIsset(field.topTitle,['资源项', '目标项'])}
      targetKeys={targetKeys}
      selectedKeys={selectedKeys}
      onChange={onChange}
      onSelectChange={onSelectChange}
      render={item => item.title}
    />)

  }

  if (field.compoType !== 'br') {
    htmls = (
      <Form.Item
        // labelCol={{ span: 6 }}
        // wrapperCol={{ span: 18 }}
        label={field.label}
        name={field.name}
        required={field.hidden === true ? false : field.required}
        tooltip={field.tooltip}
        extra={field.extra}
        className={className}
        hidden={field.hidden}
        labelAlign="right"
        rules={rules}
        initialValue={formFieldTrans(field)}
      >
        {htmls}
      </Form.Item>
    );
  }

  return (
    <Col
      span={field.col === undefined ? 12 : Number(field.col)}
      offset={Number(field.offset)}
      key={index}
    >
      {htmls}
    </Col>
  );
};
export default ItemField;
