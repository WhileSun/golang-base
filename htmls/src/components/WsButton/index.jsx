import { Button,Popconfirm } from 'antd';

const titleColor = (title, buttonColor) => {
  let newButtonColor = '';
  switch (title) {
    case '编辑':
      newButtonColor = 'bule';
      break;
    case '删除':
      newButtonColor = 'red';
      break;
    case '添加':
      newButtonColor = 'green';
      break;
    default:
      newButtonColor = buttonColor;
  }
  return newButtonColor;
};

//按钮颜色调节
const buttonColor = (buttonColor) => {
  let styleColor = {};
  switch (buttonColor) {
    case 'bule':
      styleColor = { borderColor: '#1890ff', color: '#1890ff' };
      break;
    case 'red':
      styleColor = { borderColor: '#FF5722', color: '#FF5722' };
      break;
    case 'green':
      styleColor = { borderColor: '#66cc00', color: '#66cc00' };
      break;
  }
  return styleColor;
};

const WsButton = (props) => {
  const colorName = titleColor(props.title, props.colorName);
  const clickFunc = props.onClick===undefined ? ()=>{}:props.onClick; 

  let html = (
  <Button
    type="primary"
    ghost
    size="small"
    style={buttonColor(colorName)}
    onClick={(event) => {
      if(props.pop!==true){
        clickFunc();
      }
    }}
  >
    {props.title}
  </Button>
  );

  if(props.pop===true){
    html = (
    <Popconfirm
        title={props.popTitle===undefined?"您确定删除数据吗?":props.popTitle}
        onConfirm={()=>{clickFunc()}}
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
