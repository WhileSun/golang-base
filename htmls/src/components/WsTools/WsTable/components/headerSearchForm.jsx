import { paramIsset, funcIsset, momentDate } from '../utils/tools';
import { Form, Input, Select, DatePicker, Button, Divider } from 'antd';
const { RangePicker } = DatePicker;
const { Option } = Select;


const HeaderSearchForm = (props) => {
  console.log("initSearchForm");
  const searchs = paramIsset(props.searchs, []);
  const formRef = props.formRef;
  const handleFormSubmit = paramIsset(props.handleFormSubmit, () => { })

  return (
    <>
      <Form
        layout="inline"
        form={formRef}
        initialValues={{}}
        onFinish={handleFormSubmit}
      // onValuesChange={}
      >
        {searchs.map((field, index) => initFields(field, index))}
      </Form>
    </>
  );
};

const initFields = (field, index) => {
  let htmls = '';
  let styles = {};
  let label = '';
  let fieldName = field.name;
  if (field.width !== undefined) {
    styles.width = field.width + 'px';
  }
  if (field.oneShowMode === true) {
    label = field.label;
    delete field.label;
  }
  if (field.type === 'input') {
    htmls = (
      <Input
        placeholder={field.placeholder}
        style={{ width: '150px', ...styles }}
      />
    );
  } else if (field.type === 'select') {
    htmls = (
      <Select
        placeholder={field.placeholder}
        style={{ width: '100px', ...styles }}
      >
        <Option value="" key="">
          {field.oneShowMode === true ? label : '全部'}
        </Option>
        {Object.keys(field['listData']).map((selectKey, selectIndex) => {
          return (
            <Option value={selectKey} key={selectIndex.toString()}>
              {field['listData'][selectKey]}
            </Option>
          );
        })}
      </Select>
    );
  } else if (field.type === 'selectInput') {
    let keyName = 'selectInput_' + index;
    let keys = Object.keys(field['listData']);
    htmls = (
      <Input.Group compact>
        <Form.Item name={[keyName, 'key']} noStyle initialValue={keys[0]}>
          <Select>
            {keys.map((selectKey, selectIndex) => {
              return (
                <Option value={selectKey} key={selectIndex.toString()}>
                  {field['listData'][selectKey]}
                </Option>
              );
            })}
          </Select>
        </Form.Item>
        <Form.Item name={[keyName, 'val']} noStyle>
          <Input style={{ width: '150px' }} placeholder="搜索" />
        </Form.Item>
      </Input.Group>
    );
  } else if (field.type === 'dateRange') {
    fieldName = 'dateRange_' + field.name;
    htmls = <RangePicker style={{ width: '180px' }} />;
  }

  return (
    <Form.Item
      label={field.label}
      name={fieldName}
      initialValue={field.defaultValue === undefined ? '' : field.defaultValue}
      style={{ marginBottom: '8px', marginRight: '10px' }}
      key={index}
    >
      {htmls}
    </Form.Item>
  );
};

export default HeaderSearchForm;
