import { http, jumpExport } from '@/utils/http/axios';

// 获取租房列表
export function List(params) {
  params.pageSize = 15
  return http.request({
    url: '/fang/fang/list',
    method: 'get',
    params,
  });
}

// 删除/批量删除租房
export function Delete(params) {
  return http.request({
    url: '/fang/fang/delete',
    method: 'POST',
    params,
  });
}


// 添加/编辑租房
export function Edit(params) {
  return http.request({
    url: '/fang/fang/edit',
    method: 'POST',
    params,
  });
}




// 获取租房指定详情
export function View(params) {
  return http.request({
    url: '/fang/fang/view',
    method: 'GET',
    params,
  });
}



// 导出租房
export function Export(params) {
  jumpExport('/fang/fang/export', params);
}
