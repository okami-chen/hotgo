import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
    id: number;
    status: number;
    code: string;
    name: string;
    symbol: string;
    rate: number;
    created_at: string;
    updated_at: string;
    deleted_at: string;
}

export const defaultState: State = {
    id: 0,
    status: 1,
    code: '',
    name: '',
    symbol: '',
    rate: null,
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
    yes_or_no: [],
});

export const rules = {
    status: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'number',
        message: '请输入状态',
    },
    code: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入代码',
    },
    name: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入名称',
    },
    symbol: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入符号',
    },
    rate: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'number',
        message: '请输入汇率',
    },
};

export const schemas = ref<FormSchema[]>([
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
        field: 'code',
        component: 'NInput',
        label: '代码',
        componentProps: {
            placeholder: '请输入代码',
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'name',
        component: 'NInput',
        label: '名称',
        componentProps: {
            placeholder: '请输入名称',
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
        title: '代码',
        key: 'code',
    },
    {
        title: '名称',
        key: 'name',
    },
    {
        title: '符号',
        key: 'symbol',
    },
    {
        title: '汇率',
        key: 'rate',
    },
    {
        title: '更新时间',
        key: 'updated_at',
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
