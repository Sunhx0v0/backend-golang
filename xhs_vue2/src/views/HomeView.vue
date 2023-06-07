<template>
  <div id="home">
    <!--放小红书logo-->
    <el-container>

    <el-aside style="width:180px;">
      <el-menu>
         <el-menu-item index="1-1">推荐</el-menu-item>
          <el-menu-item index="1-1">我的关注</el-menu-item>
          <el-menu-item index="1-2">美食</el-menu-item>
          <el-menu-item index="1-3">彩妆</el-menu-item>
          <el-menu-item index="1-3">穿搭</el-menu-item>
          <el-menu-item index="1-3">影视</el-menu-item>
          <el-menu-item index="1-3">职场</el-menu-item>
          <el-menu-item index="1-3">情感</el-menu-item>
          <el-menu-item index="1-3">家居</el-menu-item>
          <el-menu-item index="1-3">游戏</el-menu-item>
          <el-menu-item index="1-3">旅行</el-menu-item>
          <el-menu-item index="1-3">健身</el-menu-item>
      </el-menu>
    </el-aside>

  <el-container>
  <el-header>
    <!--设置搜索栏-->
    <div>
      <!--elementui有24分栏，span等于几就占据了几分栏，offset为偏移量，el-col里面得有div分隔才生效-->
      <el-row>
      <el-col :span="1" :offset="0" class="left">
        <div><img src="./xhs_logo.png" width="120" height="40" alt="小红书" /></div>
      </el-col>
      <el-col :span="16" :offset="3"  class="center">
        <el-input
          v-model="search"
          @focus="focus"
          @blur="blur"
          @keyup.enter.native="searchHandler"
          placeholder="搜索感兴趣的内容"
        >
          <el-button slot="append" icon="el-icon-search" id="search" @click="searchHandler">搜索</el-button>
        </el-input>
        <!---设置z-index优先级,防止被走马灯效果遮挡-->
        <el-card
          @mouseenter="enterSearchBoxHanlder"
          v-if="isSearch"
          class="box-card"
          id="search-box"
          style="position:relative;z-index:15"
        >
          <dl v-if="isHistorySearch">
            <dt class="search-title" v-show="history">历史搜索</dt>
            <dt class="remove-history" v-show="history" @click="removeAllHistory">
              <i class="el-icon-delete"></i>清空历史记录
            </dt>
            <el-tag
              v-for="search in historySearchList"
              :key="search.id"
              closable
              :type="search.type"
              @close="closeHandler(search)"
              style="margin-right:5px; margin-bottom:5px;"
            >{{search.name}}</el-tag>
            <dt class="search-title">热门搜索</dt>
            <dd v-for="search in hotSearchList" :key="search.id">{{search}}</dd>
          </dl>
          <dl v-if="isSearchList">
            <dd v-for="search in searchList" :key="search.id">{{search}}</dd>
          </dl>
        </el-card>
      </el-col>
    </el-row>
    
    </div>
  </el-header>

  <el-main>
  <el-row>
  <el-col :span="3" :offset="0" class="left">
  <el-button round icon="el-icon-s-home" size="medium" >  发现  </el-button>
  </el-col>
  <el-col :span="3" :offset="3" class="left">
  <el-button round icon="el-icon-circle-plus-outline" size="medium">  发布  </el-button>
  </el-col>
  <el-col :span="3" :offset="3" class="left">
  <el-button round icon="el-icon-tableware" size="medium">  我的  </el-button>
  </el-col>
  </el-row>
  </el-main>
  
  </el-container>
  </el-container>
  </div>

</template>

<script>
import RandomUtil from "/src/utils/randomUtil";
import Store from "/src/utils/store";
export default {
  name: 'HomeView',
  data() {
    return {
      search: "", //当前输入框的值
      isFocus: false, //是否聚焦
      hotSearchList: ["暂无热门搜索"], //热门搜索数据
      historySearchList: [], //历史搜索数据
      searchList: ["暂无数据"], //搜索返回数据,
      history: false,
      types: ["", "success", "info", "warning", "danger"] //搜索历史tag式样
    };
  },
  methods: {
    focus() {
      this.isFocus = true;
      this.historySearchList =
        Store.loadHistory() == null ? [] : Store.loadHistory();
      this.history = this.historySearchList.length == 0 ? false : true;
    },
    blur() {
      var self = this;
      this.searchBoxTimeout = setTimeout(function() {
        self.isFocus = false;
      }, 300);
    },
    enterSearchBoxHanlder() {
      clearTimeout(this.searchBoxTimeout);
    },
    searchHandler() {
      //随机生成搜索历史tag式样
      let n = RandomUtil.getRandomNumber(0, 5);
      let exist =
        this.historySearchList.filter(value => {
          return value.name == this.search;
        }).length == 0
          ? false
          : true;
      if (!exist) {
        this.historySearchList.push({ name: this.search, type: this.types[n] });
        Store.saveHistory(this.historySearchList);
      }
      this.history = this.historySearchList.length == 0 ? false : true;
    },
    closeHandler(search) {
      this.historySearchList.splice(this.historySearchList.indexOf(search), 1);
      Store.saveHistory(this.historySearchList);
      clearTimeout(this.searchBoxTimeout);
      if (this.historySearchList.length == 0) {
        this.history = false;
      }
    },
    removeAllHistory() {
      Store.removeAllHistory();
    }
  },
  computed: {
    isHistorySearch() {
      return this.isFocus && !this.search;
    },
    isSearchList() {
      return this.isFocus && this.search;
    },
    isSearch() {
      return this.isFocus;
    }
  }

}
</script>

<style>
#home{
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 0px;
}

.threeline {
top: 0ch;
width: 90px;
height: 30px;
}
.left {
  margin-left: 0px;
  left: 0cap;
}
.center {
  margin-left: 300px;
}
#search {
  background-color: #e13737;
  border-radius: 0%;
}
.search-title {
  color: #bdbaba;
  font-size: 15px;
  margin-bottom: 5px;
}
.remove-history {
  color: #bdbaba;
  font-size: 15px;
  float: right;
  margin-top: -22px;
}
#search-box {
  width: 555px;
  height: 300px;
  margin-top: 0px;
  padding-bottom: 0px;
}

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
  background-color: #d5103bcc;
}


</style>
