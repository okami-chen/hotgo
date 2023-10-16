import { http, jumpExport } from '@/utils/http/axios';

// 获取证书列表
export function List(params) {
  return http.request({
    url: '/cloudflare/ssl/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除证书
export function Delete(params) {
  return http.request({
    url: '/cloudflare/ssl/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑证书
export function Edit(params) {
  return http.request({
    url: '/cloudflare/ssl/edit',
    method: 'POST',
    params,
  });
}


// 修改证书状态
export function Status(params) {
  return http.request({
    url: '/cloudflare/ssl/status',
    method: 'POST',
    params,
  });
}



// 获取证书指定详情
export function View(params) {
  return http.request({
    url: '/cloudflare/ssl/view',
    method: 'GET',
    params,
  });
}



// 导出证书
export function Export(params) {
  jumpExport('/cloudflare/ssl/export', params);
}