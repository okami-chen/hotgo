import {h, ref} from 'vue';
import {NTag} from 'naive-ui';
import {cloneDeep} from 'lodash-es';
import {FormSchema} from '@/components/Form';
import {Dicts} from '@/api/dict/dict';

import {isNullObject} from '@/utils/is';
import {getOptionLabel, getOptionTag, Options} from '@/utils/hotgo';


export interface State {
    id: number;
    account_id: number;
    domain: string;
    status: string;
    domain_id: string;
    cert_id: string;
    cert_sub_id: string;
    type: string;
    issuer: string;
    authority: string;
    signature: string;
    created_at: string;
    updated_at: string;
    expire_at: string;
}

export const defaultState: State = {
    id: 0,
    account_id: 0,
    domain: '',
    status: 1,
    domain_id: '',
    cert_id: '',
    cert_sub_id: '',
    type: '',
    issuer: '',
    authority: '',
    signature: '',
    created_at: '',
    updated_at: '',
    expire_at: '',
};

export function newState(state: State | null): State {
    if (state !== null) {
        return cloneDeep(state);
    }
    return cloneDeep(defaultState);
}

export const options = ref<Options>({
    cloudflare_type: [],
    cloudflare_status: [],
    cloudflare_issuer: [],
    cloudflare_authority: [],
    cloudflare_signature: [],
});

export const rules = {
    account_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'number',
        message: '请输入账号',
    },
    domain: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入域名',
    },
    status: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入状态',
    },
    domain_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入域名',
    },
    cert_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入证书',
    },
    cert_sub_id: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入证书',
    },
    type: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入类型',
    },
    issuer: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入签发组织',
    },
    authority: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入授权组织',
    },
    signature: {
        required: true,
        trigger: ['blur', 'input'],
        type: 'string',
        message: '请输入签名方式',
    },
};

export const schemas = ref<FormSchema[]>([
    {
        field: 'domain',
        component: 'NInput',
        label: '域名',
        componentProps: {
            placeholder: '请输入域名',
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
        field: 'type',
        component: 'NSelect',
        label: '类型',
        defaultValue: null,
        componentProps: {
            placeholder: '请选择类型',
            options: [],
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'issuer',
        component: 'NSelect',
        label: '签发组织',
        defaultValue: null,
        componentProps: {
            placeholder: '请选择签发组织',
            options: [],
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'authority',
        component: 'NSelect',
        label: '授权组织',
        defaultValue: null,
        componentProps: {
            placeholder: '请选择授权组织',
            options: [],
            onUpdateValue: (e: any) => {
                console.log(e);
            },
        },
    },
    {
        field: 'signature',
        component: 'NSelect',
        label: '签名方式',
        defaultValue: null,
        componentProps: {
            placeholder: '请选择签名方式',
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
        title: '域名',
        key: 'domain',
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
                    type: getOptionTag(options.value.cloudflare_status, row.status),
                    bordered: false,
                },
                {
                    default: () => getOptionLabel(options.value.cloudflare_status, row.status),
                }
            );
        },
    },
    {
        title: '类型',
        key: 'type',
        render(row) {
            if (isNullObject(row.type)) {
                return ``;
            }
            return h(
                NTag,
                {
                    style: {
                        marginRight: '6px',
                    },
                    type: getOptionTag(options.value.cloudflare_type, row.type),
                    bordered: false,
                },
                {
                    default: () => getOptionLabel(options.value.cloudflare_type, row.type),
                }
            );
        },
    },
    {
        title: '签发组织',
        key: 'issuer',
        render(row) {
            if (isNullObject(row.issuer)) {
                return ``;
            }
            return h(
                NTag,
                {
                    style: {
                        marginRight: '6px',
                    },
                    type: getOptionTag(options.value.cloudflare_issuer, row.issuer),
                    bordered: false,
                },
                {
                    default: () => getOptionLabel(options.value.cloudflare_issuer, row.issuer),
                }
            );
        },
    },
    {
        title: '授权组织',
        key: 'authority',
        render(row) {
            if (isNullObject(row.authority)) {
                return ``;
            }
            return h(
                NTag,
                {
                    style: {
                        marginRight: '6px',
                    },
                    type: getOptionTag(options.value.cloudflare_authority, row.authority),
                    bordered: false,
                },
                {
                    default: () => getOptionLabel(options.value.cloudflare_authority, row.authority),
                }
            );
        },
    },
    {
        title: '签名方式',
        key: 'signature',
        render(row) {
            if (isNullObject(row.signature)) {
                return ``;
            }
            return h(
                NTag,
                {
                    style: {
                        marginRight: '6px',
                    },
                    type: getOptionTag(options.value.cloudflare_signature, row.signature),
                    bordered: false,
                },
                {
                    default: () => getOptionLabel(options.value.cloudflare_signature, row.signature),
                }
            );
        },
    },
    {
        title: '过期时间',
        key: 'expire_at',
    },
];

async function loadOptions() {
    options.value = await Dicts({
        types: [
            'cloudflare_type',
            'cloudflare_status',
            'cloudflare_issuer',
            'cloudflare_authority',
            'cloudflare_signature',
        ],
    });
    for (const item of schemas.value) {
        switch (item.field) {
            case 'type':
                item.componentProps.options = options.value.cloudflare_type;
                break;
            case 'status':
                item.componentProps.options = options.value.cloudflare_status;
                break;
            case 'issuer':
                item.componentProps.options = options.value.cloudflare_issuer;
                break;
            case 'authority':
                item.componentProps.options = options.value.cloudflare_authority;
                break;
            case 'signature':
                item.componentProps.options = options.value.cloudflare_signature;
                break;
        }
    }
}

await loadOptions();
