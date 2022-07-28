const routes = [{
    path: '/user',
    layout: false,
    routes: [{
        path: '/user',
        routes: [{
          name: 'login',
          path: '/user/login',
          component: './user/Login',
        }, ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    icon: 'icon-gongzuotai',
    path: '/daily',
    name: '日常笔记',
    routes: [{
        icon: 'icon-wodexiangmu',
        path: '/daily/work',
        name: '工作管理',
        routes: [{
            icon: 'icon-xiangmu_xiangmuguanli',
            path: '/daily/work/project',
            name: '项目管理',
            component: 'work/project',
          },
          {
            icon: 'icon-renwuguanli',
            path: '/daily/work/task',
            name: '任务管理',
            component: 'work/task',
          },
        ]
      },
      {
        icon: 'icon-wodexiangmu',
        path: '/daily/md',
        name: '文档管理',
        routes: [
          {
            path: '/daily/md/book',
            name: '文档书籍',
            component: 'md/book',
          },{
            path: '/daily/md/document/edit',
            name: '文档内容编辑',
            component: 'md/document/edit',
            layout:false,
            hideInMenu:true,
          },
          {
            path: '/daily/md/document',
            name: '文档内容',
            component: 'md/document',
            layout:false,
            hideInMenu:true,
          },
        ]
      },
      {
        component: './404',
      },
    ]
  },
  {
    icon: 'icon-xitong',
    path: '/system',
    name: '系统管理',
    routes: [{
        icon: 'icon-yonghuguanli_huaban',
        path: '/system/user',
        name: '用户管理',
        component: 'system/user',
      },
      {
        icon: 'icon-jiaoseshezhi',
        path: '/system/role',
        name: '角色管理',
        component: 'system/role',
      },
      {
        icon: 'icon-caidan',
        path: '/system/menu',
        name: '菜单管理',
        component: 'system/menu',
      },
      {
        icon: 'icon-jiedian',
        path: '/system/perms',
        name: '节点管理',
        component: 'system/perms',
      },
      {
        icon:'icon-xiugaimima',
        path: '/system/passwd',
        name: '修改密码',
        component: 'system/passwd',
      },
      // {
      //   icon:'icon-xiugaimima',
      //   path: '/system/test',
      //   name: '测试',
      //   component: 'product/index',
      // },
      {
        component: './404',
      },
    ]
  },
  {
    component: './404',
  },
]

export default routes;
