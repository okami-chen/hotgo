import { http, jumpExport } from '@/utils/http/axios';

// 获取站群列表
export function List(params) {
  return http.request({
    url: '/bee/systemSite/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除站群
export function Delete(params) {
  return http.request({
    url: '/bee/systemSite/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑站群
export function Edit(params) {
  return http.request({
    url: '/bee/systemSite/edit',
    method: 'POST',
    params,
  });
}




// 获取站群指定详情
export function View(params) {
  return http.request({
    url: '/bee/systemSite/view',
    method: 'GET',
    params,
  });
}


// 获取站群最大排序
export function MaxSort() {
  return http.request({
    url: '/bee/systemSite/maxSort',
    method: 'GET',
  });
}


// 导出站群
export function Export(params) {
  jumpExport('/bee/systemSite/export', params);
}