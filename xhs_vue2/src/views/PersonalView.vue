<template>
    <div>
      <el-backtop></el-backtop>
        <div class="personal">
            <el-container>
                <el-header>
                    
                </el-header>
                <el-main>
                      <el-dialog
                          title="资料修改"
                          :visible.sync="dialogVisible"
                          width="40%"
                          :before-close="handleClose">
                          <el-form ref="ruleForm" 
                                :model="form" 
                                label-position="left"
                                label-width="80px">
                            <el-form-item label="头像">
                              <el-row>
                              <el-upload
                                class="avatar-uploader"
                                action="https://jsonplaceholder.typicode.com/posts/"
                                :show-file-list="true"
                                :on-success="handleAvatarSuccess"
                                :before-upload="beforeAvatarUpload">
                                <img v-if="imageUrl" :src="require('../assets/'+this.imageUrl)" class="avatar">
                                <i v-else class="el-icon-plus avatar-uploader-icon"></i>
                            </el-upload>
                              </el-row>
                          </el-form-item>
                          <el-form-item label="昵称">
                            <el-input
                              placeholder="请输入昵称" 
                              v-model="form.name"
                              clearable
                              maxlength="10"
                              show-word-limit
                              >
                            </el-input>
                          </el-form-item>
                          <el-form-item label="性别">
                            <el-radio-group v-model="form.sex">
                              <el-radio label="男"></el-radio>
                              <el-radio label="女"></el-radio>
                            </el-radio-group>
                          </el-form-item>
          
                          <el-form-item label="手机号" 
                              prop="phone">
                              <el-col>
                            <el-input v-model="form.phone"
                                placeholder="请输入手机号"
                                clearable></el-input></el-col>
                          </el-form-item>

                          <el-form-item label="生日">
                            <el-col :span="5">
                              <el-date-picker
                                v-model="form.birth"
                                type="date"
                                placeholder="选择日期"
                                value-format="yyyy-MM-dd"
                                :picker-options="pickerOptions">
                              </el-date-picker>
                            </el-col>
                          </el-form-item>
                          <el-form-item>
                            <el-button @click="dialogVisible = false">取 消</el-button>
                            <el-button type="primary" @click="dialogVisible = false,submitForm('ruleForm')">保 存</el-button>
                          </el-form-item>
                          </el-form>
                      </el-dialog>

                          <div class="block"><el-avatar :size="180" :src="require('../assets/'+this.avatarurl)"></el-avatar></div>
                          <h3 class="user_name">{{form.name}}</h3>
                        <span class="id">小红书号:</span>
                        <span class="id" src="form.id">{{form.id}}</span>
                          <el-row>
                        <span class="sign">{{form.sign}}</span>
                          </el-row>
                        <el-row>
                          <el-button type="primary" icon="el-icon-edit" round class="edit" size="small" @click="dialogVisible = true">资料修改</el-button>
                        </el-row>
                        <el-row :gutter="20">
                        <el-col :span="8" :offset="10"><div>
                          <el-table
                              :data="tableData"
                              style="width:210px"
                              >
                              <el-table-column
                                prop="guanzhu"
                                label="关注"
                                width="60"
                                align="center">
                              </el-table-column>
                              <el-table-column
                                prop="fans"
                                label="粉丝"
                                width="60"
                                align="center">
                              </el-table-column>
                              <el-table-column
                                prop="liked"
                                label="获赞与收藏"
                                width="90"
                                align="center">
                              </el-table-column>
                              </el-table>
                          </div></el-col>
                        </el-row>

                  <el-row type="flex" justify="center">
                      <el-col :span="15" :offset="11">
                  <el-tabs v-model="activeName" @tab-click="handleClick">
                    <el-tab-pane label="笔记" name="first">
                      <el-button type="primary" icon="el-icon-edit" round class="edit" size="small" @click="loginVisible = true">登录</el-button>
                    </el-tab-pane>
                    <el-tab-pane label="收藏" name="second">
                                            
                    </el-tab-pane>
                    <el-tab-pane label="点赞" name="third">
                     
                    </el-tab-pane>
                  </el-tabs>
                </el-col>
                    </el-row> 

                    <loginDialog :Visible="loginVisible" @updateVisible="updateVisible"/>
                  </el-main>
            </el-container>
        </div>
    </div>
</template>

