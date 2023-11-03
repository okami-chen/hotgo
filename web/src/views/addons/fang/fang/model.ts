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
  sid: string;
  village: string;
  title: string;
  price: number;
  houseType: string;
  area: string;
  toWard: string;
  lift: string;
  water: string;
  electricity: string;
  tenancy: string;
  verify: string;
  checkIn: string;
  flag: number;
  url: string;
  createdAt: string;
  updatedAt: string;
  deletedAt: string;
}

export const defaultState: State = {
  id: 0,
  sid: '',
  village: '',
  title: '',
  price: 0,
  houseType: '',
  area: '',
  toWard: '',
  lift: '',
  water: '',
  electricity: '',
  tenancy: '',
  verify: '',
  checkIn: '',
  flag: 1,
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
  hefei_zf_flag: [],
});

export const rules = {
  sid: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入编号',
  },
  village: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入小区',
  },
  title: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入标题',
  },
  price: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    validator: validate.amount,
    message: '请输入价格',
  },
  houseType: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入户型',
  },
  toWard: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入朝向',
  },
  flag: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'number',
    message: '请输入旗帜',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'title',
    component: 'NInput',
    label: '标题',
    componentProps: {
      placeholder: '请输入标题',
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'village',
  //   component: 'NInput',
  //   label: '小区',
  //   componentProps: {
  //     placeholder: '请输入小区',
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  {
    field: 'village',
    component: 'NSelect',
    label: '小区',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择小区',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
  // {
  //   field: 'houseType',
  //   component: 'NInput',
  //   label: '户型',
  //   componentProps: {
  //     placeholder: '请输入户型',
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  // {
  //   field: 'toWard',
  //   component: 'NInput',
  //   label: '朝向',
  //   componentProps: {
  //     placeholder: '请输入朝向',
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
  // {
  //   field: 'flag',
  //   component: 'NSelect',
  //   label: '旗帜',
  //   defaultValue: null,
  //   componentProps: {
  //     placeholder: '请选择旗帜',
  //     options: [],
  //     onUpdateValue: (e: any) => {
  //       console.log(e);
  //     },
  //   },
  // },
]);

export const columns = [
  {
    title: '编号',
    key: 'id',
  },
  {
    title: '小区',
    key: 'village',
    width: 160,
  },
  {
    title: '价格',
    key: 'price',
  },
  {
    title: '户型',
    key: 'houseType',
  },
  {
    title: '面积',
    key: 'area',
  },
  {
    title: '朝向',
    key: 'toWard',
  },
  {
    title: '电梯',
    key: 'lift',
  },
  {
    title: '用水',
    key: 'water',
  },
  {
    title: '用电',
    key: 'electricity',
  },
  // {
  //   title: '入住',
  //   key: 'checkIn',
  // },
  // {
  //   title: '旗帜',
  //   key: 'flag',
  //   render(row) {
  //     if (isNullObject(row.flag)) {
  //       return ``;
  //     }
  //     return h(
  //       NTag,
  //       {
  //         style: {
  //           marginRight: '6px',
  //         },
  //         type: getOptionTag(options.value.hefei_zf_flag, row.flag),
  //         bordered: false,
  //       },
  //       {
  //         default: () => getOptionLabel(options.value.hefei_zf_flag, row.flag),
  //       }
  //     );
  //   },
  // },
  {
    title: '发布时间',
    key: 'createdAt',
    width: 160,
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'hefei_zf_flag',
      'hefei_zf_area'
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'flag':
        item.componentProps.options = options.value.hefei_zf_flag;
        break;
      case 'village':
        item.componentProps.options = options.value.hefei_zf_area;
        break;
     }
  }
}

await loadOptions();
