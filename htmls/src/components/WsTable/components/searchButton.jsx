import { paramIsset} from '../utils/tools';
import { Button} from 'antd';

const SearchButton = (props) => {
	console.log("initSearchButton");
	const searchloading = paramIsset(props.loading, false);
	const fieldLen = Object.keys(paramIsset(props.searchs, [])).length;
	const btns = paramIsset(props.btns, []);
	const formRef = props.formRef;
	const soltRightBtn = paramIsset(props.soltRightBtn,"");
	const handleFormReset = paramIsset(props.handleFormReset,()=>{})

	return (
		<>
			<div style={{ display: 'flex', justifyContent: 'space-between' }}>
				<span>
					{btns.map((btn, index) => initBtns(btn, index))}
					{fieldLen > 0 ? (
						<>
							<Button htmlType="button" onClick={handleFormReset} style={{ marginRight: '10px' }} loading={searchloading}>
								重置
							</Button>
							<Button type="primary" onClick={() => { formRef.submit(); }} loading={searchloading}>
								查询
							</Button>
						</>
					) : (
						''
					)}
				</span>
				<span>
					{soltRightBtn}
				</span>
			</div>
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


export default SearchButton;