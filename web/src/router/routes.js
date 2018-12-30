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
        path: 'post/:id',
        name: 'Post',
        component: load('Post')
      },
      {
        path: 'post/:id/edit',
        name: 'edit',
        component: load('Draft')
      },
      {
        path: 'posts',
        name: 'posts',
        component: load('Posts')
      }
    ]
  },

  { // Always leave this as last one
    path: '*',
    name: '404',
    component: load('404')
  }
]
