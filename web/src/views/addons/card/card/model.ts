import {ref} from 'vue';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {defRangeShortcuts, formatToDate} from '@/utils/dateUtil';
import {Dicts} from "@/api/dict/dict";
import {Options} from '@/utils/hotgo';


export interface State {
    id: number;
    name: string;
    title: string;
    bank: string;
    organize: string;
    cardNo: string;
    expireAt: string;
    code: string;
    remark: string;
    createdAt: string;
    updatedAt: string;
    deletedAt: string;
}

export const defaultState: State = {
    id: 0,
    name: '',
    title: '',
    bank: '',
    organize: '',
    cardNo: '',
    expireAt: '',
    code: '',
    remark: '',
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
    custom_bank: [],
    custom_name: [],
    custom_organize: [],
});

export const rules = {
    name: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入持卡人',
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
        message: '请选择银行',
    },
    organize: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请选择组织',
    },
    cardNo: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入卡号',
    },
    expireAt: {
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
        field: 'bank',
        component: 'NSelect',
        label: '银行',
        componentProps: {
            clearable: true,
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
        giProps: {
            span: 1 / 12,
        },
    },
    {
        field: 'organize',
        component: 'NSelect',
        label: '组织',
        componentProps: {
            clearable: true,
            shortcuts: defRangeShortcuts(),
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'name',
        component: 'NSelect',
        label: '持卡',
        componentProps: {
            clearable: true,
            shortcuts: defRangeShortcuts(),
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'cardNo',
        component: 'NInput',
        label: '卡号',
        componentProps: {
            clearable: true,
            shortcuts: defRangeShortcuts(),
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
]);

// @ts-ignore
export const columns = [
    {
        title: '编号',
        key: 'id',
        width: 100,
        resizable: true,
    },
    {
        title: '持卡',
        key: 'name',
        resizable: true,
    },
    {
        title: '名称',
        key: 'title',
        resizable: true,
    },
    {
        title: '银行',
        key: 'bank',
        resizable: true,
    },
    {
        title: '组织',
        key: 'organize',
        resizable: true,
    },
    {
        title: '卡号',
        key: 'cardNo',
        width: 260,
        resizable: true,
    },
    {
        title: '标识',
        key: 'code',
        resizable: true,
        edit: true,
        editComponent: 'NInput',
    },
    {
        title: '过期',
        key: 'expireAt',
        resizable: true,
        render(row) {
            return formatToDate(row.expireAt);
        },
    },
    {
        title: '备注',
        key: 'remark',
    },
];

async function loadOptions() {
    options.value = await Dicts({
        types: ['custom_bank', 'custom_organize', 'custom_name'],
    });
    for (const item of schemas.value) {
        switch (item.field) {
            case 'bank':
                item.componentProps.options = options.value.custom_bank;
                break;
            case 'organize':
                item.componentProps.options = options.value.custom_organize;
                break;
            case 'name':
                item.componentProps.options = options.value.custom_name;
                break;
        }
    }
}

await loadOptions();
