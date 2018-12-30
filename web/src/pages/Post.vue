<template>
  <q-page
    class="row justify-center"
    panding
  >
    <div
      v-if="post"
      :class="maxWidthClass"
    >
      <div>
        <div class="row q-py-lg">
          <div class="post-header col-12">{{post.title}}</div>
          <div
            v-if="userInfo"
            :class="btnBarClass"
          >
            <div v-if="isSelf">
              <q-btn
                v-if="!isIPFS"
                :label="$t('EDIT')"
                color="primary"
                class="q-mx-xs"
                rounded
                outline
                :disabled="payLoading"
                size="sm"
                :to="'/post/'+post.post_id+'/edit'"
              />
              <q-btn
                v-if="!isPub"
                :loading="payLoading"
                :label="$t('PUBLISH')"
                color="info-dark"
                class="q-mx-xs"
                @click="doPublish"
                outline
                rounded
                size="sm"
              />

              <q-btn
                v-if="isPub"
                class="q-mx-xs"
                color="grey-9"
                :label="$t('VIEW_ON_TELEGRAPH')"
                outline
                rounded
                :disabled="payLoading"
                @click="viewTelegraph"
                size="sm"
              />
              <q-btn
                v-if="isPub && !isIPFS"
                :loading="payLoading"
                :label="$t('IPFS_IT')"
                class="q-mx-xs"
                color="cyan"
                @click="doIPFS"
                outline
                rounded
                size="sm"
              />
            </div>
            <q-btn
              v-if="isIPFS"
              :label="$t('VIEW_ON_IPFS')"
              class="q-mx-xs"
              color="teal-14"
              outline
              rounded
              :disabled="payLoading"
              @click="openIPFSUrl"
              size="sm"
            />

          </div>
        </div>
        <div
          v-if="user"
          class="row author-container "
        >
          <img
            class="author-img"
            :src="authorImg"
            alt=""
          >
          <p class="author-info ">
            {{user.full_name}}
            <q-btn
              v-if="!isSelf"
              class="sub-btn"
              :color="isSub ? 'red' : 'green'"
              :loading="optLoad"
              :icon="isSub? 'remove_circle_outline' : 'add_circle_outline'"
              @click="toggleSub"
              :disabled="payLoading"
              flat
              round
              dense
              size="md"
            >
              <q-tooltip>
                {{$t(isSub? 'UNFOLLOW' : 'FOLLOW')}}
              </q-tooltip>
            </q-btn>
            <br />
            <span class="q-mt-xs">
              {{post.created_at}}
            </span>
          </p>
        </div>
      </div>
      <vue-markdown
        class="q-my-lg"
        @rendered="(html)=>htmlContent=html"
      >{{post.content}}</vue-markdown>
      <div class="row justify-center q-my-sm">

        <!-- <q-btn
          icon="favorite_border"
          round
          dense
          color="red"
          :loading="payLoading"
          :disabled="payLoading"
          @click="doDonate"
          flat
        /> -->
        <img
          class="cursor-pointer icon-svg"
          :src="donateIcon"
          alt="$t('DONATE_AUTHOR')"
          @click="doDonate"
        >
      </div>
      <div class="row justify-center q-my-sm">
        <p>
          {{$t('DONATE_AUTHOR')}}
        </p>
      </div>
    </div>
  </q-page>
</template>

<script>
import {
  QPage,
  QBtn,
  openURL
} from 'quasar'
import {
  mapActions,
  mapGetters
} from 'vuex'
import VueMarkdown from 'vue-markdown'
import {
  ASSETS,
  OPT_PIRCE,
  PAY_URL,
  IPFS_GATEWAY
} from '../utils/constants'
// import { toastError } from '../utils/util'
import MEditor from '../components/MEditor'
import uuidv4 from 'uuid/v4'
import {
  toastError
} from '../utils/util'

import {
  renderPostPage
} from '../utils/ipfsTemp'
import donateIcon from '../assets/donate.svg'
// import _ from 'lodash'

