<template>
    <div>
        <h1>创建分类</h1>
        <el-form label-width="80px" style="margin-top:20px;" @submit.native.prevent="save">
            <!-- submit事件中使用提交方法save，native是找到js原生表单，prevent是指阻止提交时默认跳转页面 -->
            <el-form-item label="（分类1）/账号">
                <el-input v-model="model.username"></el-input>
                <!-- 双向绑定model -->
            </el-form-item>
            <el-form-item label="（分类2）/密码">
                <el-input v-model="model.password"></el-input>
                <!-- 双向绑定model -->
            </el-form-item>
            <el-form-item>
                <el-button type="primary" native-type="submit">保存</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
export default {
  data () {
    return {
      model: {}
    }
  },
  methods: {
    // save () {
    //   // 这里放axios方法，用来提交数据
    //   // 自己定义的那个属性叫$AXIOS
    //   // this.$AXIOS.post()
    // }
    // 使用async..await异步请求方法
    async save () {
      // 这里放axios方法，用来提交数据

    // 接口调用之前写好的后台find_all查询接口进行测试接口连接
      // const res = await this.$AXIOS.post('/find_all')
      const res = await this.$AXIOS.post('/insert', this.model)
      console.log(res)
    }

  }
}
</script>
