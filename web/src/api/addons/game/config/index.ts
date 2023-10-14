import { http } from '@/utils/http/axios';

export function getConfig(params) {
  return http.request({
    url: '/game/config/get',
    method: 'get',
    params,
  });
}

export function updateConfig(params) {
  return http.request({
    url: '/game/config/update',
    method: 'post',
    params,
  });
}
