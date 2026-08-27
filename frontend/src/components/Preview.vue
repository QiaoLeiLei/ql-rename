<script setup lang="ts">
import {dataCenter} from "../dataCenter";
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
      <div v-for="(preview, index) in dataCenter.preview" :key="index" class="preview-item">
        <div class="item">
          <p>{{ preview.OldDisPlayName }}</p>
          <p>-></p>
          <p>{{ preview.NewDisPlayName }}</p>
        </div>

      </div>
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
  margin-top:-0.5rem;
  width: 100%;
}
.preview-item{
  height: 1.5rem;
  width: 100%;
}
.item{
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;
  overflow: hidden;
}
</style>