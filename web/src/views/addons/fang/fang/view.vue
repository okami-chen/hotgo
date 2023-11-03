<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="租房详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>编号</template>
          {{ formValue.sid }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>小区</template>
          {{ formValue.village }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>标题</template>
          {{ formValue.title }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>价格</template>
          {{ formValue.price }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>户型</template>
          {{ formValue.houseType }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>面积</template>
          {{ formValue.area }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>朝向</template>
          {{ formValue.toWard }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>电梯</template>
          {{ formValue.lift }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>用水</template>
          {{ formValue.water }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>用电</template>
          {{ formValue.electricity }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>租期</template>
          {{ formValue.tenancy }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>核验</template>
          {{ formValue.verify }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>入住</template>
          {{ formValue.checkIn }}
        </n-descriptions-item>

        <n-descriptions-item label="旗帜">
          <n-tag
            :type="getOptionTag(options.hefei_zf_flag, formValue?.flag)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.hefei_zf_flag, formValue?.flag) }}</n-tag
          >
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
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/fang/fang';
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
      message.error('编号不正确，请检查！');
      return;
    }
    formValue.value = await View({ id: id });
  });
</script>

<style lang="less" scoped></style>