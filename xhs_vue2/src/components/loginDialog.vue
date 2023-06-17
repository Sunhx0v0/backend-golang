<!-- 登陆注册 -->
<template>
<el-dialog
    title="登录"
    width="40%"
    :before-close="handleClose"
    :visible.sync="isVisible"
    @submit.prevent="submitlogin">
    <el-form ref="ruleForm" 
        label-position="left"
        label-width="80px">
                
    <el-form-item label="手机号" 
        prop="phone">
        <el-col>
    <el-input v-model="phone"
        placeholder="请输入手机号"
        clearable></el-input></el-col>
    </el-form-item>
    <el-form-item label="密码" prop="pass">
    <el-input type="password" v-model="password" autocomplete="off"></el-input>
    </el-form-item>
    <el-form-item>
    <el-button type="primary" @click="$emit('input',false),submitLogin('login')">登 录</el-button>
    <el-button @click="regVisible = true">新用户注册</el-button>
    <registerDialog :Visible="regVisible" @updateregVisible="updateregVisible"/>
    </el-form-item>
    </el-form>
    <div v-if="error">{{ error }}</div>
</el-dialog>

</template>

<script>
import registerDialog from "./registerDialog.vue"
export default {
    components:{
        registerDialog
    },
    data(){
        var phoneReg = /^[1][3,4,5,7,8][0-9]{9}$/
        var validatePhone = (rule, value, callback) => {
            if (!value) {
            return callback(new Error('号码不能为空!!'))
            }
            setTimeout(() => {
            if (!phoneReg.test(value)) {
                callback(new Error('格式有误'))
            } else {
                callback()
            }
            }, 100)
        };
        return{
            ruleForm: {
                phone: [
                    {
                    required: true, // required就是这个框必须填写
                    validator: validatePhone, // 规则
                    trigger: 'blur', // blur失去焦点，事件何时触发
                    },
                ],
            },
            phone:'',
            password:'',
            error:'',
            regVisible:false
        }
    },
    name: 'loginDialog',
    props:{
        Visible:{
            type:Boolean,
        }
    },
    computed:{
        isVisible:{
            get(){
                return this.Visible
            },
            set(val){
                this.$emit('updateVisible',val)
            }
        }
    },
    methods:{
        // 提交登录
      async submitlogin () {
      try {
        await this.$store.dispatch('login', {
          phone: this.phone,
          password: this.password
        })
        this.$router.push('/welcome')
      } catch (error) {
        this.error = error.response.data.message
      }
    },
    handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      },
      updateregVisible(){
        this.regVisible=!this.regVisible
        console.log(this.regVisible)
      }
    }
}
</script>