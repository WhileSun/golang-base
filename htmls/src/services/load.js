import requests from '@/request/index';

//获取角色基础字段
export async function getRoleNameList(){
    return requests.post('role/name/list');
}

export async function getMenuNameList(params){
    return requests.post('menu/name/list',params);
}

export async function getPermsNameList(){
    return requests.post('perms/name/list');
}

export async function getLoadWorkProjectList(){
    return requests.post('load/work_project/list/get');
}