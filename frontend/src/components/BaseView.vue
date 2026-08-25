<script setup lang="ts">
import { ref } from 'vue'
import FilesList from "./FilesList.vue";
import Preview from "./Preview.vue";
import RulesView from "./RulesView.vue";
let noFiles = ref(true);
noFiles.value = false;
</script>

<template>
  <div id="base-view" class="base-view" >
    <section class="empty-section" v-if="noFiles">
      <p>将文件或文件夹拖拽到这里</p>
      <div class="btn-group">
        <button class="btn btn-file">导入文件</button>
        <button class="btn btn-folder">导入文件夹</button>
      </div>
    </section>
    <section class="file-section" v-else>
      <FilesList/>
      <Preview/>
      <RulesView/>
    </section>
  </div>
</template>

<style scoped>
.base-view {
  width: 100vw;
  height: 100vh;
}

.empty-section {
  flex: 1;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center;
  background-color: #a918ae;
}

.btn {
  margin: 5px;
  width: 100px;
  height: 40px;
  border-radius: 5px;
  border: 2px solid #fff;
  background-color: #9c6423;
  color: #fff;
}

.file-section {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: row nowrap;
  overflow: hidden;
}
</style>

