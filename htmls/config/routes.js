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
    path: '/work',
    name: '工作管理',
    routes: [{
        icon: 'icon-wodexiangmu',
        path: '/work/task',
        name: '任务流程',
        routes: [{
            icon: 'icon-xiangmu_xiangmuguanli',
            path: '/work/task/project',
            name: '项目管理',
            component: 'work/task/project',
          },
          {
            icon: 'icon-renwuguanli',
            path: '/work/task/record',
            name: '任务管理',
            component: 'work/task/record',
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
