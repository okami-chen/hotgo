import { http, jumpExport } from '@/utils/http/axios';

// 获取小区列表
export function List(params) {
  params.pageSize = 30;
  return http.request({
    url: '/fang/village/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除小区
export function Delete(params) {
  return http.request({
    url: '/fang/village/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑小区
export function Edit(params) {
  return http.request({
    url: '/fang/village/edit',
    method: 'POST',
    params,
  });
}




// 获取小区指定详情
export function View(params) {
  return http.request({
    url: '/fang/village/view',
    method: 'GET',
    params,
  });
}



// 导出小区
export function Export(params) {
  jumpExport('/fang/village/export', params);
}
