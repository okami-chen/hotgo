<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="卡片详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>持卡</template>
          {{ formValue.name }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>名称</template>
          {{ formValue.title }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>银行</template>
          {{ formValue.bank }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>组织</template>
          {{ formValue.organize }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>卡号</template>
          {{ formValue.cardNo }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>过期时间</template>
          {{ formValue.expireAt }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>识别码</template>
          {{ formValue.code }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>备注</template>
          {{ formValue.remark }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/card/card';
  import { newState, options } from './model';
  import { getOptionLabel, getOptionTag } from '@/utils/hotgo';
  import { getFileExt } from '@/utils/urlUtils';

  const message = useMessage();
  const router = useRouter();
  const id = Number(router.currentRoute.value.params.id);
  const formValue = ref(newState(null));
  const fileAvatarCSS = computed(() => {
    return {
      '--n-merged-size': `var(--n-avatar-size-override, 80px)`,
      '--n-font-size': `18px`,
    };
  });

  //下载
  function download(url: string) {
    window.open(url);
  }

  onMounted(async () => {
    if (id < 1) {
      message.error('自动编号不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>