export default {
  name: 'Post',
  components: {
    QPage,
    MEditor,
    QBtn,
    VueMarkdown
  },
  data() {
    return {
      post: null,
      user: null,
      optLoad: false,
      isSub: false,
      payLoading: false,
      htmlContent: '',
      donateIcon
    }
  },
  async mounted() {
    let params = this.$route.params
    let id = params['id']
    let donate = this.$route.query['donate']
    if (id) {
      await this.initData(id)
      if (donate) {
        this.doDonate()
      }
    }
  },
  methods: {
    ...mapActions(['getPost', 'getUser', 'fetch', 'send']),
    async initData(id) {
      try {
        let post = await this.getPost(id)
        this.post = post
        let user = await this.getUser(post.user_id)
        this.user = user
        this.getSubStatus()
      } catch (e) {
        console.log(e)
      }
    },
    async toggleSub() {
      try {
        this.optLoad = true
        let params = {
          id: this.authorId
        }
        let res
        if (!this.isSub) {
          res = await this.fetch({
            method: 'follow',
            params
          })
        } else {
          res = await this.fetch({
            method: 'unFollow',
            params
          })
        }
        if (res) {
          console.log(res)
          this.isSub = !res.success
        }
        this.optLoad = false
      } catch (e) {
        // TODO
        this.optLoad = false

        console.log(e)
      }
    },
    viewTelegraph() {
      openURL(this.post.telegraph_url)
    },
    async getSubStatus() {
      if (!this.userInfo) return
      try {
        let res = await this.fetch({
          method: 'isSub',
          params: {
            id: this.authorId
          }
        })
        if (res) {
          this.isSub = res.isSub
        }
      } catch (e) {
        // TODO
        // this.done(false)
        console.log(e)
      }
    },
    doPublish() {
      // this.$root.$emit('payment', () => this.publish)
      this.payLoading = true
      let traceId = uuidv4()
      let url = this.getPayUrl({
        asset_id: ASSETS.CNB,
        amount: OPT_PIRCE.PUB
      }, traceId, this.post.post_id)
      // let url = `${PAY_URL}asset=${ASSETS.CNB}&amount=${OPT_PIRCE.PUB}&trace=${traceId}&memo=${this.post.post_id}`
      let tab = window.open('')
      tab.location = url
      setTimeout(() => {
        this.waitForPubPayment(traceId, tab, () => this.publish(traceId))
      }, 3000)
    },
    async publish(traceId) {
      try {
        let post = this.post
        let res = await this.send({
          method: 'publishPost',
          params: { ...post,
            id: post.post_id,
            htmlContent: this.htmlContent,
            traceId
          }
        })
        if (res.path) {
          this.payLoading = false
          this.initData(res.post_id)
        }
      } catch (e) {
        this.payLoading = false
        console.log(e)
      }
    },
    async waitForPubPayment(traceId, tab, cb = () => {}) {
      try {
        let response = await this.fetch({
          method: 'verify',
          params: {
            id: traceId
          }
        })
        if (response.success === false) {
          setTimeout(() => {
            this.waitForPubPayment(traceId, tab, cb)
          }, 3000)
          return true
        } else if (response.trace_id) {
          if (tab !== undefined) {
            tab.close()
          }
          if (cb) cb(response.trace_id)
        }
      } catch (e) {
        console.log(e)
        this.payLoading = false
        toastError(this.$t('ERROR'))
      }
    },
    doDonate() {
      this.payLoading = true
      this.$root.$emit('payment', (asset, flag) => this.donate(asset, flag))
    },
    donate(asset, flag) {
      if (flag) {
        let traceId = uuidv4()
        let url = this.getPayUrl(asset, traceId)
        let tab = window.open('')
        tab.location = url
        this.payLoading = false
        // console.log(url)
      } else {
        this.payLoading = false
      }
    },
    // async waitForDonatePayment(traceId, tab, cb = () => {}) {
    //    try {
    //     let response = await this.fetch({
    //       method: 'verify',
    //       params: {
    //         id: traceId
    //       }
    //     })
    //     if (response.success === false) {
    //       setTimeout(() => {
    //         this.waitForPubPayment(traceId, tab, cb)
    //       }, 3000)
    //       return true
    //     } else if (response.trace_id) {
    //       if (tab !== undefined) {
    //         tab.close()
    //       }
    //       if (cb)cb(response.trace_id)
    //     }
    //   } catch (e) {
    //     console.log(e)
    //     this.payLoading = false
    //     toastError(this.$t('ERROR'))
    //   }
    // },
    getPayUrl(asset, traceId, memo = '') {
      let assetId = asset.asset_id
      let amount = asset.amount
      let url = `${PAY_URL}asset=${assetId}&amount=${amount}&trace=${traceId}&memo=${this.post.post_id}`
      if (memo.length) url = `${url}&memo=${memo}`
      return url
    },
    doIPFS() {
      this.payLoading = true
      let traceId = uuidv4()
      let url = this.getPayUrl({
        asset_id: ASSETS.CNB,
        amount: OPT_PIRCE.IPFS
      }, traceId)
      let tab = window.open('')
      tab.location = url
      setTimeout(() => {
        this.waitForPubPayment(traceId, tab, () => this.ipfs(traceId))
      }, 5000)
    },
    async ipfs(traceId) {
      try {
        let post = this.post
        let postHtml = renderPostPage(post.title, post.description, this.userInfo.full_name, this.htmlContent, this.authorId, this.post.post_id)
        this.$save2IPFS(postHtml, async (err, datahash) => {
          if (err) {
            toastError(this.$t('IPFS_ERROR'))
            console.log(err)
            this.payLoading = false
            return null
          } else if (!err && datahash) {
            openURL(`${IPFS_GATEWAY}${datahash}/`)
            let res = await this.send({
              method: 'updateDraft',
              params: {
                id: post.post_id,
                ipfsId: datahash,
                traceId
              }
            })
            if (res.ipfs_id) {
              this.payLoading = false
              this.initData(res.post_id)
            }
          }
        })
      } catch (e) {
        this.payLoading = false
        console.log(e)
      }
      // this.$save2IPFS
    },
    openIPFSUrl() {
      let url = this.post && this.post.ipfs_id ? `${IPFS_GATEWAY}${this.post.ipfs_id}/` : null
      openURL(url)
    }
    // done(flag = true, cb = () => {}) {
    //   _.delay(() => {
    //     this.payLoading = false
    //     let alert = flag ? toast : toastError
    //     let msg = flag ? 'SUCCESS' : 'ERROR'
    //     alert(this.$t(msg))
    //     if (cb) cb()
    //   }, 3000)
    // }
  },
  computed: {
    ...mapGetters(['userInfo']),
    editor() {
      return this.$refs.myQuillEditor.quill
    },
    locale() {
      let isZH = this.$i18n.locale === 'zh'
      return isZH ? 'zh-CN' : 'en'
    },
    isSelf() {
      return this.userInfo && this.user ? this.userInfo.user_id === this.user.user_id : false
    },
    isPub() {
      return this.post ? this.post.path : false
    },
    isIPFS() {
      return this.post ? this.post.ipfs_id : false
    },
    authorImg() {
      if (this.user) {
        let avatarUrl = this.user.avatar_url
        let len = avatarUrl.length
        return avatarUrl.substr(0, len - 3) + '48'
      } else {
        return null
      }
    },
    authorId() {
      return this.user.user_id
    },
    maxWidthClass() {
      return this.$q.platform.is.desktop ? 'col-10 justify-center' : 'col-11 justify-center'
    },
    btnBarClass() {
      let classStr = 'col-12 flex items-center btns '
      let barClass = this.$q.platform.is.desktop ? 'justify-end' : 'justify-start'
      return classStr + barClass
    }
  }
}
</script>

<style lang="stylus" scoped>
.post-header {
  display: block;
  font-size: 2.5em;
  margin-block-start: 0.67em;
  // margin-block-end: 0.67em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}

.author-container {
  .author-img {
    width: 55px;
    height: 55px;
    border-radius: 100%;
    border: solid 0.5px;
    display: inline-block;
    vertical-align: middle;
  }

  .author-info {
    color: #79828b;
    font-family: CustomSansSerif, sans-serif;
    display: inline-block;
    vertical-align: middle;
    margin: 0 15px;
  }
}
</style>
