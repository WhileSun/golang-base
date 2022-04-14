import {
  defineConfig
} from 'umi';
import routes from './routes';
import proxy from './proxy';
import layoutSettings from './layoutSettings';
import chainWebpack from './chainWebpack';
const { REACT_APP_ENV } = process.env;
const isEnvProduction = process.env.NODE_ENV === "production";

export default defineConfig({
  base: '/',
  publicPath: './',
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  history: {
    type: 'hash',
  },
  ...chainWebpack,
  layout: {
    // https://umijs.org/zh-CN/plugins/plugin-layout
    locale: false,
    siderWidth: 208,
    ...layoutSettings,
  },
  // locale: {
  //   // default zh-CN
  //   default: 'zh-CN',
  //   antd: true,
  //   // default true, when it is true, will use `navigator.language` overwrite default
  //   baseNavigator: true,
  // },
  dynamicImport: {
    loading: '@ant-design/pro-layout/es/PageLoading',
  },
  nodeModulesTransform: {
    type: 'none',
  },
  routes: routes,
  theme: {
    // '@font-size-base':'12px',
  },
  // Fast Refresh 热更新
  // fastRefresh: {},
  mfsu: {},
  webpack5: {},
  antd: {
    compact: true, // 开启紧凑主题
  },
  proxy: proxy[REACT_APP_ENV || 'dev'],
  // 生产环境去除console日志打印
  terserOptions: {
    compress: {
      drop_console: isEnvProduction,
    },
  },
  devServer: {
    open: true,
    port: 8000,
  },
  // mock:{
  //   exclude:['mock/user.js']
  // }
});
