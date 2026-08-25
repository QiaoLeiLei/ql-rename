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
  flex: 1;
  width: 30%;
  height: 100%;
  background-color: #5d831b;
}

.preview-title{
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: bold;
  padding: 10px;
  border: 2px solid lightcyan;
}

.preview-description{
  display: flex;
  flex-direction: row;
}
</style>