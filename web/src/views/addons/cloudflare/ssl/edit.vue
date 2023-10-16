<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
       :title="params?.id > 0 ? '编辑 #' + params?.id : '添加'"
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
          <n-form-item label="账号" path="account_id">
            <n-input-number placeholder="请输入账号" v-model:value="params.account_id" />
          </n-form-item>

          <n-form-item label="域名" path="domain">
          <n-input placeholder="请输入域名" v-model:value="params.domain" />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-select v-model:value="params.status" :options="options.cloudflare_status" />
          </n-form-item>

          <n-form-item label="域名" path="domain_id">
          <n-input placeholder="请输入域名" v-model:value="params.domain_id" />
          </n-form-item>

          <n-form-item label="证书" path="cert_id">
          <n-input placeholder="请输入证书" v-model:value="params.cert_id" />
          </n-form-item>

          <n-form-item label="证书" path="cert_sub_id">
          <n-input placeholder="请输入证书" v-model:value="params.cert_sub_id" />
          </n-form-item>

          <n-form-item label="类型" path="type">
            <n-select v-model:value="params.type" :options="options.cloudflare_type" />
          </n-form-item>

          <n-form-item label="签发组织" path="issuer">
            <n-select v-model:value="params.issuer" :options="options.cloudflare_issuer" />
          </n-form-item>

          <n-form-item label="授权组织" path="authority">
            <n-select v-model:value="params.authority" :options="options.cloudflare_authority" />
          </n-form-item>

          <n-form-item label="签名方式" path="signature">
            <n-select v-model:value="params.signature" :options="options.cloudflare_signature" />
          </n-form-item>

          <n-form-item label="过期时间" path="expire_at">
            <DatePicker v-model:formValue="params.expire_at" type="datetime" />
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
  import { onMounted, ref, computed, watch } from 'vue';
  import { Edit, View } from '@/api/addons/cloudflare/ssl';
  import DatePicker from '@/components/DatePicker/datePicker.vue';
  import { rules, options, State, newState } from './model';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

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
    // 新增
    if (value.id < 1) {
      params.value = newState(value);
      loading.value = false;
      return;
    }

    loading.value = true;
    // 编辑
    View({ id: value.id })
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