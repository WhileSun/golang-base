import requests from '@/request/index';
import { request } from 'umi';

/**获取验证码 */
export async function getCaptcha() {
  return requests.get('sys/captcha/get');
}

/**角色管理 */
export async function getRoleList(params) {
  return requests.post('role/list',params);
}

export async function addRole(params){
  return requests.post('role/add',params);
}

export async function updateRole(params){
  return requests.post('role/update',params);
}

/**菜单管理 */
export async function getMenuList(params){
  return requests.post('menu/list',params);
}

export async function addMenu(params){
  return requests.post('menu/add',params);
}

export async function updateMenu(params){
  return requests.post('menu/update',params);
}

export async function deleteMenu(params){
  return requests.post('menu/delete',params);
}

/**节点管理 */
export async function getPermsList(params){
  return requests.post('perms/list',params);
}

export async function addPerms(params){
  return requests.post('perms/add',params);
}

export async function updatePerms(params){
  return requests.post('perms/update',params);
}

export async function deletePerms(params){
  return requests.post('perms/delete',params);
}
