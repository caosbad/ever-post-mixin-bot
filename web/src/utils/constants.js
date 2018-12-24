// model alike
const I18N_OPT = [{
    label: '中文简体',
    value: 'zh'
  },
  {
    label: 'English',
    value: 'en'
  }
]

const mixinHost = 'https://api.mixin.one'
const serverHost = 'http://localhost:7008'
// const dappId = {
//   development: '91a9aeeaa6609ebd73afc7c542eec02e7510c23a65d2e1db2840c46f74f67ceb',
//   production: '91a9aeeaa6609ebd73afc7c542eec02e7510c23a65d2e1db2840c46f74f67ceb'
// }

const BOT = {
  clientId: 'ee7ef479-e9e7-44ba-9755-6be4e1d826d7',
  scope: 'PROFILE:READ+ASSETS:READ',
  // scope: 'PROFILE:READ',
  clientSecret: 'f245bfa9bb79b5e076e2f905e2d66f7c658057d069a2d28c704aa263db87c0e8'
}

const urls = {
  me: () => {
    return `${serverHost}/me`
  },
  getUser: () => {
    return `${serverHost}/user/:id`
  },
  auth: () => {
    return `${serverHost}/auth`
  },
  getToken: () => {
    return `${mixinHost}/oauth/token`
  },
  verify: (traceId) => {
    return `${mixinHost}/verify/:id`
  },
  assets: () => {
    return `${mixinHost}/assets`
  },
  asset: (assetId) => {
    return `${mixinHost}/assets/:id`
  },
  getPosts: () => {
    return `${serverHost}/posts`
  },
  posts: () => {
    return `${serverHost}/posts/:id`
  },
  getDrafts: () => {
    return `${serverHost}/drafts`
  },
  drafts: () => {
    return `${serverHost}/drafts/:id`
  },
  post: () => {
    return `${serverHost}/post`
  },
  myPost: () => {
    return `${serverHost}/myPosts`
  },
  subscriber: () => {
    return `${serverHost}/subscriber/:id`
  }
}

export {
  I18N_OPT,
  urls,
  BOT
}
