import {iconfontUrl} from '@/config'
import { createFromIconfontCN } from '@ant-design/icons';

const Iconfont = createFromIconfontCN({
    scriptUrl: iconfontUrl,
});

const WsIcon = (props) => {
    return (
        <>
            <Iconfont
              type={props.type}
              style={props.style}
            />
        </>
    );
}

export default WsIcon;