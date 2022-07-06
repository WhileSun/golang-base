import { paramIsset} from '../utils/tools';
import { Button} from 'antd';

const HeaderButtonLeft = (props) => {
	const btns = paramIsset(props.btns, []);
	return (
		<>
			{btns.map((btn, index) => initBtns(btn, index))}
		</>
	);
}

//生成button
const initBtns = (btn, index) => {
	return (
		<Button
			htmlType="button"
			onClick={btn.callback}
			key={index}
			style={{ marginRight: '10px' }}
		>
			{btn.text}
		</Button>
	);
};

export default HeaderButtonLeft;