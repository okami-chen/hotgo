import {ref} from 'vue';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';


export interface State {
    id: number;
    language: string;
    language_zh: string;
    code: string;
    image: string;
    created_at: string;
    updated_at: string;
    deleted_at: string;
}

export const defaultState: State = {
    id: 0,
    language: '',
    language_zh: '',
    code: '',
    image: '',
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


export const rules = {
    language: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入语言英文',
    },
    language_zh: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入语言中文',
    },
    code: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入语言简称',
    },
};

export const schemas = ref<FormSchema[]>([
    {
        field: 'language',
        component: 'NInput',
        label: '语言英文',
        componentProps: {
            placeholder: '请输入语言英文',
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'language_zh',
        component: 'NInput',
        label: '语言中文',
        componentProps: {
            placeholder: '请输入语言中文',
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'code',
        component: 'NInput',
        label: '语言简称',
        componentProps: {
            placeholder: '请输入语言简称',
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
        title: '语言英文',
        key: 'language',
    },
    {
        title: '语言中文',
        key: 'language_zh',
    },
    {
        title: '语言简称',
        key: 'code',
    },

];
