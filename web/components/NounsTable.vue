<script lang="ts" setup>
import { CountsByNoun } from "~~/types/noun";

interface Props {
  nouns: CountsByNoun[];
}

defineProps<Props>();
</script>

<template>
  <table>
    <thead>
      <tr>
        <th>nouns</th>
        <th>count</th>
      </tr>
    </thead>
    <tbody v-for="(item, index) in nouns" :key="index">
      <tr>
        <td>{{ item.noun }}</td>
        <td>{{ item.counts }}</td>
      </tr>
    </tbody>
  </table>
</template>

<style>
table {
  width: 100%;
  text-align: center;
  border-collapse: collapse;
  border-spacing: 0;
  border: 1px solid #ccc;
  color: #353535;
}

table thead {
  background: #ede6d3;
}

table th,
table td {
  border-right: 1px solid #ccc;
  padding: 10px;
}
</style>
