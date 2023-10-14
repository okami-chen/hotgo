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
  uid: number;
  username: string;
  firstName: string;
  lastName: string;
  mobile: string;
  country: string;
  avatar: string;
  account: string;
  gender: string;
  socialMedia: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState = {
  id: 0,
  uid: 0,
  username: '',
  firstName: '',
  lastName: '',
  mobile: '',
  country: '',
  avatar: '',
  account: '',
  gender: '',
  socialMedia: '',
  createdAt: '',
  updatedAt: '',
  deletedAt: '',
};

export function newState(state: State | null): State {
  if (state !== null) {
    return cloneDeep(state);
  }
  return cloneDeep(defaultState);
}


export const rules = {
  uid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入用户',
  },
  mobile: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.phone,
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
    field: 'createdAt',
    component: 'NDatePicker',
    label: '创建时间',
    componentProps: {
      type: 'datetime',
      clearable: true,
      shortcuts: defShortcuts(),
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
    title: '用户',
    key: 'uid',
  },
  {
    title: '用户名',
    key: 'username',
  },
  {
    title: '用户名',
    key: 'firstName',
  },
  {
    title: '用户姓',
    key: 'lastName',
  },
  {
    title: '手机',
    key: 'mobile',
  },
  {
    title: '国家',
    key: 'country',
  },
  {
    title: '头像',
    key: 'avatar',
  },
  {
    title: '账号',
    key: 'account',
  },
  {
    title: '性别',
    key: 'gender',
  },
  {
    title: '联系',
    key: 'socialMedia',
  },
  {
    title: '创建时间',
    key: 'createdAt',
  },
  {
    title: '更新时间',
    key: 'updatedAt',
  },
];