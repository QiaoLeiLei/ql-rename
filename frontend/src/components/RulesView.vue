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
  flex: 1;
  width: 30%;
  height: 100%;
  background-color: #9c6423;
}
</style>