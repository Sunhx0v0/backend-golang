<template>
  <div>
    <h1>{{ msg }}</h1>
    <p>
      保留一部分汉字，才知道你是中国人,<br>
      别点右边→
      <a href="www.bilibili.com" target="_blank" rel="noopener">小绿书</a>.
    </p>
    {{ message }} ←该字段就是从后端获取的字段→
    <p>
      上面应该有一段文字
    </p>
    {{ info }}
  </div>
</template>
  
<script>
import axios from 'axios';       // 需要先npm install 并且import进来
export default {
  data() {
    return {
      message: '123',   // 如果没有成功获取，那么显示的就是默认值（即123）
      info: '不知道为什么传递第二个值就会报错 ',
    };
  },
  created() {  // created方法：首次启动该界面时加载的
    axios.get('http://localhost:8083/api/hello')       // 通过axios方法获取，get/post前后端需要保持一致
      .then(response => this.message = response.data.message) // response是从后端接收到的报文的名字，可以任取；response中有很多字段，其中.data为接收到的数据；.message为数据的名称
      // .then(response => this.info = response.data.info)
  },
};
</script>
  