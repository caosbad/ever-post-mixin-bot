import api from '../utils/api'
// import _ from 'lodash'
import {
  setCache
} from '../utils/util'

export default {
  getUserInfo: async ({
    commit,
    state
  }, args) => {
    let res = await api.getMe()
    if (res.data) {
      let user = res.data.data
      state.account = user
      setCache('account', user)
      state.account = user
      return user
    } else {
      return null
    }
  },
  auth: async ({
    commit,
    state
  }, code) => {
    let res = await api.auth(code)
    if (res.data) {
      let data = res.data.data
      setCache('Authorization', data.authentication_token)
      setCache('account', data)
      state.account = data
      return data
    } else {
      return null
    }
  },
  getUserAssets: async ({
    commit,
    state
  }, params) => {
    let res = await api.getAssets()
    if (res.data) {
      let assets = res.data.data
      state.assets = assets
      return assets
    } else {
      return null
    }
  },
  getAllPost: async ({
    commit,
    state
  }, params) => {
    let res = await api.getAllPosts(params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  postContent: async ({
    commit,
    state
  }, params) => {
    let res = await api.post(params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  getDrafts: async ({
    commit,
    state
  }, params) => {
    let res = await api.getMyDrafts(params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  getMyPosts: async ({
    commit,
    state
  }, params) => {
    let res = await api.getMyPosts(params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  getPost: async ({
    commit,
    state
  }, id) => {
    let res = await api.getPost({
      id
    })
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  getUser: async ({
    commit,
    state
  }, id) => {
    let res = await api.getUser({
      id
    })
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  fetch: async ({
    commit,
    state
  }, payload) => {
    let {
      method,
      params
    } = payload
    let res = await api[method](params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  },
  send: async ({
    commit,
    state
  }, payload) => {
    let {
      method,
      params
    } = payload
    let res = await api[method](params)
    if (res.data) {
      let data = res.data.data
      return data
    } else {
      return null
    }
  }
}
