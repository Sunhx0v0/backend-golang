<template>
    <div class="login-container">
        <!-- 使用elementui中的card组件 -->
        <el-card header="管理员登录" class="login-card">
            <!-- 监听表单的submit事件，native.prevent监听原生表单事件跳转接口并且阻止页面跳转 -->
            <el-form @submit.native.prevent="login">
                <el-form-item label="用户名">
                    <el-input v-model="model.username"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="model.password" type="password"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" native-type="submit">登录</el-button>
                </el-form-item>
            </el-form>
        </el-card>
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
    async login () {
      const res = await this.$AXIOS.post('login', this.model)
      console.log(res)
    }
  }
}
</script>

<style>
.login-card{
    width: 35rem;
    margin: 6rem auto;
}
</style>
