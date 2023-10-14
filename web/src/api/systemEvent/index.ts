import { http, jumpExport } from '@/utils/http/axios';

// 获取系统事件列表
export function List(params) {
  return http.request({
    url: '/systemEvent/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除系统事件
export function Delete(params) {
  return http.request({
    url: '/systemEvent/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑系统事件
export function Edit(params) {
  return http.request({
    url: '/systemEvent/edit',
    method: 'POST',
    params,
  });
}


// 修改系统事件状态
export function Status(params) {
  return http.request({
    url: '/systemEvent/status',
    method: 'POST',
    params,
  });
}



// 获取系统事件指定详情
export function View(params) {
  return http.request({
    url: '/systemEvent/view',
    method: 'GET',
    params,
  });
}



// 导出系统事件
export function Export(params) {
  jumpExport('/systemEvent/export', params);
}