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
    if (res.data && !res.data.error) {
      let user = res.data.data
      state.account = user
      setCache('account', user)
      state.account = user
      return user
    } else {
      return {
        success: false
      }
    }
  },
  auth: async ({
    commit,
    state
  }, code) => {
    let res = await api.auth(code)
    if (res.data && !res.data.error) {
      let data = res.data.data
      setCache('Authorization', data.authentication_token)
      setCache('account', data)
      state.account = data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getUserAssets: async ({
    commit,
    state
  }, params) => {
    let res = await api.getAssets()
    if (res.data && !res.data.error) {
      let data = res.data.data
      let assets = data.assets
      if (assets.length) {
        data.assets = assets.filter(asset => asset.balance > 0)
      }
      state.assets = data.assets
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getAllPost: async ({
    commit,
    state
  }, params) => {
    let res = await api.getAllPosts(params)
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  postContent: async ({
    commit,
    state
  }, params) => {
    let res = await api.post(params)
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getDrafts: async ({
    commit,
    state
  }, params) => {
    let res = await api.getMyDrafts(params)
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getMyPosts: async ({
    commit,
    state
  }, params) => {
    let res = await api.getMyPosts(params)
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getPost: async ({
    commit,
    state
  }, id) => {
    let res = await api.getPost({
      id
    })
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  },
  getUser: async ({
    commit,
    state
  }, id) => {
    let res = await api.getUser({
      id
    })
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
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
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
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
    if (res.data && !res.data.error) {
      let data = res.data.data
      return data
    } else {
      return {
        success: false
      }
    }
  }
}
