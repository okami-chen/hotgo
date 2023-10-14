<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="站群详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>客服类型</template>
          {{ formValue.messageType }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>排序</template>
          {{ formValue.sort }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>网站名称</template>
          {{ formValue.siteName }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>网站网址</template>
          {{ formValue.siteUrl }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>网站logo</template>
          {{ formValue.siteLogo }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>货币，多个逗号分隔</template>
          <span v-html="formValue.currencys"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>语言，多个逗号分隔</template>
          <span v-html="formValue.languages"></span></n-descriptions-item>

        <n-descriptions-item>
          <template #label>网站备注</template>
          <span v-html="formValue.siteRemark"></span></n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/bee/systemSite';
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