<script>
import axios from 'axios'
import loginDialog from "../components/loginDialog.vue"
  export default {
    // 组件引用
    components:{
      loginDialog,
    },
    data () {
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

      return {
        pickerOptions: {
          disabledDate(time) {
            return time.getTime() > Date.now();
          },
        },
        ruleForm: {
          phone: [
            {
              required: true, // required就是这个框必须填写
              validator: validatePhone, // 规则
              trigger: 'blur', // blur失去焦点，事件何时触发
            },
          ],
      },
        tableData: [{
            guanzhu:'0',
            fans:'0',
            liked:'0'
          }],
        activeName:"first",
        dialogVisible: false,
        // 数据库里只存文件名，全部存到assets里
        imageUrl: 'head1.webp',
        avatarurl:'head1.webp',
        Visible:false,
        loginVisible:false,
        form: {
          name: 'Avenir.',
          sex:'',
          birth:'2003-10-03',
          phone: '12345678',
          id:'111',
          sign:'nothing here'
        },
      }
    },
    // 传个人资料修改数据给后端
    // created(){
    //   axios.post('http://localhost:8805/host',this.form[0])
    //     .then(response=>{
    //     this.form[0].name=response.data.userName;
    //     this.form[0].sex=response.data.gender;
    //     this.form[0].birth=response.data.birthday;
    //     this.form[0].phone=response.data.phoneNumber;
    //     this.form[0].sign=response.data.introduction;
    //     this.imageUrl=response.data.portrait;
    //     this.avatarurl=response.data.portrait;
    //   }),
    //   axios.get('http://localhost:8805/host')
    //     .then(response=>{
    //       this.tableData[0].guanzhu=response.data.guanzhu;
    //       this.tableData[0].fans=response.data.fans;
    //       this.tableData[0].liked=response.data.liked;
    //       this.avatarurl=response.data.avatarurl;
    //       this.form[0].id=response.data.id;
    //       this.form[0].birth=response.data.birth;
    //       this.form[0].name=response.data.name;
    //       this.form[0].sex=response.data.sex;
    //       this.form[0].birth=response.data.birth;
    //       this.form[0].phone=response.data.phone;
    //       this.form[0].sign=response.data.sign;
    //     })
    // },
    methods: {
      handleClick(tab, event) {
        console.log(tab, event);
      },
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      },
      handleAvatarSuccess(res, file) {
        this.imageUrl = file.name;
        console.log(this.imageUrl)
      },
      beforeAvatarUpload(file) {
        const isJPG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 2;

        if (!isJPG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 2MB!');
        }
        return isJPG && isLt2M;
      },
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            const data = {
              // newPassword: this.ruleForm.pass
            }
            this.avatarurl=this.imageUrl
          } else {
            return false
          }
        })
      },
      updateVisible(){
        this.loginVisible=!this.loginVisible
      }
  }
}
</script>

<style>
  .el-header, .el-footer {
    background-color: #FFFFFF;
    color: #303133;
    text-align: center;
    line-height: 60px;
  }
  
  .el-main {
    background-color: #FFFFFF;
    color: #303133;
    text-align: center;
    line-height: 30px;
  }
  
  body > .el-container {
    margin-bottom: 40px;
  }
  .user_name{
    letter-spacing:1px;
    margin-top:0px;
    margin-bottom:1px;
  }
  .id{
    color:#909399;
    letter-spacing:1px;
    margin-top:0px;
    margin-bottom:1px;
    font-weight:500;
    font-size:10px;
  }
  .sign{
    letter-spacing:1px;
    margin-top:0px;
    margin-bottom:1px;
    font-weight:500;
    font-size:13px;
  }
  .input{
    el-form:inline=true;
  }
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 0px;
  }
  .row-bg {
    padding: 10px 0;
  }
  .el-table{
    border:none;
    fit:"true";
    size:"mini";
    text-align:center;
    color:#303133;
    font-weight:400;
    
  }
  .group > .el-table--enable-row-hover .el-table__body tr:hover > td {
  background-color: transparent !important;
}
.inline-block{
  display:inline-block;
  vertical-align: middle;
}
/* dialog的圆角设置但是失败了.note {
  border-radius: 800px;
} */
</style>

<style lang="scss" scoped>
//去掉el-tab-pane底部灰色线条
.el-tabs__nav-wrap::after {
  height: 0 !important;
}
.el-tabs /deep/.el-tabs__nav-wrap::after{
   background-color: #fff;
}
//修改样式
/deep/.el-tabs__item {
  padding: 0 20px 0 0;
}
.edit{
  color: white;
  // background-color:red;
}
.avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 5px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 120px;
    height: 120px;
    line-height: 120px;
    text-align: center;
  }
  .avatar {
    width: 120px;
    height: 120px;
    display: block;
  }
</style>
