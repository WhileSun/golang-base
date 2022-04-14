import requests from '@/request/index';

/**获取验证码 */
export async function getCaptcha() {
  return requests.get('sys/loginCaptcha/get');
}

/**用户登录 */
export async function loginUser(data) {
  return requests.post('user/login',data);
}

/**用户退出 */
export async function outLoginUser() {
  await requests.post('user/outLogin');
  return requests.logout();
}

/**获取用户信息 */
export async function getUserInfo() {
  return requests.get('user/info/get');
}

