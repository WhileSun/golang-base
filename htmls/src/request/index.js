import { extend  } from 'umi-request';
import { history } from 'umi';
import {getToken,deleteToken} from "@/token"
import {loginPath,AUTH_FAIL} from '@/config'
import { message } from 'antd';
const request = extend({
    prefix: '/api/',
    errorConfig: {},
    middlewares: [],
    // errorHandler,
});

request.interceptors.request.use(async (url, options) => {
    if (getToken()!=""){
        options['headers']['Authorization'] =  "Bearer "+getToken();
    }
    return ({
          url: url,
          options: { ...options},
    });
})

request.interceptors.response.use(async response => {
    const resp = await response.clone().json();
    if(resp.code == AUTH_FAIL){
        //失效先删除
        deleteToken();
        history.push(loginPath);
    }
    return response;
});

const requests = {
    postFile(api,data){
        return request(api,{
            method:'POST',
            data:data,
        })
    },
    post(api,data,options,timeout=5000){
        return request(api,{
            method: 'POST',
            requestType:'form',
            timeout:timeout,
            data:{...data},
            ...(options || {}),
          }
        );
    },
    get(api,param,options,timeout=5000){
        return request(api, {
            method: 'GET',
            headers: { },
            timeout:timeout,
            param:{...param},
            ...(options || {}),
            }
        );
    },
    logout(){
        deleteToken();
    }
}
export default requests