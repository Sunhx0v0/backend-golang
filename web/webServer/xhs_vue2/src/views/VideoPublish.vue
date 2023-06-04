<template>
    <div class="videoPublish">
        <!-- <el-row>
            <el-col :span="3" ><div class = "publish">发布图文</div></el-col>
        </el-row> -->

        <!-- <br>
        <el-row>
        <el-col :span="3" ><div class = "edit">图片编辑</div></el-col>
        </el-row> -->

        <!-- <br>
        <el-row>
        <el-col :span="12" >
        <el-input
            placeholder="请输入内容"
            v-model="input"
             clearable>
        </el-input></el-col></el-row> -->

<!-- 开始 -->
        <el-container  style="height: 100%;">
          <el-header>
            <div>
               <el-row>
                <el-col :span="1" :offset="0">
                  <div><img src="@/assets/xhs_logo.png" width="120" height="40" alt="小红书" /></div>
                </el-col>
              </el-row>
            </div>
          </el-header>
        <el-container>

          <el-aside width="350px">Aside



          </el-aside>

          <el-main>
            <el-row>
              <el-col :span="3"><div class = "left" style = "font-size:25px;white-space:nowrap;">发布图文<i class="el-icon-picture-outline"></i></div></el-col>
            </el-row><br>
            <el-row>
              <el-col :span="3"><div class = "left" style ="font-size:20px;transform: scale(1, 0.975);">图片编辑</div></el-col>
            </el-row><br>
            <el-row>
              <el-col><div class = "left"> 
              <el-image style="width: 100px; height: 120px" :src="require(`@/assets/photoo.jpg`)" ></el-image></div></el-col>
            </el-row><br>

            <el-row><el-col><div class = "left" style = "width: 500px" >
              <el-input type="text" placeholder="请输入内容" v-model="text" maxlength="20" show-word-limit>
              </el-input></div></el-col>
            </el-row>
            <!-- 需要引入components的TextArea -->
            <el-row><el-col><div id = "textarea" style="margin: 20px 0;"></div>
              <el-input type="textarea" placeholder="请输入内容" v-model="textarea" maxlength="1000" show-word-limit>
              </el-input></el-col>
            </el-row>
            <TextArea/>





          </el-main>
        </el-container>
      </el-container>
    </div>
  </template>
  
  <script>  
  import TextArea from '@/components/TextArea.vue'
  export default {
    name: 'VideoPublish',
    data() {
      return {
        text: '',
        textarea :''
      }
    },
    components: {
      TextArea
    }
  }
  </script>



  <style>
    .left{
      text-align: left
    }

    .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    /* 高度 */
    line-height: 500px; 
    
  }
  </style>
  