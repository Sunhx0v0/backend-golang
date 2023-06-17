<template>
  <div id="textarea">

<el-input
  v-model="textarea"
  class="chatText"
  resize="none"
  type="textarea"
  placeholder="请输入内容"
  id='texttexttext'
  maxlength="1000" 
  show-word-limit
  rows= "5"
  @keydown.native="listen($event)"
  >
</el-input>

<!-- <el-input
  v-model="textarea"
  class="chatText"
  resize="none"
  type="textarea"
  placeholder="请输入内容"
  id='texttexttext'
  maxlength="1000" 
  show-word-limit
  rows= "5"
  @keyup.enter.native="sendInfo"
   @keydown.native="listen($event)"
  >
  <template #suffix>
    <el-button type="primary" style="height: 45px;">获取动态码</el-button>
</template>
</el-input> -->


  <div class="chatIcon">
    <!-- el-popover是弹出框 -->
    <el-popover placement="top-start" width="400" trigger="click" class="emoBox">
        <div class="emotionList">
            <a href="javascript:void(0);" @click="getEmo(index)" v-for="(item,index) in faceList" :key="index" class="emotionItem">{{item}}</a>
        </div>
        <el-button
        class="emotionSelect"
        slot="reference"
        >(◍•ᴗ•◍)表情</el-button>
    </el-popover>
  </div>

        <!-- v-for:数组遍历 -->
  <div class="topic">
    <el-popover placement="top-start" width="200" trigger="click" class="topicBox">
      <div class = "topicList">
          <div href="javascript:void(0);" @click="getTopic(index)" v-for="(item,index) in topicList" :key= "index" class ="topicone">{{item}}</div>
      </div>
      <el-button
        class="emotionSelect"
        slot="reference"
        >#话题</el-button>
    </el-popover>
  </div>

  <div class="others">
    <el-popover placement="top-start" width="200" trigger="click" class="othersBox">
      <div class = "othersList">
          <div href="javascript:void(0);" @click="getothers(index)" v-for="(item,index) in othersList" :key= "index" class="othersone">{{item}}</div>
      </div>
      <el-button
        class="emotionSelect"
        slot="reference"
        >@用户</el-button>
    </el-popover>
  </div>

  </div>
