<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="设备详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>分类名称</template>
          {{ formValue.title }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>图片</template>
          {{ formValue.image }}
        </n-descriptions-item>

        <n-descriptions-item label="状态">
          <n-tag
              :type="getOptionTag(options.sys_normal_disable, formValue?.status)"
              size="small"
              class="min-left-space"
          >{{ getOptionLabel(options.sys_normal_disable, formValue?.status) }}
          </n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>排序</template>
          {{ formValue.sort }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>网址</template>
          {{ formValue.url }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {useMessage} from 'naive-ui';
import {View} from '@/api/addons/game/device';
import {newState, options} from './model';
import {getOptionLabel, getOptionTag} from '@/utils/hotgo';

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
  formValue.value = await View({deviceId: id});
});
</script>

<style lang="less" scoped></style>
