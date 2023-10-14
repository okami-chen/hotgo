import { h, ref } from 'vue';
import {NTag } from 'naive-ui';
import { cloneDeep } from 'lodash-es';
import { FormSchema } from '@/components/Form';
import { Dicts } from '@/api/dict/dict';

import {isNullObject } from '@/utils/is';
import {defShortcuts } from '@/utils/dateUtil';
import { getOptionLabel, getOptionTag, Options } from '@/utils/hotgo';


export interface State {
  gameId: number;
  gameSku: string;
  name: string;
  title: string;
  status: number;
  sort: number;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState = {
  gameId: 0,
  gameSku: '',
  name: '',
  title: '',
  status: 1,
  sort: 0,
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

export const options = ref<Options>({
  sys_normal_disable: [],
});

export const rules = {
  gameSku: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入唯一值',
  },
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入游戏全名',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入游戏名称',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态',
  },
  sort: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入排序',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'gameId',
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
    field: 'status',
    component: 'NSelect',
    label: '状态',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择状态',
      options: [],
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
    key: 'gameId',
  },
  {
    title: '唯一值',
    key: 'gameSku',
  },
  {
    title: '游戏全名',
    key: 'name',
  },
  {
    title: '游戏名称',
    key: 'title',
  },
  {
    title: '状态',
    key: 'status',
    render(row) {
      if (isNullObject(row.status)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.yes_or_no, row.status),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.yes_or_no, row.status),
        }
      );
    },
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

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'yes_or_no',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'status':
        item.componentProps.options = options.value.yes_or_no;
        break;
     }
  }
}

await loadOptions();
