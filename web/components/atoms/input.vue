<script lang="ts" setup>
interface Props {
  placeholder: string;
}

interface Emits {
  (e: "update:modelValue", value: string): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();

const handleOnInput = (e: InputEvent): void => {
  emit("update:modelValue", (e.target as HTMLInputElement).value);
};
</script>

<template>
  <div class="input-wrap">
    <input type="text" @input="handleOnInput" :placeholder="placeholder" />
    <label>{{ placeholder }}</label>
  </div>
</template>

<style lang="scss" scoped>
.input-wrap {
  margin: 0;
  margin-bottom: 5px;
  padding-top: 12px;
  position: relative;
}
.input-wrap label {
  position: absolute;
  top: 14px;
  left: 4px;
  color: #aaa;
  transition: all 0.2s;
  cursor: text;
  font-family: sans-serif;
  font-size: 16px;
  line-height: 24px;
  display: none;
}
.input-wrap input {
  height: 24px;
  padding: 2px;
  border: none;
  border-bottom: 1px solid #ccc;
  font-family: sans-serif;
  font-size: 16px;
  line-height: 20px;
  color: #555;
  -webkit-font-smoothing: antialiased;
  outline: none;
  width: 100%;
}

.input-wrap input:focus,
.input-wrap input:not(:placeholder-shown) {
  border-bottom: 1px solid #aaa;
}
.input-wrap input:focus + label {
  display: none;
}

.input-wrap input:not(:placeholder-shown) + label {
  font-size: 11px;
  line-height: 12px;
  top: 0px;
  color: #777;
  display: block;
}
</style>
