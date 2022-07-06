import requests from '@/request/index';

/**项目管理 */
export async function getWorkProjectList(params) {
  return requests.post('work_project/list/get', params);
}

export async function addWorkProject(params) {
  return requests.post('work_project/add', params);
}

export async function updateWorkProject(params) {
  return requests.post('work_project/update', params);
}

/**任务管理 */
export async function getWorkTaskList(params) {
  return requests.post('work_task/list/get', params);
}

export async function addWorkTask(params) {
  return requests.post('work_task/add', params);
}

export async function updateWorkTask(params) {
  return requests.post('work_task/update', params);
}

export async function deleteWorkTask(params){
    return requests.post('work_task/delete', params);
}