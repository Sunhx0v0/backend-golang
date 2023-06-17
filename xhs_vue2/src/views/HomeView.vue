<template>
  <div id="home">
    <!--放小红书logo-->
    <el-container>
  
    <el-aside style="width:180px;">
      <!-- 设置当前菜单状态跟路由保持一致 -->
      <el-menu :default-active="$route.path">
        <el-menu-item :class="index == number ? 'btn1':'btn'" @click="tab(index)" v-for="(item ,index) in dataList"
                :key="index">{{item.option}}
        </el-menu-item>
      </el-menu>
    </el-aside>

  <el-container>
  <el-header>

  <div>
      <!--elementui有24分栏，span等于几就占据了几分栏，offset为偏移量，el-col里面得有div分隔才生效-->
      <el-row>

      <el-col :span="2" :offset="0" class="left">
        <div><img src="./xhs_logo.png" width="120" height="40" alt="小红书" /></div>
      </el-col>

       <!--设置搜索栏,组件插入，减少本页代码量-->
      <el-col :span="8" :offset="2"  class="center">
        <div> <SearchBox></SearchBox></div>
      </el-col>
     
    <el-col :span="3" :offset="0" class="left">
    <el-button round icon="el-icon-s-home" size="medium"  @click="finded"  :plain= "find"  type="danger">  发现  </el-button>
     </el-col>

     <el-col :span="3" :offset="0" class="left">
     <el-button round icon="el-icon-circle-plus-outline" size="medium"  @click="post" :plain= "posted"  type="danger">  发布  </el-button>
     </el-col>

     <el-col :span="3" :offset="0" class="left">
     <el-button round icon="el-icon-tableware" size="medium"  @click="mine" :plain= "my"  type="danger">  我的  </el-button>
     </el-col>

     <!-- 弹窗 -->
     <el-col :span="3" :offset="0" class="left">
      <!-- 点击button进入外层 -->
      <el-button size="medium" @click="outerVisible = true" :plain= "my"  type="danger">  登录 </el-button>
      <el-dialog title="外层 Dialog" :visible.sync="outerVisible">
      <el-dialog width="30%" title="内层 Dialog" :visible.sync="innerVisible" append-to-body>
        <!-- 内部dialog的东西放这里 -->
      </el-dialog>
      <!-- 外部dialog底部按钮以上的东西放这里 -->
    <div slot="footer" class="dialog-footer">
      <el-button @click="outerVisible = false">取 消</el-button>
      <el-button type="primary" @click="innerVisible = true">打开内层 Dialog</el-button>
    </div>
  </el-dialog> 
    </el-col>

    </el-row>   
    </div>
  </el-header>

  <el-main class="tab">
  <!-- 11个div可以进行切换，number=0是默认第一个div -->

  <div v-show="number == 0">
  <!-- 下面的模块是走马灯 -->
  <div>
  <el-carousel :interval="4000" type="card" height="300px">
    <el-carousel-item v-for="item in imgList" :key="item.id">
      <img :src="item.url" class="image">
    </el-carousel-item>
  </el-carousel>
  </div>

  <!-- 下面的模块是图片三列排布 -->
  <div style="display: flex" class="average_container">
      <div v-for="(column, index) in columns" :key="index" >
        <div v-for="image in column" :key="image.id">
          <img :src="image.url" >
        </div>
      </div>
  </div>


  </div>

  <div v-show="number == 1">
        <span>C语言：C语言是一门面向过程的、抽象化的通用程序设计语言，广泛应用于底层开发。</span>
  </div>

  <div v-show="number == 2">
        <span>JavaScript：是一种具有函数优先的轻量级，解释型或即时编译型的编程语言。</span>
  </div>

  <div v-show="number == 3">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 4">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 5">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 6">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 7">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 8">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 9">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  <div v-show="number == 10">
        <span>PHP：是在服务器端执行的脚本语言，尤其适用于Web开发并可嵌入HTML中。</span>
  </div>
  </el-main>
  
  </el-container>
  </el-container>
  </div>

</template>

<script>
import axios from "axios";
import SearchBox from '@/components/SearchBox.vue'

