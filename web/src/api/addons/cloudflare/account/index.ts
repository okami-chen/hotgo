import { http, jumpExport } from '@/utils/http/axios';

// 获取账号列表
export function List(params) {
  return http.request({
    url: '/cloudflare/account/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除账号
export function Delete(params) {
  return http.request({
    url: '/cloudflare/account/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑账号
export function Edit(params) {
  return http.request({
    url: '/cloudflare/account/edit',
    method: 'POST',
    params,
  });
}




// 获取账号指定详情
export function View(params) {
  return http.request({
    url: '/cloudflare/account/view',
    method: 'GET',
    params,
  });
}



// 导出账号
export function Export(params) {
  jumpExport('/cloudflare/account/export', params);
}