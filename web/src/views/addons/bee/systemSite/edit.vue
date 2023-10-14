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
          <n-form-item label="客服类型" path="messageType">
            <n-input-number placeholder="请输入客服类型" v-model:value="params.messageType" />
          </n-form-item>

          <n-form-item label="排序" path="sort">
            <n-input-number placeholder="请输入排序" v-model:value="params.sort" />
          </n-form-item>

          <n-form-item label="网站名称" path="siteName">
          <n-input placeholder="请输入网站名称" v-model:value="params.siteName" />
          </n-form-item>

          <n-form-item label="网站网址" path="siteUrl">
          <n-input placeholder="请输入网站网址" v-model:value="params.siteUrl" />
          </n-form-item>

          <n-form-item label="网站logo" path="siteLogo">
          <n-input placeholder="请输入网站logo" v-model:value="params.siteLogo" />
          </n-form-item>

          <n-form-item label="货币，多个逗号分隔" path="currencys">
            <n-input type="textarea" placeholder="货币，多个逗号分隔" v-model:value="params.currencys" />
          </n-form-item>

          <n-form-item label="语言，多个逗号分隔" path="languages">
            <n-input type="textarea" placeholder="语言，多个逗号分隔" v-model:value="params.languages" />
          </n-form-item>

          <n-form-item label="网站备注" path="siteRemark">
            <n-input type="textarea" placeholder="网站备注" v-model:value="params.siteRemark" />
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
  import { Edit, MaxSort, View } from '@/api/addons/bee/systemSite';
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
    loading.value = true;

    // 新增
    if (value.id < 1) {
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