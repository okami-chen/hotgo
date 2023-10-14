import { http, jumpExport } from '@/utils/http/axios';

// 获取事件列表
export function List(params) {
  return http.request({
    url: '/system/event/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除事件
export function Delete(params) {
  return http.request({
    url: '/system/event/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑事件
export function Edit(params) {
  return http.request({
    url: '/system/event/edit',
    method: 'POST',
    params,
  });
}


// 修改事件状态
export function Status(params) {
  return http.request({
    url: '/system/event/status',
    method: 'POST',
    params,
  });
}



// 获取事件指定详情
export function View(params) {
  return http.request({
    url: '/system/event/view',
    method: 'GET',
    params,
  });
}



// 导出事件
export function Export(params) {
  jumpExport('/system/event/export', params);
}