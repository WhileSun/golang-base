import React, { useState} from 'react';
import {useModel} from 'umi';
import {WsForm} from '@/components/WsTools';
import { message } from 'antd';
import {userOutLogin} from '@/services/user';
import { loginPath } from '@/config';

export default (props) => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const { currentUser } = initialState;
  const formRef = WsForm.useForm();
  return (
    <>
    <WsForm
        form={formRef}
        widthCol={10}
        mode="Form"
        title="修改密码"
        data = {{realname:currentUser?.name}}
        fields={[
          {name:"realname",col:24,label:'用户名称',compoType:'text',disabled:true,required:true},
          {name:"old_passwd",col:24,label:'旧密码',compoType:'inputPasswd',required:true},
          {name:"new_passwd",col:24,label:'新密码',compoType:'inputPasswd',required:true},
          {name:"repeat_passwd",col:24,label:'确认新密码',compoType:'inputPasswd',required:true},
        ]}
        idKey="realname"
        updateApi = "user/passwd/update"
        onBeforeSubmit={(params, cb) => {
          var passwd_patt = /(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[#@!~%^&*.])[a-zA-Z\d#@!~%^&*.]{6,}/;
          if(!passwd_patt.test(params.new_passwd)){
            message.error("密码需要符合规则[大小写字母+数字+特殊符号+6位以上]");
            return;
          }
          if(params.new_passwd != params.repeat_passwd){
            message.error("新密码二次输入不相同，请确认！");
            return;
          }
          cb();
        }}
        onSucc={async ()=>{
          await userOutLogin();
          history.push(loginPath);
        }}
      />
    </> 
  );
};