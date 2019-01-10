export default {
  userInfo: state => {
    return state.account
  },
  getAssets: state => {
    return state.assets
  },
  getLocale: state => {
    return state.local
  }
}
