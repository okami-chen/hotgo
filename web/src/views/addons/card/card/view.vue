<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="卡片详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item label="持卡">
          <n-tag
            :type="getOptionTag(options.addons_card_name, formValue?.name)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.addons_card_name, formValue?.name) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>名称</template>
          {{ formValue.title }}
        </n-descriptions-item>

        <n-descriptions-item label="银行">
          <n-tag
            :type="getOptionTag(options.addons_card_bank, formValue?.bank)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.addons_card_bank, formValue?.bank) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item label="组织">
          <n-tag
            :type="getOptionTag(options.addons_card_organize, formValue?.organize)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.addons_card_organize, formValue?.organize) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>卡号</template>
          {{ formValue.card_no }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>过期时间</template>
          {{ formValue.expire_at }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>识别码</template>
          {{ formValue.code }}
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