export default {
  name: 'HomeView',
  components:{
          SearchBox
      },
  data() {
    return {
      number:0,//设置菜单的默认选中值
      find:false,
      posted:true,
      my:true,
      outerVisible: false,
      innerVisible: false,
      // dialogVisible:true,
      dataList: [
          {option: '推荐'},
          {option: '我的关注'},
          {option: '美食'},
          {option: '彩妆'},
          {option: '穿搭'},
          {option: '影视'},
          {option: '情感'},
          {option: '职场'},
          {option: '家居'},
          {option: '游戏'},
          {option: '旅行'},
        ],
        imgList: [
                {id:0,url:require('./beauty.jpg')},
                {id:1,url:require('./2.jpg')},
                {id:2,url:require('./3.jpg')},
                {id:3,url:require('./4.jpg')}            
            ],
        screenWidth :0,
        //图片和列
        images: [],
        columns: [[], [], []],
    };
    
  },

  methods: {
  //下面是走马灯图片自适应的第一个函数，通过图片宽度计算高度
  setSize:function () {
                // 通过浏览器宽度(图片宽度)计算高度
                this.bannerHeight = 400 / 1920 * this.screenWidth;
              },
    
    tab(index) {
        this.number = index;
        // console.log(index, this.number);
      },
    // 即可实现点击按钮的跳转，对应路由里面的path
    finded(){
      this.find=false//背景为透明色
      this.posted=true
      this.my=true
      // this.$router.push('/about');可实现跳转页面
    },
    post() {
      this.find=true//背景为透明色
      this.posted=false
      this.my=true
    },
    mine() {
      this.find=true
      this.posted=true
      this.my=false
     },  
     //images先定义为向本地请求的照片的名称，即阿拉伯数字1-10
    //  fetchImages() {
    //   this.images=[ 
    //             {id:0,url:require('./beauty.jpg')},
    //             {id:1,url:require('./2.jpg')},
    //             {id:2,url:require('./3.jpg')},
    //             {id:3,url:require('./4.jpg')} ,
    //             {id:4,url:require('./beauty.jpg')}
    //           ];
    //   this.splitImagesIntoColumns()
    //   // 向后端发送 GET 请求
    //   axios.get('/user/userinfo')
    //     .then(res => {
    //       console.log(res);
    //       this.images=[ 
    //             {id:0,url:require('./beauty.jpg')},
    //             {id:1,url:require('./2.jpg')},
    //             {id:2,url:require('./3.jpg')},
    //             {id:3,url:require('./4.jpg')} ,
    //             {id:4,url:require('./beauty.jpg')}
    //           ];
    //       this.splitImagesIntoColumns()
    //     })
    //     .catch(error => {
    //       console.log(error)
    //     })
    // },
    splitImagesIntoColumns() {

          this.images=[ 
                {id:0,url:require('./beauty.jpg')},
                {id:1,url:require('./2.jpg')},
                {id:2,url:require('./3.jpg')},
                {id:3,url:require('./4.jpg')} ,
                {id:4,url:require('./beauty.jpg')}
              ];
      let columnCount = 3
      let imagesPerColumn = Math.ceil(this.images.length / columnCount)
      for (let i = 0; i < columnCount; i++) {
        for (let j = 0; j < imagesPerColumn; j++) {
          let image = this.images[i * imagesPerColumn + j]
          if (image) {
            this.columns[i].push(image)
          }
        }
      }
    }, 
  },
  
  //这个是走马灯的第二个图片自适应的函数
  mounted() {
              // 首次加载时,需要调用一次
              this.screenWidth =  window.innerWidth;
              this.setSize();
              // 窗口大小发生改变时,调用一次
              window.onresize = () =>{
              this.screenWidth =  window.innerWidth;
              this.setSize();
              //首次加载时去请求所有的笔记
              //this.fetchImages();
              //this.splitImagesIntoColumns()
            }
          },
          created:function(){
            this.splitImagesIntoColumns()
        }
}
</script>

<style>

.el-container {
  height: 100%;
}
.el-header{
  position: absolute;
    line-height: 50px;
    top: 0px;
    left: 0px;
    right: 0px;
    background-color: #3881ce01;
    margin-top: 0px;
}

.el-footer {
  background-color: #bdbaba0e;
}

.el-aside {
  position: absolute;
    width: 200px;
    top: 50px;  /* 距离上面50像素 */
    left: 0px;
    bottom: 0px;
    overflow-y: auto; /* 当内容过多时y轴出现滚动条 */
    background-color: #def3f927;

}

.el-main {
  position: absolute;
    top: 50px;
    left: 200px;
    bottom: 0px;
    right: 0px;  /* 距离右边0像素 */
    padding: 10px;
    overflow-y: auto; /* 当内容过多时y轴出现滚动条 */
    /* background-color: red; */
}

.el-menu {
  background-color: #e7486bcc;
}

/*选中时*/
.btn1 {
  background-color:#fff !important;
      color: #ec2929  !important;
      /* 设置选中时的字体颜色和背景颜色 */
 }

 /* 走马灯的style设置 */
 .el-carousel__item h3 {
      color: #475669;
      font-size: 14px;
      opacity: 0.75;
      line-height: 200px;
      margin: 0;
    }

    .el-carousel__item:nth-child(2n) {
      background-color: #99a9bf;
    }

    .el-carousel__item:nth-child(2n+1) {
      background-color: #d3dce6;
    }
    img{
     /*设置图片宽度和浏览器宽度一致*/
      width:100%;
      height:inherit;
    }
    /* 实现布局三列平分 */
    .average_container>div{
    flex: 1;
    }
    .average_container div {  
            border: solid 1px rgba(237, 226, 226, 0.174); 
            margin-right: calc(80px / 3);
            box-sizing: border-box;
            margin-top:20px;
            margin-bottom:20px;
    }

</style>
