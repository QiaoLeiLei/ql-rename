<script setup lang="ts">
import { ref } from 'vue'
let oldNames = ref<string[]>([]);
let newNames = ref<string[]>([]);
</script>

<template>
  <div id="preview" class="preview" >
    <div class="preview-title">
      <p>预览</p>
      <div class="preview-description">
        <div>原始名</div>
        <div>→</div>
        <div>新名称</div>
      </div>
    </div>

    <section>
    </section>
  </div>
</template>

<style scoped>
.preview {
  width: 26%;
  height: calc(100% - 18px);
  background-color: rgb(36, 44, 61);
  border-radius: 15px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
  margin-right: 6px;
  display: flex;
  overflow: hidden;
  flex-flow: column nowrap;
  justify-content: space-between;
}

.preview-title{
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: bold;
  padding: 10px;
}

.preview-description{
  display: flex;
  flex-direction: row;
}
</style>