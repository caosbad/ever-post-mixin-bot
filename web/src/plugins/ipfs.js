import {
  IPFS_NODE,
  IPFS_GATEWAY
} from '../utils/constants'
import IPFS from 'ipfs-mini'

// leave the export, even if you don't use it
export default ({
  app,
  router,
  Vue
}) => {
  let ipfs = new IPFS()
  ipfs.setProvider({
    host: IPFS_NODE,
    protocol: 'https'
  })
  Vue.prototype.$ipfs = ipfs
  Vue.prototype.$ipfs_gateway = IPFS_GATEWAY
  Vue.prototype.$save2IPFS = function (d, c) {
    ipfs.add(d, c)
  }
}
