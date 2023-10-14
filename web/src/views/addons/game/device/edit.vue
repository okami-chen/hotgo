<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
          v-model:show="isShowModal"
          :show-icon="false"
          preset="dialog"
          :title="params?.deviceId > 0 ? '编辑 #' + params?.deviceId : '添加'"
          :style="{
          width: dialogWidth,
        }"
      >
        <n-form
            :model="params"
            :rules="rules"
            ref="formRef"
            label-placement="left"
            :label-width="100"
            class="py-4"
        >
          <n-form-item label="分类名称" path="title">
            <n-input placeholder="请输入分类名称" v-model:value="params.title"/>
          </n-form-item>

          <n-form-item label="图片" path="image">
            <n-input placeholder="请输入图片" v-model:value="params.image"/>
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-select v-model:value="params.status" :options="options.sys_normal_disable"/>
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number placeholder="请输入排序" v-model:value="params.sort"/>
          </n-form-item>

          <n-form-item label="网址" path="url">
            <n-input placeholder="请输入网址" v-model:value="params.url"/>
          </n-form-item>


        </n-form>
        <template #action>
          <n-space>
            <n-button @click="closeForm">取消</n-button>
            <n-button type="info" :loading="formBtnLoading" @click="confirmForm">确定</n-button>
          </n-space>
        </template>
      </n-modal>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, watch} from 'vue';
import {Edit, MaxSort, View} from '@/api/addons/game/device';
import {newState, options, rules, State} from './model';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

const emit = defineEmits(['reloadTable', 'updateShowModal']);

interface Props {
  showModal: boolean;
  formParams?: State;
}

const props = withDefaults(defineProps<Props>(), {
  showModal: false,
  formParams: () => {
    return newState(null);
  },
});

const isShowModal = computed({
  get: () => {
    return props.showModal;
  },
  set: (value) => {
    emit('updateShowModal', value);
  },
});

const loading = ref(false);
const params = ref<State>(props.formParams);
const message = useMessage();
const formRef = ref<any>({});
const dialogWidth = ref('75%');
const formBtnLoading = ref(false);

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      Edit(params.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          isShowModal.value = false;
          emit('reloadTable');
        });
      });
    } else {
      message.error('请填写完整信息');
    }
    formBtnLoading.value = false;
  });
}

onMounted(async () => {
  adaModalWidth(dialogWidth);
});

function closeForm() {
  isShowModal.value = false;
}

function loadForm(value) {
  loading.value = true;

  // 新增
  if (value.deviceId < 1) {
    params.value = newState(value);
    MaxSort()
        .then((res) => {
          params.value.sort = res.sort;
        })
        .finally(() => {
          loading.value = false;
        });
    return;
  }

  // 编辑
  View({deviceId: value.deviceId})
      .then((res) => {
        params.value = res;
      })
      .finally(() => {
        loading.value = false;
      });
}

watch(
    () => props.formParams,
    (value) => {
      loadForm(value);
    }
);
</script>

<style lang="less"></style>
