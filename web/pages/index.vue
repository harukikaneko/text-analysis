<script lang="ts" setup>
import { fetchCountsByNoun } from "~~/gateway/getCountsByNoun";
import { CountsByNoun } from "~~/types/noun";

const tokens = ref<CountsByNoun[]>([]);
const handleOnChange = async (x: InputEvent) => {
  const value = (x.target as HTMLInputElement).value;
  const { data: nouns } = await fetchCountsByNoun(value);
  tokens.value = nouns.value.sort((a, b) => b.counts - a.counts);
};
</script>

<template>
  <div>
    <input type="text" @input="handleOnChange" />
    <table>
      <thead>
        <tr>
          <th colspan="2">nouns</th>
          <th>count</th>
        </tr>
      </thead>
      <tbody v-for="(item, index) in tokens" :key=index>
        <tr>
          <td>{{ item.noun }}</td>
          <td>{{ item.counts }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
