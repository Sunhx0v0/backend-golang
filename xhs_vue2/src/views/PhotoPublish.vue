<template>
    <div class="photoPublish">
        <el-container>
          <el-header>
               <!-- <el-row> -->
                <el-col :span="1" :offset="0">
                  <div><img src="@/assets/xhs_logo.png" width="120" height="40" alt="小红书" /></div>
                </el-col>
              <!-- </el-row> -->
          </el-header>
        <el-container>

        <el-aside width="350px">Aside
        </el-aside>

        <el-main>
            <el-row>
              <el-col :span="3"><div class = "left" style = "font-size:25px;white-space:nowrap;">发布图文<i class="el-icon-picture-outline"></i></div></el-col>
            </el-row><br>

            <el-row>
              <el-col :span="3"><div class = "left" style ="font-size:20px;transform: scale(1, 0.975);">上传图片</div></el-col>
            </el-row><br>

            <!-- <el-row><el-col :span="9">
              <el-upload class="upload-demo" drag action="#" :http-request="uploadImage" :before-upload="beforeUpload" :before-remove="beforeRemove" multiple:limit="2" :on-exceed="handleExceed" :file-list="fileList">
                <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将图片拖拽到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">Tips:只能上传jpg/jpeg/png文件，且不超过2M</div>
              </el-upload></el-col>
            </el-row><br> -->

            <!-- action:上传的地址 -->
            <el-row><el-col :span="9">
            <el-upload class="upload-demo"
            drag
            action="http://localhost:8080/photoPublish"
            :before-upload="beforeUpload"
            :on-preview="handlePreview"	
            :on-remove="handleRemove"		
            :file-list="fileList"
            :list-type="picture"		
            :on-success="handleSuccess"
            multiple
            :limit="2" 
            :on-exceed="handleExceed"	
            >
            <i class="el-icon-upload"></i>
                  <div class="el-upload__text">将图片拖拽到此处，或<em>点击上传</em></div>
                  <div class="el-upload__tip" slot="tip">Tips:只能上传jpg/jpeg/png文件，且不超过2M</div>
            <!-- <el-button size="small" type="primary">点击上传</el-button> -->
            </el-upload></el-col>
            </el-row>

            <el-row>
            <el-dialog
              title="图片预览"
              :visible.sync="previewVisible"
              width="50%"
            >
              <img :src="previewPath" class="previewImg"/>
            </el-dialog>
            </el-row><br>

            <el-row><el-col><div class = "left" style = "width: 500px" >
              <el-input type="text" placeholder="请输入标题" v-model="title" maxlength="20" show-word-limit>
              </el-input></div></el-col>
            </el-row><br>
        
            <TextArea/>
            <br><br><br><br>

            <el-row>
              <el-col :span="3"><div class = "left" style ="font-size:20px;transform: scale(1, 0.975);">发布设置</div></el-col>
            </el-row><br>

            <el-form ref="form" :model="form" label-width="85px" label-position=left>
              <el-form-item label="自定义地点">
                <el-input v-model="form.location" placeholder="请输入自定义地点" maxlength="20" show-word-limit></el-input>
              </el-form-item>
              <el-form-item label="所属分栏">
                <el-checkbox-group v-model="form.classi">
                  <el-checkbox label="美食" name="type"></el-checkbox>
                  <el-checkbox label="彩妆" name="type"></el-checkbox>
                  <el-checkbox label="穿搭" name="type"></el-checkbox>
                  <el-checkbox label="影视" name="type"></el-checkbox>
                  <el-checkbox label="职场" name="type"></el-checkbox>
                  <el-checkbox label="情感" name="type"></el-checkbox>
                  <el-checkbox label="家居" name="type"></el-checkbox>
                  <el-checkbox label="游戏" name="type"></el-checkbox>
                  <el-checkbox label="旅游" name="type"></el-checkbox>
                  <el-checkbox label="健身" name="type"></el-checkbox>
                  <el-checkbox label="其他" name="type"></el-checkbox>
                </el-checkbox-group>
              </el-form-item>
            </el-form>

            <el-row>
              <el-button round size="medium" @click="submit()">发布</el-button>
            </el-row><br>

            <DetailsCom/>


        </el-main>
        </el-container>
      </el-container>
    </div>
  </template>

  
  <script>  
  // import axios from 'axios';
  import TextArea from '@/components/TextArea.vue'
  import DetailsCom from '@/components/DetailsCom.vue'
  export default {
    name: 'PhotoPublish',
    data() {
      return {
        fileList:[],
        //控制图片预览窗口的显示与隐藏
        previewVisible: false,
        //图片地址
        previewPath: '',
        title: '',
        textarea :'',
        form:{
          location:'',
          classi:[],
        }
      }
    },
    components: {
      TextArea,
      DetailsCom,
    },
    methods: {
      submit()
      {
        console.log(this.title)
        console.log(this.fileList)
        console.log(this.form.location)
        console.log(this.form.classi)
      },

      beforeUpload(file){
      // 只能上传jpg(jpeg)和png的图片
      var isImage = file.type==="image/jpg"||file.type==="image/png"||file.type==="image/jpeg"
      // 大小限制为2M    
      var isLt2M = file.size/1024/1024 <= 2
      if(!isImage){
        this.$message.error("图片格式错误！")
      }
      if(!isLt2M ){
        this.$message.error("图片大小错误！")
      }
      return isImage && isLt2M 
      },

      handleExceed(files, fileList) {
        console.log("124335")
        this.$message.warning(`当前限制选择 2 个文件，本次选择了 ${files.length} 个文件，共选择了 ${files.length + fileList.length} 个文件`);
      },
    
      beforeRemove(file) {
        return this.$confirm(`确定移除 ${ file.name }？`);
      },

    // async uploadImage(params){
    // // 一般上传文件使用的是formdata格式 我们需要使用formdata进行传参
    //   const formdata = new FormData()
    //   formdata.append("file",params.file)
    //   //const res = await uploadFile(formdata)
    //   //if (res.success) {
    //   this.$message.success('图片上传成功！');
    //   //}
    // },
    handlePreview(file){
      this.previewVisible = true;
      this.previewPath = file.response.url;
    },

    //点击删除图片
    handleRemove(file){   //file是将要移除的那个图片的信息,可以console.log(file)，查看file有什么信息
      //1.获取将要删除的图片的临时路径
      const filePath = file.response.data.tmp_path;
      //2.从pics数组中找到图片对应的索引值
      //形参‘x’，是数组pic里的每一项
      const i = this.addForm.pics.findIndex(x =>
        x.pic == filePath);
      //调用数组的splice方法，把图片的信息对象从pics数组中移除
      this.addForm.pics.splice(i,1);
      this.$message.success("图片删除成功！");
    },
    handleSuccess(response){
        //拼接得到图片信息对象
        const picInfo = {pic:response.data.tmp_path};
        //将图片信息对象，push到pics数组中
        this.addForm.pics.push(picInfo);
        //console.log(this.addForm)
      }
  }
}
  </script>



  <style scoped>
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
  .upload-demo{
    position:left;
  }
  .el-upload__tip{
    text-align:center;
  }
  </style>
  