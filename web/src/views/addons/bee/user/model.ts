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
  uid: number;
  siteId: number;
  shareUid: number;
  email: string;
  password: string;
  grade: number;
  status: number;
  consumeMoney: number;
  nonMoney: number;
  freezeMoney: number;
  amount: number;
  payment: string;
  defaultPayment: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState = {
  uid: 0,
  siteId: 0,
  shareUid: 0,
  email: '',
  password: '',
  grade: 0,
  status: 1,
  consumeMoney: 0,
  nonMoney: 0,
  freezeMoney: 0,
  amount: 0,
  payment: '',
  defaultPayment: '',
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
  yes_or_no: [],
});

export const rules = {
  siteId: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入站点编号',
  },
  email: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.email,
  },
  password: {
    required: false,
    trigger: ['blur', 'input'],
    type: 'string',
    validator: validate.password,
  },
  grade: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入等级',
  },
  status: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入状态',
  },
  consumeMoney: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入消费金额',
  },
  nonMoney: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入非货币',
  },
  freezeMoney: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入冻结金额',
  },
  amount: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入账户余额',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'uid',
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
    title: '编号',
    key: 'uid',
  },
  {
    title: '站点',
    key: 'siteId',
  },
  {
    title: '邮箱',
    key: 'email',
  },
  {
    title: '等级',
    key: 'grade',
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
    title: '消费',
    key: 'consumeMoney',
  },
  {
    title: '非货',
    key: 'nonMoney',
  },
  {
    title: '冻结',
    key: 'freezeMoney',
  },
  {
    title: '余额',
    key: 'amount',
  },
  {
    title: '创建',
    key: 'createdAt',
  }
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
