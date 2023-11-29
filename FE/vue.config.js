const { defineConfig } = require('@vue/cli-service')
const path = require('path')
const MiniCssExtractPlugin = require("mini-css-extract-plugin");


module.exports = defineConfig({
  transpileDependencies: [
    'quasar'
  ],

  pluginOptions: {
    quasar: {
      importStrategy: 'kebab',
      rtlSupport: false
    }
  },

  configureWebpack: {
    plugins: [new MiniCssExtractPlugin()],
  },

  devServer: {
    // outputDir: path.resolve(__dirname,'../BE/public'),
    proxy:{
      '^/api':{
        target: 'http://localhost:3000'
      }
    }
  }
})
