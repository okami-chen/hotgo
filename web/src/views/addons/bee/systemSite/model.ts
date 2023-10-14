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
  messageType: number;
  sort: number;
  siteName: string;
  siteUrl: string;
  siteLogo: string;
  currencys: string;
  languages: string;
  siteRemark: string;
  createdAt: string;
  updatedAt: string;
}

export const defaultState = {
  id: 0,
  messageType: 0,
  sort: 0,
  siteName: '',
  siteUrl: '',
  siteLogo: '',
  currencys: '',
  languages: '',
  siteRemark: '',
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
  messageType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入客服类型',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
  siteName: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入网站名称',
  },
  siteUrl: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入网站网址',
  },
  currencys: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入货币，多个逗号分隔',
  },
  languages: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入语言，多个逗号分隔',
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
    title: '客服类型',
    key: 'messageType',
  },
  {
    title: '排序',
    key: 'sort',
  },
  {
    title: '网站名称',
    key: 'siteName',
  },
  {
    title: '网站网址',
    key: 'siteUrl',
  },
  {
    title: '网站logo',
    key: 'siteLogo',
  },
  {
    title: '货币，多个逗号分隔',
    key: 'currencys',
  },
  {
    title: '语言，多个逗号分隔',
    key: 'languages',
  },
  {
    title: '网站备注',
    key: 'siteRemark',
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