</template>
  
  <script>
  const appData = require(`@/assets/emojis.json`);//require(`@/assets/emojis.json`)
  const topicData = require(`@/assets/topics.json`);
  const othersData = require(`@/assets/others.json`);
  export default {
    name: 'TextArea',

    data() {
    return {
      faceList: [],
      textarea: "",
      topicList:[],
      othersList:[],
      aiteList:[],
      huatiList:[]
    };
    },

    mounted() {//按下表情即执行
      for(let i in appData){
        this.faceList.push(appData[i].char);
      }
      for(let j in topicData){
        this.topicList.push("#"+topicData[j].string);
      }
      for(let h in othersData){
        this.othersList.push(othersData[h].image+othersData[h].string);
      }
    },

    methods: {
      getEmo(index){//按下表情中的某一个才执行
        const textArea=document.getElementById("texttexttext");
        function changeSelectedText(obj, str) {
          if (window.getSelection) {
            // 非IE浏览器
            textArea.focus();
            textArea.setRangeText(str);
            // 在未选中文本的情况下，重新设置光标位置
            textArea.selectionStart += str.length;
            textArea.focus()
          } else if (document.selection) {
            // OE浏览器
            obj.focus();
            var sel = document.selection.createRange();
            sel.text = str;
          }
        }
        changeSelectedText(textArea,this.faceList[index]);
        this.textarea = textArea.value;// 要同步data中的数据
        return;
      },

    getTopic(index){
        const textArea=document.getElementById("texttexttext");
        function changeSelectedText(obj, str) {
          if (window.getSelection) {
            // 非IE浏览器
            textArea.focus();
            textArea.setRangeText(" "+str+" ");
            // 在未选中文本的情况下，重新设置光标位置
            textArea.selectionStart += (str.length+2);
            textArea.focus();
          } else if (document.selection) {
            //OE浏览器
            obj.focus();
            var sel = document.selection.createRange();
            sel.text = str;
          }
        }
        this.huatiList.push(this.topicList[index]);
        changeSelectedText(textArea,this.topicList[index]);

        // var banners = document.getElementById("class");
        // console.log(banners);
        // //创建一个div
        // var div = document.createElement("div");
        // //设置div的id
        // div.id = "slide_show_";
        // //设置div的class
        // div.className = "class1";
        // //设置style
        // div.style.color="red";
        // //设置子div的其他属性，直接通过字符串连接的方式连接即可。动态值放在引号外面就行
        // div.innerHTML = "wenjunhuuuui";
        // banners.appendChild(div);
        // console.log("lail");
        // console.log(banners)

        this.textarea = textArea.value;// 要同步data中的数据
        return;
    },
    getothers(index){
        const textArea=document.getElementById("texttexttext");
        function changeSelectedText(obj, str) {
          if (window.getSelection) {
            // 非IE浏览器
            textArea.focus();
            str = str.substring(2);
            textArea.setRangeText(" @"+str+" ");
            // 在未选中文本的情况下，重新设置光标位置
            textArea.selectionStart += (str.length+3);
            textArea.focus()
          } else if (document.selection) {
            // IE浏览器
            obj.focus();
            var sel = document.selection.createRange();
            sel.text = str;
          }
        }
        this.aiteList.push(this.othersList[index]);
        changeSelectedText(textArea,this.othersList[index]);

        // var res = new RegExp(this.othersList[index],"g");
        //textArea.value = textArea.value.replace(res,'<span style="color:#802eb1">'+this.topicList[index]+'</span>'); 
        this.textarea = textArea.value;// 要同步data中的数据
        return;
    },

    listen(event)
    {
      if(event.which===8) 
      { //删除时，#和@整体删除
        var textArea=document.getElementById("texttexttext");
        var msgLen =textArea.selectionStart;
        var lastword = this.textarea.substring(msgLen);
        for(let j in this.aiteList)
        {
          var loc1=msgLen-this.aiteList[j].length;
          var ailoc = 2;//2是因为表情char占位置了
          var loc2 = loc1;
          if(this.textarea[loc1] =="@")
          {
              loc1++;
             while(this.textarea[loc1] == this.aiteList[j][ailoc]  && ailoc <this.aiteList[j].length)
             {
              loc1++;
              ailoc++;
             }
             if(ailoc == this.aiteList[j].length)
             {
              this.textarea=this.textarea.substring(0,loc2)+lastword+" ";
             }
          }
        }
        for(let k in this.huatiList)
        {
          var loc3=msgLen-this.huatiList[k].length;
          var huatiloc = 1;
          var loc4 = loc3;
          if(this.textarea[loc3] =="#")
          {
             loc3++;
             while(this.textarea[loc3] == this.huatiList[k][huatiloc]  && huatiloc <this.huatiList[k].length)
             {
              loc3++;
              huatiloc++;
             }
            if(huatiloc == this.huatiList[k].length)
            {
              this.textarea=this.textarea.substring(0,loc4)+lastword+" ";
            }
          }
        }
        }
    }
}
  }
  </script>


<style lang="scss">
/* el-popover是和app同级的，所以scoped的局部属性设置了无效 */
/* 需要设置全局style */
  .el-popover{
    height:200px;
    width:400px;
    overflow: scroll;
    overflow-x:auto;
  }
</style>

<style scoped>
.chatIcon {
  padding: 0 10px;
  font-size: 25px;
}
.emotionList{
  display: flex;
  flex-wrap: wrap;
  padding:5px;
}

.topicList{
  display: block;
  flex-wrap: wrap;
}

.emotionSelect{
  float:left;
}
.emotionItem{
  width:10%;
  font-size:20px;
}

.topicone{
  margin-bottom:5px;
  border-bottom:1px solid #D3D3D3;
}

.othersone{
  margin-bottom:5px;
  border-bottom:1px solid #D3D3D3;
}

/*包含以下四种的链接*/
.emotionItem {
    text-decoration: none;
}
/*正常的未被访问过的链接*/
.emotionItem:link {
    text-decoration: none;
}
/*已经访问过的链接*/
.emotionItem:visited {
    text-decoration: none;
}
/*鼠标划过(停留)的链接*/
.emotionItem:hover {
    text-decoration: none;
}
/* 正在点击的链接*/
.emotionItem:active {
    text-decoration: none;
}
</style>