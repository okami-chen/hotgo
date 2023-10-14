import {http, jumpExport} from '@/utils/http/axios';

// 获取设备列表
export function List(params) {
    return http.request({
        url: '/game/device/list',
        method: 'get',
        params,
    });
}

// 删除/批量删除设备
export function Delete(params) {
    return http.request({
        url: '/game/device/delete',
        method: 'POST',
        params,
    });
}


// 添加/编辑设备
export function Edit(params) {
    return http.request({
        url: '/game/device/edit',
        method: 'POST',
        params,
    });
}


// 修改设备状态
export function Status(params) {
    return http.request({
        url: '/game/device/status',
        method: 'POST',
        params,
    });
}


// 获取设备指定详情
export function View(params) {
    return http.request({
        url: '/game/device/view',
        method: 'GET',
        params,
    });
}


// 获取设备最大排序
export function MaxSort() {
    return http.request({
        url: '/game/device/maxSort',
        method: 'GET',
    });
}


// 导出设备
export function Export(params) {
    jumpExport('/game/device/export', params);
}
