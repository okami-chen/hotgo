import {ref} from 'vue';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';


export interface State {
    id: number;
    account_id: number;
    domain_id: string;
    domain: string;
    created_at: string;
    updated_at: string;
}

export const defaultState: State = {
    id: 0,
    account_id: 0,
    domain_id: '',
    domain: '',
    created_at: '',
    updated_at: '',
};

export function newState(state: State | null): State {
    if (state !== null) {
        return cloneDeep(state);
    }
    return cloneDeep(defaultState);
}


export const rules = {
    account_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'number',
        message: '请输入关联账号',
    },
    domain_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入域名编号',
    },
    domain: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入域名名称',
    },
};

export const schemas = ref<FormSchema[]>([
    {
        field: 'domain',
        component: 'NInput',
        label: '域名名称',
        componentProps: {
            placeholder: '请输入域名名称',
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
        title: '域名名称',
        key: 'domain',
    },
    {
        title: '创建时间',
        key: 'created_at',
    },
];
