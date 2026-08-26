<script setup lang="ts">

</script>

<template>
  <div id="rules-view" class="rules-view">
    <div>规则</div>
    <div id="tabBox" class="tab-box">
      <div class="tab-title" >
        <ul>
          <li class="">前缀</li>
          <li class="">后缀</li>
          <li class="">替换</li>
          <li class="">删除</li>
          <li class="">顺序</li>
          <li class="active">大小写</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>
.rules-view {
  width: 43%;
  height: calc(100% - 18px);
  background-color: rgb(36, 44, 61);
  border-radius: 15px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.5);
}
</style>