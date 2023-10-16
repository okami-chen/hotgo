import { http, jumpExport } from '@/utils/http/axios';

// 获取语言列表
export function List(params) {
  return http.request({
    url: '/system/language/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除语言
export function Delete(params) {
  return http.request({
    url: '/system/language/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑语言
export function Edit(params) {
  return http.request({
    url: '/system/language/edit',
    method: 'POST',
    params,
  });
}




// 获取语言指定详情
export function View(params) {
  return http.request({
    url: '/system/language/view',
    method: 'GET',
    params,
  });
}



// 导出语言
export function Export(params) {
  jumpExport('/system/language/export', params);
}