<template>
  <q-page
    class="row justify-center"
    panding
  >

  </q-page>
</template>

<script>
import {
  QPage,
  QBtn
} from 'quasar'
import {
  mapActions,
  mapGetters
} from 'vuex'
import VueMarkdown from 'vue-markdown'

// import { toastError } from '../utils/util'
import MEditor from '../components/MEditor'

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
      isSub: false
    }
  },
  mounted() {
    let params = this.$route.params
    let id = params['id']
    if (id) {
      this.initData(id)
    }
  },
  methods: {
    ...mapActions(['getPost', 'getUser', 'fetch', 'send']),
    async initData(id) {
      try {
        let post = await this.getPost(id)
        this.post = post
        if (!post.post_id) this.$router.push('/')
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
        if (this.isSub) {
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
        this.optLoad = true

        console.log(e)
      }
    },
    async getSubStatus() {
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
        console.log(e)
      }
    },
    doPublish() {
      this.$root.$emit('pay', () => this.publish)
    },
    async publish() {
      try {
        let post = this.post
        let res = await this.send({method: 'updatePost', params: {...post, id: post.post_id}})
        if (res) {
          this.post = res
          // todo
        }
      } catch (e) {

      }
    }
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
      return this.userInfo.user_id === this.user.user_id
    },
    isPub() {
      return this.post ? this.post.path : false
    },
    isEver() {
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
    }
  }
}
</script>

<style lang="stylus" scoped>
.post-header {
  display: block;
  font-size: 2em;
  margin-block-start: 0.67em;
  margin-block-end: 0.67em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  font-weight: bold;
}

.author-container {
  .author-img {
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
