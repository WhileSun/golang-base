import requests from '@/request/index';

/**用户登录 */
export async function userLogin(data) {
  return requests.post('user/login',data);
}

/**用户退出 */
export async function userOutLogin() {
  await requests.post('user/out.login');
  return requests.logout();
}

/**获取用户信息 */
export async function getUserInfo() {
  return requests.get('user/info/get');
}

/**获取用户路由列表 */
export async function getUserRouteList(){
  return requests.get('user/route/list/get');
}

export async function getUserList(params){
  return requests.post('user/list/get',params);
}

export async function addUser(params){
  return requests.post('user/add',params);
}

export async function updateUser(params){
  return requests.post('user/update',params);
}