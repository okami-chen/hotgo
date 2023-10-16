import { http, jumpExport } from '@/utils/http/axios';

// 获取货币列表
export function List(params) {
  return http.request({
    url: '/system/currency/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除货币
export function Delete(params) {
  return http.request({
    url: '/system/currency/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑货币
export function Edit(params) {
  return http.request({
    url: '/system/currency/edit',
    method: 'POST',
    params,
  });
}


// 修改货币状态
export function Status(params) {
  return http.request({
    url: '/system/currency/status',
    method: 'POST',
    params,
  });
}



// 获取货币指定详情
export function View(params) {
  return http.request({
    url: '/system/currency/view',
    method: 'GET',
    params,
  });
}



// 导出货币
export function Export(params) {
  jumpExport('/system/currency/export', params);
}