import { http, jumpExport } from '@/utils/http/axios';

// 获取用户信息列表
export function List(params) {
  return http.request({
    url: '/bee/user/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除用户信息
export function Delete(params) {
  return http.request({
    url: '/bee/user/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑用户信息
export function Edit(params) {
  return http.request({
    url: '/bee/user/edit',
    method: 'POST',
    params,
  });
}


// 修改用户信息状态
export function Status(params) {
  return http.request({
    url: '/bee/user/status',
    method: 'POST',
    params,
  });
}



// 获取用户信息指定详情
export function View(params) {
  return http.request({
    url: '/bee/user/view',
    method: 'GET',
    params,
  });
}



// 导出用户信息
export function Export(params) {
  jumpExport('/bee/user/export', params);
}