<template>
  <div>
    <div class="n-layout-page-header">
      <n-card :bordered="false" title="证书详情"> <!-- CURD详情页--> </n-card>
    </div>
    <n-card :bordered="false" class="proCard mt-4" size="small" :segmented="{ content: true }">
      <n-descriptions label-placement="left" class="py-2" column="4">
        <n-descriptions-item>
          <template #label>账号</template>
          {{ formValue.account_id }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>域名</template>
          {{ formValue.domain }}
        </n-descriptions-item>

        <n-descriptions-item label="状态">
          <n-tag
            :type="getOptionTag(options.cloudflare_status, formValue?.status)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.cloudflare_status, formValue?.status) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>域名</template>
          {{ formValue.domain_id }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>证书</template>
          {{ formValue.cert_id }}
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>证书</template>
          {{ formValue.cert_sub_id }}
        </n-descriptions-item>

        <n-descriptions-item label="类型">
          <n-tag
            :type="getOptionTag(options.cloudflare_type, formValue?.type)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.cloudflare_type, formValue?.type) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item label="签发组织">
          <n-tag
            :type="getOptionTag(options.cloudflare_issuer, formValue?.issuer)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.cloudflare_issuer, formValue?.issuer) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item label="授权组织">
          <n-tag
            :type="getOptionTag(options.cloudflare_authority, formValue?.authority)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.cloudflare_authority, formValue?.authority) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item label="签名方式">
          <n-tag
            :type="getOptionTag(options.cloudflare_signature, formValue?.signature)"
            size="small"
            class="min-left-space"
            >{{ getOptionLabel(options.cloudflare_signature, formValue?.signature) }}</n-tag
          >
        </n-descriptions-item>

        <n-descriptions-item>
          <template #label>过期时间</template>
          {{ formValue.expire_at }}
        </n-descriptions-item>


      </n-descriptions>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
  import { computed, onMounted, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { useMessage } from 'naive-ui';
  import { View } from '@/api/addons/cloudflare/ssl';
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