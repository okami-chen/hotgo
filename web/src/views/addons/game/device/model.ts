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
  deviceId: number;
  title: string;
  image: string;
  status: number;
  sort: number;
  url: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState = {
  deviceId: 0,
  title: '',
  image: '',
  status: 1,
  sort: 0,
  url: '',
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
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入分类名称',
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
  url: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入网址',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'deviceId',
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
      type: 'datetimerange',
      clearable: true,
      shortcuts: defRangeShortcuts(),
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '自动编号',
    key: 'deviceId',
  },
  {
    title: '分类名称',
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
    title: '网址',
    key: 'url',
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
