import { paramIsset} from '../utils/tools';
import WsButton from '../../WsButton';

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
		<WsButton
		type='primary'
		size="middle"
		key={index}
		style={{ marginRight: '10px' }}
		onClick={btn.onClick}
		title = {btn.title}
		color = {btn.color}
		/>
	);
};

export default HeaderButtonLeft;