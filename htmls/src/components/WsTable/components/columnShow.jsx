import React, { useState, useMemo } from 'react';
import { Dropdown, Menu, Tree, Checkbox } from 'antd';

const ColumnShow = (props) => {
  const { solt, data, showColumns, setShowColumns } = props;
  const [indeterminate, setIndeterminate] = useState(true);
  const [checkAll, setCheckAll] = useState(false);
  const { allData, allKeys } = data;

  const onCheck = (checkedKeys, info) => {
    setShowColumns(checkedKeys);
    setIndeterminate(!!checkedKeys.length && checkedKeys.length < allKeys.length);
    setCheckAll(checkedKeys.length === allKeys.length);
  };
  const onDrop = (info) => {
    console.log('info', info);
  }
  const onCheckAllChange = (e) => {
    setIndeterminate(false);
    setCheckAll(e.target.checked)
    if (e.target.checked) {
      setShowColumns(allKeys);
    } else {
      setShowColumns([]);
    }
  }
  const menu = (
    <Menu
      className="table-column-show"
    >
      <Checkbox
        style={{ padding: "5px 10px" }}
        indeterminate={indeterminate}
        onChange={onCheckAllChange}
        checked={checkAll}>
        列展示
      </Checkbox>
      <Menu.Divider />
      <Tree
        checkable
        checkedKeys={showColumns}
        // draggable
        selectable={false}
        // height={150}
        onCheck={onCheck}
        onDrop={onDrop}
        treeData={allData}
      />
    </Menu>
  );
  return (
    <>
      <Dropdown
        overlay={menu}
        trigger={['click']}
      >
        {solt}
      </Dropdown>
    </>
  );
};

export default ColumnShow;
