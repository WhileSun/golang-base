import { LockOutlined, SafetyOutlined, UserOutlined } from '@ant-design/icons';
import { message, Divider } from 'antd';
import React, { useState, useEffect, useRef } from 'react';
import { history, useModel } from 'umi';
import Footer from '@/components/Footer';
import { getCaptcha, userLogin } from '@/services/api';
import { setToken, getToken} from '@/token';
import styles from './index.less';
import { Form, Input, Button, Checkbox } from 'antd';

const Login = () => {
  const { initialState, setInitialState } = useModel('@@initialState');
  const [loadingState, setLoadingState] = useState(false);
  const [captcha, setCaptcha] = useState({});
  const [form] = Form.useForm();

  const getCaptchaInfo = async () => {
    const resp = await getCaptcha();
    if (resp.code == 0) {
      setCaptcha(resp.data);
      form.resetFields(['captcha'], []);
    }
  };
  
  //初始化获取验证码
  useEffect(() => {
    if(getToken() != ""){
      history.push('/');
    }
    getCaptchaInfo();
  }, []);

  const fetchLoginApis = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();
    if (userInfo) {
      await setInitialState((s) => ({ ...s, currentUser: userInfo }));
    }
    const menuData = await initialState?.fetchMenuData?.(userInfo?.name);
    if (menuData) {
      await setInitialState((s) => ({ ...s, menuData: menuData }));
    }
  };

  const handleSubmit = async (values) => {
    try {
      setLoadingState(true);
      // 登录
      const resp = await userLogin({ ...values, captcha_id: captcha.id });
      setLoadingState(false);
      if (resp.code === 0) {
        message.success('登录成功！');
        setToken(resp.data.token);
        await fetchLoginApis();
        /** 此方法会跳转到 redirect 参数所在的位置 */
        if (!history) return;
        const { query } = history.location;
        const { redirect } = query;
        history.push(redirect || '/');
        return;
      }
      message.error(resp.msg);
      getCaptchaInfo();
    } catch (error) {
      setLoadingState(false);
      message.error('登录失败,Error:' + error);
    }
  };

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles.loginForm}>
          <p className={styles.loginTitle}>
            <span>登 录</span>
          </p>
          <Divider />
          <Form
            name="basic"
            initialValues={{ remember: true }}
            onFinish={handleSubmit}
            autoComplete="off"
            form={form}
            size="large"
          >
            <Form.Item
              name="username"
              rules={[{ required: true, message: '请输入用户名！' }]}
            >
              <Input
                prefix={<UserOutlined className={styles.prefixIcon} />}
                placeholder="用户名"
              />
            </Form.Item>
            <Form.Item
              name="password"
              rules={[{ required: true, message: '请输入密码！' }]}
            >
              <Input.Password
                prefix={<LockOutlined className={styles.prefixIcon} />}
                type="password"
                placeholder="密码"
              />
            </Form.Item>
            <Form.Item style={{ marginBottom: 0 }}>
              <Form.Item
                name="captcha_code"
                // noStyle
                className={styles.captchaItemLeft}
                rules={[
                  {
                    required: true,
                    message: '请输入验证码！',
                  },
                ]}
              >
                <Input
                  prefix={<SafetyOutlined className={styles.prefixIcon} />}
                  type="captcha"
                  placeholder="验证码"
                />
              </Form.Item>
              <Form.Item
                className={styles.captchaItemRight}
              >
                <img
                  alt="验证码"
                  src={captcha.src}
                  className={styles.captchaRight}
                  onClick={getCaptchaInfo}
                />
              </Form.Item>
            </Form.Item>

            <Form.Item name="remember" valuePropName="checked">
              <Checkbox>自动登录</Checkbox>
            </Form.Item>
            <Form.Item>
              <Button
                type="primary"
                htmlType="submit"
                className={styles.loginButton}
                loading={loadingState}
              >
                登录
              </Button>
            </Form.Item>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default Login;
