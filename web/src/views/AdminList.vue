<template>
    <div>
        <h1>分类列表</h1>
        <el-table :data="items">
            <el-table-column prop="_id" label="ID" width="220">
            </el-table-column>
            <el-table-column prop="username" label="用户名">
            </el-table-column>
            <!-- 列表页没必要将用户密码显示 -->
            <!-- <el-table-column prop="password" label="密码">
            </el-table-column> -->
            <el-table-column
            fixed="right"
            label="操作"
            width="100">
                <template slot-scope="scope">
                    <el-button type="text" size="small" @click="$router.push('/admins/edit/' + scope.row._id)">编辑</el-button>
                    <el-button @click="remove(scope.row)" type="text" size="small">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script>
export default {
  data () {
    return {
      items: []
    }
  },
  methods: {
    async fetch () {
      const res = await this.$AXIOS.get('rest/admins')
      this.items = res.data
    },
    remove (row) {
      this.$confirm('是否确定要删除"' + row.name + '"的账号?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async () => {
        // 要想使用await，函数必须使用async
        // await异步执行，待调用接口获取数据完成后再将值传给res，进行下一步操作
        const res = await this.$AXIOS.delete('rest/admins/' + row._id)
        this.$message({
          type: 'success',
          message: '删除成功!'
        })
        if (res.status === 200) {
          // 接口调用成功后，刷新页面
          this.fetch()
        }
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        })
      })
    }
  },
  created () {
    this.fetch()
  }
}
</script>
