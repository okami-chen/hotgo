import { h, ref } from 'vue';
import { NAvatar, NImage, NTag, NSwitch, NRate } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import { isArray, isNullObject } from '@/utils/is';
import { getFileExt } from '@/utils/urlUtils';
import { defRangeShortcuts, defShortcuts, formatToDate } from '@/utils/dateUtil';
import { validate } from '@/utils/validateUtil';
import { getOptionLabel, getOptionTag, Options, errorImg } from '@/utils/hotgo';


export interface State {
  id: number;
  country: string;
  country_name: string;
  dialling_code: string;
  img: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export const defaultState: State = {
  id: 0,
  country: '',
  country_name: '',
  dialling_code: '',
  img: '',
  created_at: '',
  updated_at: '',
  deleted_at: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


export const rules = {
  country: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入国家缩写',
  },
  country_name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入国家名称全程',
  },
  dialling_code: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入国家区号',
  },
  img: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入国旗图片',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'country',
    component: 'NInput',
    label: '国家缩写',
    componentProps: {
      placeholder: '请输入国家缩写',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'country_name',
    component: 'NInput',
    label: '国家名称全程',
    componentProps: {
      placeholder: '请输入国家名称全程',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'dialling_code',
    component: 'NInput',
    label: '国家区号',
    componentProps: {
      placeholder: '请输入国家区号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '自动编号',
    key: 'id',
  },
  {
    title: '国家缩写',
    key: 'country',
  },
  {
    title: '国家名称全程',
    key: 'country_name',
  },
  {
    title: '国家区号',
    key: 'dialling_code',
  },
  {
    title: '国旗图片',
    key: 'img',
  },
];