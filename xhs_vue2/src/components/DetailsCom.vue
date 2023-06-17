<template>
    <div class ="details">
        <!-- 笔记详情 -->
        <el-button round size="medium" @click="dialogVisible=true;">详情</el-button>

        <!-- <el-container>
          <el-aside width="200px">Aside</el-aside>
          <el-main>Main</el-main>
            </el-container> -->
        <el-dialog
        :visible.sync="dialogVisible"
        style="border-radius: 100px;"
        width="80%">

          <div class = "photoshow">
            <el-carousel trigger="click" height="400px">
              <el-carousel-item v-for="item in imgList" :key="item">
                <img :src="item" id = "toEnlargeImg" class="image" >
                <div class="black_overlay" id="black_overlay"></div>
                <!--预览容器，存放点击放大后的图片-->
                <div class="enlargeContainer" id="enlargeContainer">
                <!-- 关闭按钮 -->
                  <img src="@/assets/close.png" class="close" id="close">
                </div>
              </el-carousel-item>
            </el-carousel>
          </div>

        </el-dialog> 


    
 
    </div>

</template>

<script>  
  window.onload=function(){
    let black_overlay = document.getElementById('black_overlay');
  let enlargeContainer = document.getElementById('enlargeContainer');
  let closeBtn = document.getElementById('close');
  let toEnlargeImg = document.getElementById('toEnlargeImg');
    toEnlargeImg.addEventListener('click', function () {
    // 获取当前图片的路径
    console.log("click")
    let imgUrl = this.src;
    // 显示黑色遮罩和预览容器
    black_overlay.style.display = 'block';
    enlargeContainer.style.display = 'block';
    let img = new Image();
    img.src = imgUrl;
    img.classList.add('enlargePreviewImg');
    if (closeBtn.nextElementSibling) {
      enlargeContainer.removeChild(closeBtn.nextElementSibling);
    }
    enlargeContainer.appendChild(img);
  });

  closeBtn.addEventListener('click', function () {
    black_overlay.style.display = 'none';
    enlargeContainer.style.display = 'none';
  });

}

  export default {
    name: 'DetailsCom',
    data() {
      return {
        dialogVisible:false,
        imgList:['https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg','https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg'],
      }
    },
    
    methods: {
      // viewUploadImages(params).then((response) => {
      //   this.photoLists = response.data
      //   for(let item of this.uploadImages){
      //     this.srcList.push(item.sfUrl)
      //   }
      // })
      getPreviewImgList(index) {
        let arr = []
        let i = 0;
        for(i;i < this.srcList.length;i++){
          arr.push(this.srcList[i+index])
          if(i+index >= this.srcList.length-1){
            index = 0-(i+1);
           }
        }
        return arr;
      },
    }
}
</script>

<style scoped>
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
.photoshow{
  width:50%;
  height:100%;
}
img{
	width: 100%;
	height: 100%;
	object-fit: contain;
}

.black_overlay {
      display: none;
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: rgba(0, 0, 0, 0.7);
      z-index: 100;
    }

    .enlargeContainer {
      display: none;
    }

    .enlargePreviewImg {
      /*这里我设置的是：预览后放大的图片相对于整个页面定位*/
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      
      /*宽度设置为页面宽度的70%，高度自适应*/
      width: 70%;
      z-index: 200;
    }

	/*关闭预览*/
    .close {
      position: absolute;
      top: 20px;
      right: 20px;
      width: 20px;
      height: 20px;
      cursor: pointer;
      z-index: 200;
    }
</style>