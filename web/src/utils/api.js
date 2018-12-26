import axios from './axiosWrap'
import {
  urls,
  BOT
} from './constants'
import {
  getCache
} from '../utils/util'

axios.interceptors.request.use(config => {
  const token = getCache('Authorization')
  config.headers.common['Authorization'] = 'Bearer ' + token
  return config
}, error => {
  return Promise.reject(error)
})

axios.interceptors.response.use(response => {
  if (response.data.error && response.data.error.code === 401) {
    location.replace(`https://mixin.one/oauth/authorize?client_id=${BOT.clientId}&scope=${BOT.scope}&response_type=code`)
    return Promise.reject(response)
  }
  return response
}, function (error) {
  if (error.response) {
    switch (error.response.status) {
      case 401:
    }
  }
  return Promise.reject(error)
})

const fetch = (url, method, data = {}, postHeaders) => {
  for (let k in data) {
    if (url.indexOf(':' + k) !== -1) {
      url = url.replace(':' + k, data[k])
      delete data[k]
    }
  }

  let type = method.toLowerCase()
  let res = {}
  if (type === 'get') {
    res = axios.get(url + '?' + json2url(data))
  } else if (type === 'post') {
    res = axios.post(url, data, postHeaders)
  } else if (type === 'put') {
    res = axios.put(url, data, postHeaders)
  } else if (type === 'delete') {
    res = axios.delete(url, data, postHeaders)
  }
  return res
  // return axios[type](url, params)
}
const json2url = json => {
  var arr = []
  var str = ''
  for (var i in json) {
    str = i + '=' + json[i]
    arr.push(str)
  }
  return arr.join('&')
}

const api = {}

api.getMe = () => {
  return fetch(urls.me(), 'get')
}
api.verify = (params) => {
  return fetch(urls.verify(), 'get', params)
}
api.auth = (code) => {
  return fetch(urls.auth(), 'post', {
    code
  })
}
api.getAssets = (params) => {
  return fetch(urls.assets(), 'get', params)
}
api.getAsset = (params) => {
  return fetch(urls.asset(), 'get', params)
}
api.getMyPosts = (params) => {
  return fetch(urls.getPosts(), 'get', params)
}
api.getMyDrafts = (params) => {
  return fetch(urls.getDrafts(), 'get', params)
}

api.getAllPosts = (params) => {
  return fetch(urls.getPosts(), 'get', params)
}
api.getPost = (params) => {
  return fetch(urls.posts(), 'get', params)
}
api.post = (params) => {
  return fetch(urls.posts(), 'post', params)
}
api.savePost = (params) => {
  return fetch(urls.posts(), 'post', params)
}
api.updatePost = (params) => {
  return fetch(urls.posts(), 'put', params)
}
api.publishPost = (params) => {
  return fetch(urls.posts(), 'post', params)
}

api.getUser = (params) => {
  return fetch(urls.getUser(), 'get', params)
}

api.follow = (params) => {
  return fetch(urls.subscriber(), 'post', params)
}
api.unFollow = (params) => {
  return fetch(urls.subscriber(), 'delete', params)
}
api.isSub = (params) => {
  return fetch(urls.subscriber(), 'get', params)
}

api.createDraft = (params) => {
  return fetch(urls.getDrafts(), 'post', params)
}
api.updateDraft = (params) => {
  return fetch(urls.drafts(), 'put', params)
}
api.deleteDraft = (params) => {
  return fetch(urls.drafts(), 'delete', params)
}

// api.deposit = params => {
//   return createInTransfer(urls.dappId[process.env.NODE_ENV], 'koumei.KMC', params.amount, params.secret, params.secondPassword || '')
// }

export default api
