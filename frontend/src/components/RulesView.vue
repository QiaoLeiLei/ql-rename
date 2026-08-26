<script setup lang="ts">
function onRuleChange(rule: number) {
  console.log("[js:当前规则]:",rule)
}
</script>

<template>
  <div id="rules-view" class="rules-view">
    <p class="title">规则</p>
    <div id="radioBox" class="radioBox">
      <input type="radio" name="rules" id="tab1" value="1" @change="onRuleChange(1)" checked/>
      <label class="radio-label" for="tab1">添加前缀</label>
      <input type="radio" name="rules" id="tab2" value="2" @change="onRuleChange(2)"/>
      <label class="radio-label" for="tab2">添加后缀</label>
      <input type="radio" name="rules" id="tab3" value="3" @change="onRuleChange(3)"/>
      <label class="radio-label" for="tab3">替换文本</label>
      <input type="radio" name="rules" id="tab4" value="4" @change="onRuleChange(4)"/>
      <label class="radio-label" for="tab4">顺序编号</label>
      <input type="radio" name="rules" id="tab5" value="5" @change="onRuleChange(5)"/>
      <label class="radio-label" for="tab5">改大小写</label>
      <input type="radio" name="rules" id="tab6" value="6" @change="onRuleChange(6)"/>
      <label class="radio-label" for="tab6">非法字符</label>
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

.title{
  font-size: 1rem;
  padding: 0 0 15px 12px;
  background-color: rgb(35, 41, 57);
  border-bottom: 1px solid #515151;
  text-align: left;
  font-weight: bold;
}

input[type="radio"]{
  appearance: none;
  margin-right: 5px;
}

.radio-label{
  padding: 5px 5px 5px 5px;
  font-size: 0.8rem;
  background-color: rgb(69, 76, 87);
  border: 1px solid rgb(57, 57, 57);
  border-radius: 5px;
}

input[type="radio"]:checked + .radio-label{
  background-color: #2b85e4;
}
</style>