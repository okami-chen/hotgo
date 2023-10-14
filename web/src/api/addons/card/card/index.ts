import { http, jumpExport } from '@/utils/http/axios';

// 获取卡片列表
export function List(params) {
  return http.request({
    url: '/card/card/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除卡片
export function Delete(params) {
  return http.request({
    url: '/card/card/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑卡片
export function Edit(params) {
  return http.request({
    url: '/card/card/edit',
    method: 'POST',
    params,
  });
}




// 获取卡片指定详情
export function View(params) {
  return http.request({
    url: '/card/card/view',
    method: 'GET',
    params,
  });
}



// 导出卡片
export function Export(params) {
  jumpExport('/card/card/export', params);
}