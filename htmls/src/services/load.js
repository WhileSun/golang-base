import requests from '@/request/index';

//获取角色基础字段
export async function getLoadRoleList(){
    return requests.post('load/role/list/get');
}

export async function getLoadMenuList(params){
    return requests.post('load/menu/list/get');
}

export async function getLoadPermsList(){
    return requests.post('load/perms/list/get');
}

export async function getLoadWorkProjectList(){
    return requests.post('load/work_project/list/get');
}