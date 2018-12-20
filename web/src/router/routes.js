function load(component, path = false) {
  let comp = path ? () =>
    import(`src/layouts/${component}.vue`) : () =>
    import(`src/pages/${component}.vue`)
  return comp
}

export default [{
    path: '/',
    component: load('Main', true),
    name: 'main',
    // redirect: 'overall',
    children: [{
        path: '/',
        name: 'index',
        component: load('Index')
      },
      {
        path: '/draft',
        name: 'Draft',
        component: load('Draft')
      },
      {
        path: 'somepage/:id',
        name: 'marketDetail',
        component: load('MarketDetail')
      },
      {
        path: 'post/:id',
        name: 'Post',
        component: load('Post')
      }
    ]
  },

  { // Always leave this as last one
    path: '*',
    name: '404',
    component: load('404')
  }
]
