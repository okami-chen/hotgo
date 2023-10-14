import { http, jumpExport } from '@/utils/http/axios';

// 获取完成时间列表
export function List(params) {
  return http.request({
    url: '/order/log/complete/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除完成时间
export function Delete(params) {
  return http.request({
    url: '/order/log/complete/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑完成时间
export function Edit(params) {
  return http.request({
    url: '/order/log/complete/edit',
    method: 'POST',
    params,
  });
}




// 获取完成时间指定详情
export function View(params) {
  return http.request({
    url: '/order/log/complete/view',
    method: 'GET',
    params,
  });
}



// 导出完成时间
export function Export(params) {
  jumpExport('/order/log/complete/export', params);
}
