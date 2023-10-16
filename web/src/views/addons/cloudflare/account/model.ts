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
  email: string;
  token: string;
  remark: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState: State = {
  id: 0,
  email: '',
  token: '',
  remark: '',
  createdAt: '',
  updatedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


export const rules = {
  email: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
  token: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入令牌',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'id',
    component: 'NInputNumber',
    label: '自动编号',
    componentProps: {
      placeholder: '请输入自动编号',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'email',
    component: 'NInput',
    label: 'email',
    componentProps: {
      placeholder: '请输入email',
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
    title: 'email',
    key: 'email',
  },
  {
    title: '令牌',
    key: 'token',
  },
  {
    title: '备注',
    key: 'remark',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
];