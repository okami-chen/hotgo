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
  name: string;
  title: string;
  bank: string;
  organize: string;
  card_no: string;
  expire_at: string;
  code: string;
  remark: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export const defaultState: State = {
  id: 0,
  name: '',
  title: '',
  bank: '',
  organize: '',
  card_no: '',
  expire_at: '',
  code: '',
  remark: '',
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

export const options = ref<Options>({
  addons_card_bank: [],
  addons_card_organize: [],
  addons_card_name: [],
});

export const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入持卡',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入名称',
  },
  bank: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入银行',
  },
  organize: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入组织',
  },
  card_no: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入卡号',
  },
  expire_at: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入过期时间',
  },
  code: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入识别码',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NSelect',
    label: '持卡',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择持卡',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'bank',
    component: 'NSelect',
    label: '银行',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择银行',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'organize',
    component: 'NSelect',
    label: '组织',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择组织',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  {
    field: 'card_no',
    component: 'NInput',
    label: '卡号',
    componentProps: {
      placeholder: '请输入卡号',
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
    title: '持卡',
    key: 'name',
    render(row) {
      if (isNullObject(row.name)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.addons_card_name, row.name),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.addons_card_name, row.name),
        }
      );
    },
  },
  {
    title: '名称',
    key: 'title',
  },
  {
    title: '银行',
    key: 'bank',
    render(row) {
      if (isNullObject(row.bank)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.addons_card_bank, row.bank),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.addons_card_bank, row.bank),
        }
      );
    },
  },
  {
    title: '组织',
    key: 'organize',
    render(row) {
      if (isNullObject(row.organize)) {
        return ``;
      }
      return h(
        NTag,
        {
          style: {
            marginRight: '6px',
          },
          type: getOptionTag(options.value.addons_card_organize, row.organize),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.addons_card_organize, row.organize),
        }
      );
    },
  },
  {
    title: '卡号',
    key: 'card_no',
  },
  {
    title: '过期时间',
    key: 'expire_at',
    render(row) {
      return formatToDate(row.expire_at);
    },
  },
  {
    title: '识别码',
    key: 'code',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'addons_card_bank',
    'addons_card_organize',
    'addons_card_name',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'bank':
        item.componentProps.options = options.value.addons_card_bank;
        break;
      case 'organize':
        item.componentProps.options = options.value.addons_card_organize;
        break;
      case 'name':
        item.componentProps.options = options.value.addons_card_name;
        break;
     }
  }
}

await loadOptions();