const AUTH_FAIL = 10008;
const loginPath  =  '/user/login';
const isDev = true;//process.env.NODE_ENV === 'development';
const defaultPasswd = "\u521d\u59cb\u5316\u5bc6\u7801\u4e3a\u0061\u0031\u0032\u0033\u0034\u0035\u0036\u0021";
const roleSuperName = "super_admin"; //超级管理员权限标识
const userSuperName = "system";
const iconfontUrl = "//at.alicdn.com/t/font_2215021_guez560qs.js";

export {loginPath,AUTH_FAIL,isDev,defaultPasswd,roleSuperName,userSuperName,iconfontUrl};