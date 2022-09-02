<script lang="ts" setup>
import { fetchCountsByNoun } from "~~/gateway/getCountsByNoun";
import { CountsByNoun } from "~~/types/noun";

const tokens = ref<CountsByNoun[]>([]);
const target = ref("");
const placeholder = "分析したい文章を入れてください";
const loading = ref(false);

watch(
  () => target.value,
  async () => await handleOnInput()
);

const handleOnInput = async () => {
  loading.value = true;
  const { data: nouns, pending } = await fetchCountsByNoun(target.value);
  tokens.value = nouns.value.sort((a, b) => b.counts - a.counts);
  loading.value = pending.value;
};
</script>

<template>
  <div>
    <AtomsInput v-model="target" :placeholder="placeholder" />
    <NounsTable v-if="!loading" :nouns="tokens" />
  </div>
</template>
