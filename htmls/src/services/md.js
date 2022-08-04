import requests from '@/request/index';

/**文档管理 */
export async function getMdDocumentNameList(params){
    return requests.post('md_document_name/list/get', params);
}

export async function addMdDocumentName(params) {
  return requests.post('md_document_name/add', params);
}

export async function updateMdDocumentName(params){
    return requests.post('md_document_name/update', params);
}

export async function deleteMdDocumentName(params){
  return requests.post('md_document_name/delete',params)
}

export async function dragMdDocumentName(params){
  return requests.post('md_document_name/drag',params)
}

export async function uploadDocumentNameFile(params){
  return requests.postFile('md_document_name/upload_file',params)
}

export async function getMdDocumentText(params){
  return requests.post('md_document_text/get',params)
}

export async function updateMdDocumentText(params){
  return requests.post('md_document_text/update',params)
}
/**书籍管理 */
export async function getMdBookList(params){
  return requests.post('md_book/list/get',params)
}

export async function addMdBook(params){
  return requests.post('md_book/add',params)
}

export async function updateMdBook(params){
  return requests.post('md_book/update',params)
}
