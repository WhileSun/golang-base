import { Button, Popconfirm } from 'antd';
import React, { useMemo } from 'react';
import { titleColorType, buttonStyle} from './func/initColor';
import { paramIsset } from '../utils/tools';
import './index.less';

const WsButton = (props) => {
  //基础配置
  const config = useMemo(() => {
    let param = {};
    param.title = paramIsset(props.title, '按钮'); //按钮名称
    param.onClick = paramIsset(props.onClick, () => { }); //点击事件
    param.block = paramIsset(props.block, false); //将按钮宽度调整为其父宽度的选项
    param.disabled = paramIsset(props.disabled, false); //按钮失效状态
    param.ghost = paramIsset(props.ghost, false); //幽灵属性，使按钮背景透明
    param.icon = paramIsset(props.icon, ''); //设置按钮的图标组件
    param.loading = paramIsset(props.loading, false); //设置按钮载入状态
    param.size = paramIsset(props.size, 'small'); //设置按钮载入状态
    param.type = paramIsset(props.type, 'default');
    param.colorType = paramIsset(props.color, titleColorType(props.title));
    param.pop = paramIsset(props.pop,false);
    param.popTitle = paramIsset(props.popTitle,'您确定删除选中的数据吗?');
    param.style = paramIsset(props.style,{}); //自定义样式
    return param;
  }, [props])

  let html = (
    <Button
      type={config.type}
      size={config.size}
      style={config.style}
      icon={config.icon}
      disabled={config.disabled}
      block={config.block}
      loading={config.loading}
      className={'ws-btn-'+config.colorType}
      onClick={(event) => {
        if (config.pop !== true) {
          config.onClick.call(this,event)
        }
      }}
    >
      {config.title}
    </Button>
  );

  if (config.pop === true) {
    html = (
      <Popconfirm
        title={config.popTitle}
        onConfirm={() => { config.onClick.call(this) }}
        // onCancel={cancel}
        okText="确定"
        cancelText="取消"
      >
        {html}
      </Popconfirm>
    )
  }

  return (
    <>
      {html}
    </>
  );
};
export default WsButton;