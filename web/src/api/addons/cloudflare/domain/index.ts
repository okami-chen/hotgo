import { http, jumpExport } from '@/utils/http/axios';

// 获取名列表
export function List(params) {
  return http.request({
    url: '/cloudflare/domain/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除名
export function Delete(params) {
  return http.request({
    url: '/cloudflare/domain/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑名
export function Edit(params) {
  return http.request({
    url: '/cloudflare/domain/edit',
    method: 'POST',
    params,
  });
}




// 获取名指定详情
export function View(params) {
  return http.request({
    url: '/cloudflare/domain/view',
    method: 'GET',
    params,
  });
}



// 导出名
export function Export(params) {
  jumpExport('/cloudflare/domain/export', params);
}