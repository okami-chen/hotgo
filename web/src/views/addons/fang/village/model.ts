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
  data_1: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export const defaultState: State = {
  id: 0,
  name: '',
  data_1: '',
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
  hefei_zf_area: [],
});

export const rules = {
  name: {
    required: true,
    trigger: ['blur', 'input'],
    type: 'string',
    message: '请输入名称',
  },
};

export const schemas = ref<FormSchema[]>([
  {
    field: 'name',
    component: 'NSelect',
    label: '名称',
    defaultValue: null,
    componentProps: {
      placeholder: '请选择名称',
      options: [],
      onUpdateValue: (e: any) => {
        console.log(e);
      },
    },
  },
]);

export const columns = [
  {
    title: '编号',
    key: 'id',
  },
  {
    title: '名称',
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
          type: getOptionTag(options.value.hefei_zf_area, row.name),
          bordered: false,
        },
        {
          default: () => getOptionLabel(options.value.hefei_zf_area, row.name),
        }
      );
    },
  },
  {
    title: '创建时间',
    key: 'created_at',
  },
  {
    title: '更新时间',
    key: 'updated_at',
  },
];

async function loadOptions() {
  options.value = await Dicts({
    types: [
      'hefei_zf_area',
   ],
  });
  for (const item of schemas.value) {
    switch (item.field) {
      case 'name':
        item.componentProps.options = options.value.hefei_zf_area;
        break;
     }
  }
}

await loadOptions();
