<script setup lang="ts">
import {ref} from "vue";
import {dataCenter} from "../dataCenter";

const currentRule = ref(1);

function onRuleChange(rule: number) {
  console.log("[js:当前规则]:", rule)
  currentRule.value = rule;
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
      <label class="radio-label" for="tab3">查找替换</label>
      <input type="radio" name="rules" id="tab4" value="4" @change="onRuleChange(4)"/>
      <label class="radio-label" for="tab4">顺序编号</label>
      <input type="radio" name="rules" id="tab5" value="5" @change="onRuleChange(5)"/>
      <label class="radio-label" for="tab5">改大小写</label>
      <input type="radio" name="rules" id="tab6" value="6" @change="onRuleChange(6)"/>
      <label class="radio-label" for="tab6">非法字符</label>
    </div>
    <div id="all-content">
      <section id="tab1-content" class="tab-content" v-if="currentRule === 1">
        <p>给原文件名添加前缀</p>
        <label for="prefix">前缀: </label>
        <input type="text" id="prefix" v-model="dataCenter.rules.Prefix" placeholder="test_" required>
        <div class="description">
          <p>例如：<br><br>原文件名： "123.txt"<br>输入前缀： "test_"<br>重命名结果为： "test_123.txt"</p>
        </div>
      </section>
      <section id="tab2-content" class="tab-content" v-if="currentRule === 2">
        <p>给原文件名添加后缀</p>
        <label for="suffix">后缀: </label>
        <input type="text" id="suffix" v-model="dataCenter.rules.Suffix" placeholder="test_" required>
        <div class="description">
          <p>例如：<br><br>原文件名： "123.txt"<br>输入后缀： "_test"<br>重命名结果为： "123_test.txt"</p>
        </div>
      </section>
      <section id="tab3-content" class="tab-content" v-if="currentRule === 3">
        <p>查找替换指定字符串</p>
        <div>
          <label for="oldStr">查找字符串: </label>
          <input type="text" id="oldStr" v-model="dataCenter.rules.ReplaceObj!.OldStr" placeholder="123" required>
        </div>
        <div>
          <label for="newStr">替换字符串: </label>
          <input type="text" id="newStr" v-model="dataCenter.rules.ReplaceObj!.NewStr" placeholder="456" required>
        </div>
       <div class="description">
          <p>例如：<br><br>原文件名： "123.txt"<br>查找字符串： "123"<br>替换字符串： "456"<br>重命名结果为： "456.txt"</p>
        </div>
      </section>
      <section id="tab4-content" class="tab-content" v-if="currentRule === 4">
        <p>安顺序编号</p>
      </section>
      <section id="tab5-content" class="tab-content" v-if="currentRule === 5">
        <p>将文件名转换为大小写/小写</p>
        <select v-model="dataCenter.rules.ToUpperCase">
          <option value="0" selected>转小写</option>
          <option value="1">转大写</option>
        </select>
        <div class="description" v-if="dataCenter.rules.ToUpperCase">
          <p>例如：<br><br>原文件名： "abc.txt"<br>转换为大写： "ABC.txt"<br>重命名结果为： "ABC.txt"</p>
        </div>
        <div class="description" v-else>
          <p>例如：<br><br>原文件名： "ABC.txt"<br>转换为小写： "abc.txt"<br>重命名结果为： "abc.txt"</p>
        </div>
      </section>
      <section id="tab6-content" class="tab-content" v-if="currentRule === 6">
        <p>删除文件名中的非法字符</p>
        <div class="description">
          <p class="tips"><span style="color: lightcoral;"> * </span><em>只保留文件名中的汉字、字母、数字、下划线、连字符</em></p>
          <br>
          <p>例如：<br><br>原文件名： "a_ b-@&c记录.txt"<br>重命名结果为： "a_b-c记录.txt"</p>
        </div>
      </section>
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

.title {
  font-size: 1rem;
  padding: 0 0 15px 12px;
  background-color: rgb(35, 41, 57);
  border-bottom: 1px solid #515151;
  text-align: left;
  font-weight: bold;
}

input[type="radio"] {
  appearance: none;
  margin-right: 5px;
}

.radio-label {
  padding: 5px 5px 5px 5px;
  font-size: 0.8rem;
  background-color: rgb(69, 76, 87);
  border: 1px solid rgb(57, 57, 57);
  border-radius: 5px;
}

input[type="radio"]:checked + .radio-label {
  background-color: #2b85e4;
}

.tab-content {
  padding: 20px;
  margin-top: 20px;
}

.description {
  text-align: left;
  padding: 30px 30px;
}

input[type="text"] {
  background-color: rgb(36, 44, 61);
  border: none;
  color: #fff;
  border-radius: 2px;
  padding: 5px 5px;
  font-size: 0.8rem;
  font-weight: bold;
  outline: 2px solid #2b85e4;
  margin: 10px;
}

select {
  background-color: rgb(36, 44, 61);
  border: none;
  color: #fff;
  height: 30px;
  width: 150px;
  font-size: 0.8rem;
  font-weight: bold;
  outline: 2px solid #2b85e4;
  margin-top: 10px;
  text-align: center;
  text-align-last: center;
}

option {
  text-align: center;
}

.tips {
  font-size: 0.7rem;
}

</style>