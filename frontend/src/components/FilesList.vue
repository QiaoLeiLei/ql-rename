<script setup lang="ts">
import {importFiles,importFolder} from "../composables/userImportFiles";
import {dataCenter} from "../dataCenter";
</script>

<template>
  <div id="files-list" class="files-list" >
    <div class="title">
      <p>文件列表</p>
      <div class="btn-title">
        <button class="btn btn-file-title" @click="importFiles">导入文件</button>
        <button class="btn btn-folder-title" @click="importFolder">导入文件夹</button>
      </div>
    </div>

    <section class="list-section">
      <div class="file-item" v-for="previewData in dataCenter.preview">
        {{ previewData.OldDisPlayName }}
      </div>
    </section>

    <div class="files-statistics">
      <p class="file-count">共 {{ dataCenter.preview.length }} 个文件</p>
      <p class="selected-count">已选 {{ dataCenter.preview.length }} 个文件</p>
    </div>
  </div>
</template>

<style scoped>
#files-list{
  flex: 1;
  height: 100%;
  display: flex;
  overflow: hidden;
  flex-flow: column nowrap;
  justify-content: space-between;
}
.title {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  font-size: 1.2rem;
  font-weight: bold;
  padding: 3px;
  border: 2px solid lightcyan;
}

.btn {
  margin: 5px;
  width: 80px;
  height:30px;
  border: none;
  border-radius: 10px;
  font-size: 0.75rem;
  font-weight: bold;
  color: rgb(83,67,71);
}

.btn:hover {
  background-color: rgb(83,67,71);
  color: #fff;
}

.list-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
  border: 2px solid lightcyan;
}

.file-item {
  margin: 2px;
  padding: 2px 5px;
}

.files-statistics {
  background-color: #d18a70;
  display: flex;
  flex-flow: row nowrap;
  justify-content: space-between;
  border: 2px solid lightcyan;
}

.file-count {
  margin: 5px;
  background-color: #9c6423;
}
.selected-count {
  margin: 5px;
  background-color: #9c6423;
}
</style>

