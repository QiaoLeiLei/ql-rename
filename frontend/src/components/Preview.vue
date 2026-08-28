<script setup lang="ts">
import {dataCenter} from "../dataCenter";
import { ExecsRename } from "../../wailsjs/go/main/App";
function confirmRename(): void {
  console.log("[js:confirmRename]: 确认重命名");
  ExecsRename().catch((err: Error) => {
    console.error("[js:confirmRename]: 重命名失败", err);
  });
}
</script>

<template>
  <div id="preview" class="preview" >
    <div class="preview-title">
      <h2>预览</h2>
      <hr style="width: 100%; color: #3e3e3e;">
      <div class="preview-description">
        <div>原始名</div>
        <div>→</div>
        <div>新名称</div>
      </div>
    </div>

    <section class="preview-content">
      <div v-for="(preview, index) in dataCenter.preview" :key="index" class="preview-item" :class="{'is-selected': preview.Selected}">
        <div class="item">
          <p>{{ preview.OldDisPlayName }}</p>
          <p>-></p>
          <p>{{ preview.NewDisPlayName }}</p>
        </div>
      </div>
    </section>

    <div class="bottom">
      <button class="btn btn-confirm" @click="confirmRename">确认重命名</button>
    </div>
  </div>
</template>

<style scoped>
.preview {
  width: 32%;
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
  flex-direction: column;
  align-items: flex-start;
  font-weight: bold;
  font-size: 1rem;
  background-color: rgb(35, 41, 57);
  border-bottom: 1px solid #676767;
  text-align: left;
}

h2{
  font-size: 1rem;
  padding-left: 1rem;
  margin-bottom: 0.5rem;
}

.preview-description{
  width: 90%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin:0 auto 0.5rem;
}

.preview-content{
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  width: 100%;
  overflow: hidden;
}

.preview-item{
  height: auto;
  width: 100%;
  margin: 2px 0;
  align-items: center;
}

.preview-item.is-selected{
  background-color: rgba(43, 133, 228, 0.25);
  font-weight: bold;
}

.item{
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  overflow: hidden;
  font-weight: lighter;
  font-size: 0.8rem;
  margin: -0.5rem auto;
}

.bottom {
  height: 3rem;
  text-align: right;
  background-color: rgb(35, 41, 57);
  border-top: 1px solid #515151;
}

.btn-confirm {
  margin: 5px;
  width: 100px;
  height: 35px;
  border: none;
  border-radius: 5px;
  font-size: 0.9rem;
  background-color: #2b85e4;
  color: white;
  transition: background-color 0.2s ease-in-out;
}

.btn-confirm:hover {
  background-color: rgb(30, 100, 170);
  color: #fff;
}
</style>