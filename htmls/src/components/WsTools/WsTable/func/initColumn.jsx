import { Input, Space, Button } from 'antd';

const filterIndex = (record, dataIndex, value) => {
  let res = record[dataIndex]
    .toString()
    .toLowerCase()
    .includes(value.toLowerCase());
  if (!res && record.children) {
    if (record.children.some(item => item[dataIndex].toString()
      .toLowerCase()
      .includes(value.toLowerCase()))) {
      res = true;
    } else {
      for (let i = 0; i < record.children.length; i++) {
        res = filterIndex(record.children[i], dataIndex, value);
        if (res) {
          break;
        }
      }
    }
  }
  return res;
}
/**
 * 初始化column
 * @param {Array} fields table的字段
 * @param {Array} showColumns table展示的字段
 * @returns 
 */
const initColumnFunc = (fields, showColumns) => {
  console.log('initColumn');
  var columns = [];
  var sumWidth = 0;
  var tableScroll = {};
  for (var i = 0, len = fields.length; i < len; i++) {
    let column = {};
    var d = fields[i];
    let showColumnKey = d['name'] + '_' + i;
    column['ellipsis'] = true;
    // column['className'] = 'cell-row ';
    for (let dkey in d) {
      switch (dkey) {
        case 'title':
          column['title'] = d['title'];
          break;
        case 'name':
          column['dataIndex'] = d['name'];
          break;
        case 'width':
          // px-100,%-20%
          if (showColumns.includes(showColumnKey)) {
            sumWidth += d['width'];
          }
          column['width'] = d['width'];
          break;
        case 'align':
          column['align'] = d['align'];
          break;
        case 'className':
          column['className'] = d['className'];
          break;
        case 'editable':
          //是否可编辑 默认false
          column['editable'] = d['editable'];
          break;
        case 'ellipsis':
          //超过宽度将自动省略 默认false
          column['ellipsis'] = d['ellipsis'];
          break;
        case 'fixed':
          //可选 true (等效于 left) left right
          column['fixed'] = d['fixed'];
          break;
        case 'render':
          //生成复杂数据的渲染函数，参数分别为当前行的值，当前行数据，行索引
          //function(text, record, index) {}
          column['render'] = d['render'];
          break;
        case 'sorter':
          column['sorter'] = d['sorter'];
          break;
      }
    }
    //暂时不用，后续使用在实现
    if (d.filter) {
      column['onFilter'] = (value, record) => {
        let dataIndex = column['dataIndex'];
        return filterIndex(record, dataIndex, value)
      }
      column['filterDropdown'] = ({ setSelectedKeys, selectedKeys, confirm, clearFilters }) => {
        const handleSearch = (selectedKeys, confirm) => {
          confirm();
        };
        const handleReset = clearFilters => {
          clearFilters();
        };
        console.log(selectedKeys);
        return (
          <div style={{ padding: 8 }}>
            <Input
              value={selectedKeys[0]}
              onChange={e => setSelectedKeys(e.target.value ? [e.target.value] : [])}
              onPressEnter={() => handleSearch(selectedKeys, confirm)}
              style={{ marginBottom: 8, display: 'block' }}
            />
            <Space>
              <Button
                type="primary"
                onClick={() => handleSearch(selectedKeys, confirm)}
                // icon={<SearchOutlined />}
                size="small"
                style={{ width: 90 }}
              >
                Search
              </Button>
              <Button onClick={() => handleReset(clearFilters)} size="small" style={{ width: 90 }}>
                Reset
              </Button>
            </Space>
          </div>
        )
      }
    }
    if (showColumns.includes(showColumnKey)) {
      columns.push(column);
    }
  }
  //防止撑大
  columns.push({});
  sumWidth += 0;
  //设置table宽高
  tableScroll = { x: sumWidth, y: 5000 };
  return { 'columns': columns, 'tableScroll': tableScroll };
};

export default initColumnFunc;