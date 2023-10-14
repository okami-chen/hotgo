<template>
  <div>
    <n-spin :show="loading" description="请稍候...">
      <n-modal
        v-model:show="isShowModal"
        :show-icon="false"
        preset="dialog"
       :title="params?.uid > 0 ? '编辑 #' + params?.uid : '添加'"
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
          <n-form-item label="站点编号" path="siteId">
            <n-input-number placeholder="请输入站点编号" v-model:value="params.siteId" />
          </n-form-item>

          <n-form-item label="电子邮箱" path="email">
          <n-input placeholder="请输入电子邮箱" v-model:value="params.email" />
          </n-form-item>

          <n-form-item label="密码" path="password">
          <n-input placeholder="请输入密码" v-model:value="params.password" />
          </n-form-item>

          <n-form-item label="等级" path="grade">
            <n-input-number placeholder="请输入等级" v-model:value="params.grade" />
          </n-form-item>

          <n-form-item label="状态" path="status">
            <n-select v-model:value="params.status" :options="options.yes_or_no" />
          </n-form-item>

          <n-form-item label="消费金额" path="consumeMoney">
            <n-input-number placeholder="请输入消费金额" v-model:value="params.consumeMoney" />
          </n-form-item>

          <n-form-item label="非货币" path="nonMoney">
            <n-input-number placeholder="请输入非货币" v-model:value="params.nonMoney" />
          </n-form-item>

          <n-form-item label="冻结金额" path="freezeMoney">
            <n-input-number placeholder="请输入冻结金额" v-model:value="params.freezeMoney" />
          </n-form-item>

          <n-form-item label="账户余额" path="amount">
            <n-input-number placeholder="请输入账户余额" v-model:value="params.amount" />
          </n-form-item>

          <n-form-item label="支付方式" path="payment">
          <n-input placeholder="请输入支付方式" v-model:value="params.payment" />
          </n-form-item>

          <n-form-item label="默认支付方式" path="defaultPayment">
          <n-input placeholder="请输入默认支付方式" v-model:value="params.defaultPayment" />
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
  import { Edit, View } from '@/api/addons/bee/user';
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
    if (value.uid < 1) {
      params.value = newState(value);
      loading.value = false;
      return;
    }

    loading.value = true;
    // 编辑
    View({ uid: value.uid })
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
