import {http, jumpExport} from '@/utils/http/axios';

// 获取游戏列表
export function List(params) {
    return http.request({
        url: '/game/game/list',
        method: 'get',
        params,
    });
}

// 删除/批量删除游戏
export function Delete(params) {
    return http.request({
        url: '/game/game/delete',
        method: 'POST',
        params,
    });
}


// 添加/编辑游戏
export function Edit(params) {
    return http.request({
        url: '/game/game/edit',
        method: 'POST',
        params,
    });
}


// 修改游戏状态
export function Status(params) {
    return http.request({
        url: '/game/game/status',
        method: 'POST',
        params,
    });
}


// 获取游戏指定详情
export function View(params) {
    return http.request({
        url: '/game/game/view',
        method: 'GET',
        params,
    });
}


// 获取游戏最大排序
export function MaxSort() {
    return http.request({
        url: '/game/game/maxSort',
        method: 'GET',
    });
}


// 导出游戏
export function Export(params) {
    jumpExport('/game/game/export', params);
}
