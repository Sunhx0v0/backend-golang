<template>
  <div>
    <div class="personal">
      <el-container>
        <el-header>

        </el-header>
        <el-main>
          <div class="block"><el-avatar :size="180" :fit="fit" :src="require('../assets/head1.webp')"></el-avatar></div>
          <h3 class="user_name">Avenir.</h3>
          <span class="id">小红书号:</span>
          <span class="id">xxx</span>
          <!-- <el-row> -->
          <el-button type="primary" icon="el-icon-edit" round class="edit">资料修改</el-button>
          <!-- <h4>签名</h4> -->
          <el-row type="flex" class="row-bg" justify="space-between">
            <el-col :span="4">
              <div></div>
            </el-col>
            <el-col :span="4">
              <div>
                <el-table :data="tableData" style="width:210px">
                  <el-table-column prop="guanzhu" label="关注" width="60" align="center">
                  </el-table-column>
                  <el-table-column prop="fans" label="粉丝" width="60" align="center">
                  </el-table-column>
                  <el-table-column prop="liked" label="获赞与收藏" width="90" align="center">
                  </el-table-column>
                </el-table>
              </div>
            </el-col>
            <el-col :span="4">
              <div></div>
            </el-col>
          </el-row>
        </el-main>
        <el-footer>
          <el-row type="flex" justify="center">
            <el-col :span="15" :offset="11">
              <el-tabs v-model="activeName" @tab-click="handleClick">
                <el-tab-pane label="笔记" name="first">

                </el-tab-pane>
                <el-tab-pane label="收藏" name="second">

                </el-tab-pane>
                <el-tab-pane label="点赞" name="third">

                </el-tab-pane>
              </el-tabs>
            </el-col>
          </el-row>
        </el-footer>
      </el-container>
    </div>
  </div>
</template>

<script>
import axios from 'axios';       // 需要先npm install 并且import进来
export default {
  data() {
    return {
      tableData: [{
        guanzhu: '0',
        fans: '16',
        liked: '0'
      }],
      activeName: "first",
    }
  },
  created() {  // created方法：首次启动该界面时加载的
    axios.post('http://localhost:8085/host',this.tableData[0].fans)       // 通过axios方法获取，get/post前后端需要保持一致
      .then(response => {
        this.tableData[0].guanzhu = response.data.guanzhu;
        this.tableData[0].fans = response.data.fans;
        this.tableData[0].liked = response.data.liked;
      }) // response是从后端接收到的报文的名字，可以任取；response中有很多字段，其中.data为接收到的数据；
      // axios.get('http://localhost:8085/host')       // 通过axios方法获取，get/post前后端需要保持一致
      // .then(response => {
      //   this.tableData[0].guanzhu = response.data.guanzhu;
      //   this.tableData[0].fans = response.data.fans;
      //   this.tableData[0].liked = response.data.liked;
      // })
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  },
}
</script>

<style>
.el-header,
.el-footer {
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

body>.el-container {
  margin-bottom: 40px;
}

.user_name {
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
}

.id {
  color: #909399;
  letter-spacing: 1px;
  margin-top: 0px;
  margin-bottom: 1px;
  font-weight: 500;
  font-size: 10px;
}

.el-row {
  margin-bottom: 20px;
  /* &:last-child {
      margin-bottom: 0;
    } */
}

.el-col {
  border-radius: 0px;
}

.row-bg {
  padding: 10px 0;
}

.el-table {
  border: none;
  fit: "true";
  size: "mini";
  text-align: center;
  color: #303133;
  font-weight: 400;

}

.group>.el-table--enable-row-hover .el-table__body tr:hover>td {
  background-color: transparent !important;
}
</style>

<style lang="scss" scoped>
//去掉el-tab-pane底部灰色线条
.el-tabs__nav-wrap::after {
  height: 0 !important;
}

.el-tabs /deep/.el-tabs__nav-wrap::after {
  background-color: #fff;
}

//修改样式
/deep/.el-tabs__item {
  padding: 0 20px 0 0;
}

.edit {
  color: white;
}
</style>
