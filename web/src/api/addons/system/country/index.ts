import { http, jumpExport } from '@/utils/http/axios';

// 获取国家列表
export function List(params) {
  return http.request({
    url: '/system/country/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除国家
export function Delete(params) {
  return http.request({
    url: '/system/country/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑国家
export function Edit(params) {
  return http.request({
    url: '/system/country/edit',
    method: 'POST',
    params,
  });
}




// 获取国家指定详情
export function View(params) {
  return http.request({
    url: '/system/country/view',
    method: 'GET',
    params,
  });
}



// 导出国家
export function Export(params) {
  jumpExport('/system/country/export', params);
}