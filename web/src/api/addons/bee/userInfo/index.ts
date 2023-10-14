import { http, jumpExport } from '@/utils/http/axios';

// 获取客户信息扩展列表
export function List(params) {
  return http.request({
    url: '/bee/userInfo/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除客户信息扩展
export function Delete(params) {
  return http.request({
    url: '/bee/userInfo/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑客户信息扩展
export function Edit(params) {
  return http.request({
    url: '/bee/userInfo/edit',
    method: 'POST',
    params,
  });
}




// 获取客户信息扩展指定详情
export function View(params) {
  return http.request({
    url: '/bee/userInfo/view',
    method: 'GET',
    params,
  });
}



// 导出客户信息扩展
export function Export(params) {
  jumpExport('/bee/userInfo/export', params);
}