const { defineConfig } = require('@vue/cli-service')

// module.exports = defineConfig({
//   transpileDependencies: true
// })

module.exports = {
  publicPath: '/',
  outputDir: 'dist',
  devServer: {
    open: true,
    host: 'localhost',
    port: '8081',
    proxy: {
      '/api': {
        target: 'http://localhost:8080', // 要请求的地址
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/api': ''
        }
      }
    }
  }
};
