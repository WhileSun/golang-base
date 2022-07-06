import { PageLoading } from '@ant-design/pro-layout';
import { history } from 'umi';
import Footer from '@/components/Footer';
import RightContent from '@/components/RightContent';
import { getUserInfo } from '@/services/user';
import { loginPath,isDev,iconfontUrl} from '@/config';
import { loadApi} from '@/utils/tools';
import {getUserRouteList} from '@/services/api';
import routes from '../config/routes';

let PagePerms = {};
let menuRedirect = {};
let homeMenu = false;
let authMenus = [];
let userRoutes = {};
//routes打入页面权限
const putRoutePagePerms = (routes) => {
  routes.forEach((menu) => {
    if (PagePerms[menu.path] !== undefined) {
      menu.perms = PagePerms[menu.path];
    }
    if (menu.routes != undefined && menu.routes.length > 0) {
      putRoutePagePerms(menu.routes);
    }
  });
};
//子目录跳转记录
const getMenuRedirect = (menuData) => {
  menuData.forEach((menu) => {
    if(menu.path == undefined || menu.redirect!=undefined){
      return;
    }
    authMenus.push(menu.path);
    //页面权限集合
    if (menu.pagePerms !== undefined && menu.pagePerms.length > 0) {
      PagePerms[menu.path] = menu.pagePerms;
    }
    if (menu.path != undefined && menu?.routes?.length > 0 && menu.layout !== false) 
    {
      if (!homeMenu) {
        menuRedirect['/'] = menu.routes[0].path;
        homeMenu = true;
      }
      //一级栏目跳转集合
      menuRedirect[menu.path] = menu.routes[0].path;
      getMenuRedirect(menu.routes);
    }
  });
};

const getRouteList = (menuData)=>{
  menuData.forEach((menu) => {
    if(menu.path == undefined || menu.redirect!=undefined){
      return;
    }
    userRoutes[menu.path] = menu;
  });
}

//动态获取菜单栏
const fetchMenuData = async (name) => {
  if (isDev) {
    getMenuRedirect(routes);
    return routes;
  }
  if (name === undefined) {
    return [];
  }
  const serverMenuData = await loadApi(getUserRouteList, {}, (data) => {
    return data;
  });
  getMenuRedirect(serverMenuData);
  console.log(serverMenuData);
  return serverMenuData;
};


/** 获取用户信息比较慢的时候会展示一个 loading */
export const initialStateConfig = {
  loading: <PageLoading />,
};

export async function getInitialState() {
  console.log('initstate');
  const fetchUserInfo = async () => {
    try {
      const resp = await getUserInfo();
      return resp.data;
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };

  // 如果是登录页面，不执行
  if (history.location.pathname !== loginPath) {
    const currentUser = await fetchUserInfo();
    //获取菜单权限
    const menuData = await fetchMenuData(currentUser?.name);
    return {
      fetchUserInfo,
      fetchMenuData,
      settings: {},
      currentUser,
      menuData,
    };
  }
  return {
    fetchUserInfo,
    fetchMenuData,
    settings: {},
  };
}

export const layout = ({initialState}) => {
  return {
    iconfontUrl: iconfontUrl,
    rightContentRender: () => <RightContent />,
    disableContentMargin: false,
    //水印
    waterMarkProps: {
      content: initialState?.currentUser?.name,
    },
    // footerRender: () => <Footer />,
    onPageChange: ({pathname}) => {
      const { location } = history; // 如果没有登录，重定向到 login
      if (!initialState?.currentUser && location.pathname !== loginPath) {
        history.push(loginPath);
        return;
      }
      console.log('onPageChange',menuRedirect);
      if (pathname != loginPath) {
        //实现mix页面跳转
        if (menuRedirect[pathname] !== undefined) {
          history.push(menuRedirect[pathname]);
          return;
        }
        // //无权限页面404展示
        if (!authMenus.includes(pathname)) {
          history.push('/404');
          return;
        }
      }
    },
    contentStyle: {
      margin:'0px',
      backgroundColor: '#fff',
      padding: '15px 10px 0px 15px',
    },
    menu: {
      params: {
        menuData: initialState?.menuData,
      },
      request:(params, menuData) => {
        console.log(params);
        //打入页面权限
        putRoutePagePerms(menuData)
        return initialState?.menuData;
      },
    },
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    ...initialState?.settings,
  };
};