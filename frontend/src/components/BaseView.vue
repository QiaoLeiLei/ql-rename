<script setup lang="ts">
import FilesList from "./FilesList.vue";
import Preview from "./Preview.vue";
import RulesView from "./RulesView.vue";
import {dataCenter} from "../dataCenter";
import {importFiles,importFolder} from "../composables/userImportFiles";
</script>

<template>
  <div id="base-view" class="base-view" >
    <section class="empty-section" v-if="dataCenter.preview.length == 0">
      <p class="title">批量重命名文件</p>
      <p class="tips">将文件或文件夹拖拽到这里</p>
      <div class="btn-group">
        <button class="btn btn-file" @click="importFiles">导入文件</button>
        <button class="btn btn-folder" @click="importFolder">导入文件夹</button>
      </div>
      <p class="description">支持添加前/后缀、替换指定字符串、按序号重命名、大小写转换、删除非法字符</p>
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
  text-align: center;
}

.empty-section {
  width: 100%;
  height: 100%;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center;
}

.title {
  font-size: 1.5rem;
  font-weight: bolder;
  margin: 1rem;
}
.tips {
  font-size: 1.2rem;
  font-weight: bold;
  margin: 1rem;
}
.description {
  font-size: 0.8rem;
  font-weight: lighter;
  margin: 1rem;
}

.btn {
  margin: 5px 10px;
  width: 100px;
  height: 40px;
  color: rgb(83,67,71);
  font-size: 1rem;
  border: none;
  border-radius: 10px;
}

.btn:hover {
  background-color: rgb(83,67,71);
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

