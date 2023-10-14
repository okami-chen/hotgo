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
          <n-form-item label="用户" path="uid">
            <n-input-number placeholder="请输入用户" v-model:value="params.uid" />
          </n-form-item>

          <n-form-item label="用户名" path="username">
          <n-input placeholder="请输入用户名" v-model:value="params.username" />
          </n-form-item>

          <n-form-item label="用户名" path="firstName">
          <n-input placeholder="请输入用户名" v-model:value="params.firstName" />
          </n-form-item>

          <n-form-item label="用户姓" path="lastName">
          <n-input placeholder="请输入用户姓" v-model:value="params.lastName" />
          </n-form-item>

          <n-form-item label="手机" path="mobile">
          <n-input placeholder="请输入手机" v-model:value="params.mobile" />
          </n-form-item>

          <n-form-item label="国家" path="country">
          <n-input placeholder="请输入国家" v-model:value="params.country" />
          </n-form-item>

          <n-form-item label="头像" path="avatar">
            <n-input type="textarea" placeholder="头像" v-model:value="params.avatar" />
          </n-form-item>

          <n-form-item label="账号" path="account">
          <n-input placeholder="请输入账号" v-model:value="params.account" />
          </n-form-item>

          <n-form-item label="性别" path="gender">
          <n-input placeholder="请输入性别" v-model:value="params.gender" />
          </n-form-item>

          <n-form-item label="联系" path="socialMedia">
          <n-input placeholder="请输入联系" v-model:value="params.socialMedia" />
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
  import { Edit, View } from '@/api/addons/bee/userInfo';
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