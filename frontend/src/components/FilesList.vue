<script setup lang="ts">
import {importFiles, importFolder} from "../composables/userImportFiles";
import {dataCenter} from "../dataCenter";
import {ref} from "vue";

function selectFile(index: number): void {
  console.log("[js:selectFile]:", index);
  dataCenter.preview[index].Selected = !dataCenter.preview[index].Selected;
  selectedCount.value = dataCenter.preview.filter(previewData => previewData.Selected).length;
}

const selectedCount = ref(0);
</script>

<template>
  <div id="files-list" class="files-list">
    <div class="title">
      <p>文件列表</p>
      <div class="btn-title">
        <button class="btn btn-file-title" @click="importFiles">导入文件</button>
        <button class="btn btn-folder-title" @click="importFolder">导入文件夹</button>
      </div>
    </div>

    <section class="list-section">
      <div class="file-item" v-for="(previewData,index) in dataCenter.preview" :class="{ 'selected': previewData.Selected }">
        <input v-bind:id="'file-checkbox-' + index" type="checkbox" :checked="previewData.Selected" @change="selectFile(index)">
        <label v-bind:for="'file-checkbox-' + index">{{ previewData.OldDisPlayName }}</label>
      </div>
    </section>

    <div class="files-statistics">
      <p class="selected-count"><img id="notice-icon" alt="notice icon" src="../assets/images/notice-white.svg"/>已选
        {{ selectedCount }} 个文件</p>
    </div>
  </div>
</template>

<style scoped>
#files-list {
  width: 26%;
  height: calc(100% - 18px);
  background-color: rgb(36, 44, 61);
  border-radius: 15px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
  margin-left: 6px;
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
  font-size: 1rem;
  padding: 0 10px;
  background-color: rgb(35, 41, 57);
  border-bottom: 1px solid #515151;
}

.btn {
  margin: 5px;
  width: 75px;
  height: 25px;
  border: none;
  border-radius: 5px;
  font-size: 0.75rem;
  background-color: #2b85e4;
  color: white;
  transition: background-color 0.2s ease-in-out;
}

.btn:hover {
  background-color: rgb(30, 100, 170);
  color: #fff;
}

.list-section {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
}

.file-item {
  width: 100%;
  padding: 2px 0 5px 0;
  margin: 2px;
  border-radius: 5px;
  transition: background-color 0.2s ease-in-out;
  text-align: left;
}

.file-item.selected {
  background-color: rgba(43, 133, 228, 0.25);
  font-weight: bold;
}

input[type="checkbox"] {
  appearance: none;
  width: 16px;
  height: 16px;
  border: 1px solid #2b85e4;
  border-radius: 3px;
  margin:0 5px;
  vertical-align: middle;
}

input[type="checkbox"]:checked {
  background-color: rgb(30, 100, 170);
  position: relative;
}

input[type="checkbox"]:checked::after {
  content: "✓";
  color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  font-size: 1rem;
}

.files-statistics {
  text-align: left;
  background-color: rgb(35, 41, 57);
  border-top: 1px solid #515151;
}

#notice-icon {
  width: 25px;
  height: 25px;
  vertical-align: middle;
  margin-right: 5px;
}

.selected-count {
  font-size: 0.8rem;
  margin: 10px 10px;
  padding: 2px 2px;
}
</style>

