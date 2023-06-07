<!-- 当URL正好对应这个组件的时候，它就会被渲染到App.vue中的<router-view>标签里面。 -->

<template>
    <div id = "PublishPage">

      <!-- 加个logo -->
<!-- 用成Tabs 标签页更好 -->
      <el-button class = 'button' round  icon='el-icon-picture-outline' autofocus="true" @click="photo">上传图文</el-button>
      <el-button class = 'button' round  icon='el-icon-video-camera' @click="video">上传视频</el-button>

      <br><br>
      <el-upload
        class="upload-demo"
        drag
        action="#"
        :http-request="uploadImage"
        :before-upload="beforeUpload"
        multiple>
        <i class="el-icon-upload"></i>
          <div class="el-upload__text">将图片拖拽到此处，或<em>点击上传</em></div>
        <div class="el-upload__tip" slot="tip">只能上传jpg/png文件，且不超过2M</div>
      </el-upload>

    </div>
</template>


<script>
// export function uploadFile(params){
//   return this.$axios({
//     url:"/url/url/url",
//     method:'POST',
//     data:params
//   })
// }
export const uploadFile = (params) => this.$axios.post("/url/url/url", params)
export default {
  name: 'PublishPage',
  props: {
    msg: String
  },
  data(){
    return {
    }
  },
  methods:{
    photo(){
      this.$router.push('/videoPublish')
    },
    video(){

    },
    async uploadImage(params){
    // 一般上传文件使用的是formdata格式 我们需要使用formdata进行传参
      const formdata = new FormData()
      formdata.append("file",params.file)
      const res = await uploadFile(formdata)
      if (res.success) {
        this.$Message.success('This is a success tip');
      }
    },
//     UploadFiles: function () { 
//    var f = this.$ref.photo;
//    console.dir(f.files[0]);
//    let param = new FormData();
//    param.append('img', f.files[0]);//通过append向form对象添加数据
//    console.log(param.get('img')); //FormData私有类对象，访问不到，可以通过get判断值是否传进去
//    let config = {
//      headers: { 'Content-Type': 'multipart/form-data' }
//    };  //添加请求头
//    this.$axios.post(this.$store.state.https + this.$store.state.geturl.upload + '?token=' + this.$store.state.tokenstatic, param, config)
//    .then(response => {
//       this.form.url = this.$store.state.geturl.photourl + response.data.data.url;
//          console.log(this.form.url)
//          if (this.form.url) {
//            alert('图片上传成功')
//          }
//   })
// },




    // uploadImage (param) {
    //   const formData = new FormData()
    //   formData.append('ModelName', param.file) // 要提交给后台的文件
    //   formData.append('projectFolder', this.routeProjectId) // 这个接口必要的项目id
    //   formData.append('subFolder', this.routeFloorId) // 这个接口必要的其他的id
    //   UploadFiles(formData).then(response => {  // UploadFiles 是封装的接口
    //     if (response !== undefined) {
    //       Message({
    //         message: response.Msg,
    //         type: 'success',
    //         duration: 5 * 1000
    //       })
    //       param.onSuccess()
    //       this.$store.state.basics.ShuxinTable = true
    //     }
    //   })
    // },

    beforeUpload(file){
      // 只能上传jpg和png的图片
      const isImage = file.type==="image/jpg"||file.type==="image/png"
      // 大小限制为2M    
      const isLt2M = file.size/1024/1024 <= 2
 
      if(!isImage){
        this.$message.error("图片格式错误！")
      }
      if(!isLt2M ){
        this.$message.error("图片大小错误！")
      }
      return isImage && isLt2M 
    }

    

  }
}
</script>



<style scoped>
.button 
{
  border-style:none;
  font-weight:bold;
  font-size:20px;
  width:300px;
  height: 50px;
}

</style>

