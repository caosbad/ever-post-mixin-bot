<template>
  <q-page pandding>
    <div class="row justify-center q-py-lg">
      <div
        v-if="draftable"
        class="col-lg-9 col-xs-11 col-md-10"
      >
        <q-field
          class="col-8 no-shadow q-px-sm "
          :label-width="2"
          :count="64"
        >
          <q-input
            v-model="post.title"
            :float-label="$t('TITLE')"
            color="gray"
            hide-underline
            :loading="loading"
          />
        </q-field>
        <!-- <q-field
          class="col-8 q-mb-lg q-px-sm"
          :label-width="2"
          :count="16"
        >
          <q-input
            v-model="post.name"
            :float-label="$t('NAME')"
            color="gray"
            clearable
            hide-underline
          />
        </q-field> -->
        <q-field
          class="col-8 q-mb-lg q-px-sm"
          :label-width="2"
          :count="140"
        >
          <q-input
            v-model="post.description"
            :float-label="$t('DESC')"
            type="textarea"
            color="gray"
            hide-underline
          />
        </q-field>
        <q-field class="q-mb-lg">
          <m-editor
            :language="locale"
            v-model="post.content"
            class=" "
            @save="onSave"
            @change="onChange"
            :placeholder="$t('NEW_POST_PLACEHOLDER')"
          />
        </q-field>
        <div class="row justify-between">
          <q-btn
            v-if="!isPub"
            :loading="loading"
            :disabled="$v.post.valid"
            class="col-auto m-t-1 self-center"
            color="primary"
            outline
            @click="submit"
          >
            {{$t('SAVE')}}
          </q-btn>
          <q-btn
            v-if="isPub"
            :loading="loading"
            :disabled="$v.post.valid"
            class="col-auto m-t-1 self-center"
            color="primary"
            outline
            @click="rePublish"
          >
            {{$t('REPUBLISH')}}
          </q-btn>
          <q-btn
            v-if="postId"
            :disabled="loading"
            class="col-auto m-t-1 self-center"
            color="primary"
            outline
            @click="preview"
          >
            {{$t('PREVIEW')}}
          </q-btn>

          <q-btn
            outline
            color="info"
            :disabled="loading"
            class="col-auto m-t-1 self-center"
            @click="$router.go(-1)"
          >
            {{$t('CANCEL')}}
          </q-btn>
          <q-btn
            v-if="!isPub"
            color="negative"
            outline
            :disabled="loading"
            class="col-auto m-t-1 self-center"
            @click="callDelete"
          >
            {{$t('DELETE')}}
          </q-btn>
        </div>
      </div>
      <!-- <vue-markdown
        class="q-my-lg "
        @rendered="(html)=>htmlContent=html"
      >{{post.content}}</vue-markdown> -->
    </div>
  </q-page>
</template>

<style>
</style>

<script>
import {
  QField,
  QInput,
  QRadio,
  QBtnToggle,
  QSlider,
  QPage,
  QBtn
} from 'quasar'
import {
  required,
  maxLength,
  minLength
} from 'vuelidate/lib/validators'
import {
  mapActions,
  mapGetters
} from 'vuex'
import {
  toastError,
  toast,
  toastWarn,
  confirm
} from '../utils/util'
import VueMarkdown from 'vue-markdown'
import MEditor from '../components/MEditor'
import _ from 'lodash'

export default {
  name: 'Draft',
  components: {
    QField,
    QInput,
    QRadio,
    QBtnToggle,
    QSlider,
    QPage,
    MEditor,
    QBtn,
    VueMarkdown
  },
  data() {
    return {
      post: {
        title: '',
        content: '',
        description: ''
      },
      htmlContent: '',
      path: '',
      loading: false,
      draftable: true
    }
  },
  validations: {
    // TODO figure out why vuelidate can not bind vue 'this' in rule config
    post: {
      title: {
        required,
        maxLength: maxLength(64),
        minLength: minLength(5)
      },
      // name: {
      //   maxLength: maxLength(16),
      //   minLength: minLength(3)
      // },
      content: {
        required,
        minLength: minLength(10)
      },
      description: {
        maxLength: maxLength(140)
      }
    }
  },
  mounted() {
    if (this.userInfo) {
      let {
        id
      } = this.$route.params
      if (id) this.initData(id)
    } else {
      this.$router.push('/')
    }
  },
  methods: {
    ...mapActions(['getPost', 'send']),
    async initData(id) {
      let post = await this.getPost(id)
      if (!post.post_id) this.$router.push('/')
      if (post.ipfs_id) {
        toastWarn(this.$t('CANNOT_SAVE_IPFS_POST'))
        this.draftable = false
        _.delay(() => this.$router.go(-1), 2500)
      }
      this.checkEditPromise(post)
      if (post) {
        this.post = {
          post_id: post.post_id,
          title: post.title,
          content: post.content,
          description: post.description
        }
        this.path = post.path
      } else {
        this.$router.push('/')
      }
    },
    onChange(val, html) {
      this.post.content = val
      this.htmlContent = html
      // this.$v.post.content.$touch()
    },
    submit() {
      this.updateDraft(this.post.content, this.htmlContent)
    },
    async callDelete() {
      let t = this.$t
      confirm({
        title: t('DELETE_DRAFT'),
        confirm: t('DELETE'),
        cancel: t('CANCEL')
      }, () => {}, () => this.deleteDraft())
    },
    onSave(val, html) {
      if (!this.checkForm() || this.loading) return
      if (this.isNewDraft) {
        this.createDraft(val)
      } else {
        this.updateDraft(val, html)
      }
    },
    async createDraft(val) {
      this.loading = true
      try {
        let post = this.post
        let res = await this.send({
          method: 'createDraft',
          params: { ...post,
            content: val
          }
        })
        if (res) {
          // this.post = res
          this.post = _.merge(res, this.post)
          console.log(this.post)
          // toast(this.$t('SAVE_SUCCESS'))
          this.done()
        }
      } catch (e) {
        this.done(false)
      }
    },
    rePublish() {
      this.updatePost(this.post.content, this.htmlContent)
    },
    async updateDraft(val, htmlContent) {
      this.loading = true
      try {
        let post = this.post
        let res = await this.send({
          method: 'updateDraft',
          params: { ...post,
            id: post.post_id,
            content: val,
            htmlContent
          }
        })
        if (res) {
          // this.post = res
          // toast(this.$t('SAVE_SUCCESS'))
          this.done()
        }
      } catch (e) {
        console.log(e)
        this.done(false)
      }
    },
    async updatePost(val, htmlContent) {
      this.loading = true
      try {
        let post = this.post
        let res = await this.send({
          method: 'updatePost',
          params: { ...post,
            id: post.post_id,
            content: val,
            htmlContent
          }
        })
        if (res) {
          this.done()
        }
      } catch (e) {
        console.log(e)
        this.done(false)
      }
    },
    done(flag = true, cb = () => {}) {
      _.delay(() => {
        this.loading = false
        let alert = flag ? toast : toastError
        let msg = flag ? 'SUCCESS' : 'ERROR'
        alert(this.$t(msg))
        if (cb) cb()
      }, 3000)
    },
    publish() {
      // TODO
      if (this.checkForm()) {}
      // let article = document.createElement('div')
      // article.innerHTML = this.postHtml
    },
    checkForm() {
      this.$v.post.$touch()
      if (this.$v.post.$invalid) {
        toastError(this.$t('FORM_INVALID'))
        return false
      }
      return true
    },
    checkEditPromise(post) {
      if (this.userInfo.user_id !== post.user_id) {
        this.$router.push('/post/' + post.post_id)
      }
    },
    async deleteDraft() {
      try {
        let post = this.post
        let res = await this.send({
          method: 'deleteDraft',
          params: {
            id: post.post_id
          }
        })
        if (res) {
          this.post = res
        }
        this.done(true, () => this.$router.push('/'))
      } catch (e) {
        this.done(false)
      }
    },
    clear() {
      this.post = {
        title: '',
        content: '',
        description: ''
      }
    },
    preview() {
      this.$router.push(`/post/${this.postId}`)
    }
  },
  computed: {
    ...mapGetters(['userInfo']),
    isNewDraft() {
      return !this.post.post_id
    },
    isPub() {
      return !!this.path
    },
    locale() {
      let isZH = this.$i18n.locale === 'zh'
      return isZH ? 'zh-CN' : 'en'
    },
    draftId() {
      return this.$route.params ? this.$route.params.id : ''
    },
    postId() {
      return this.post ? this.post.post_id : ''
    }
  },
  watch: {
    draftId(val) {
      this.clear()
      if (val) this.initData(val)
    }
  }
}
</script>

<style lang="stylus">
.editor-container {
  font-family: CustomSerif, Georgia, Cambria, 'Times New Roman', serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  line-height: 1.58;
  padding: 0;
  margin: 0;
  color: rgba(0, 0, 0, 0.8);
}

.title-input {
  font-family: CustomSansSerif, 'Lucida Grande', Arial, sans-serif;
  font-weight: 700;
  font-style: normal;
  font-size: 32px;
  margin: 21px 21px 12px;
  line-height: 34px;

  ::before {
    position: absolute;
    left: 0;
    right: 0;
    content: attr(data-placeholder);
    color: rgba(0, 0, 0, 0.44);
    font-weight: 400;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    pointer-events: none;
  }

  ::after {
    position: absolute;
    content: attr(data-label);
    color: rgba(0, 0, 0, 0.44);
    border-right: 1px solid rgba(0, 0, 0, 0.15);
    right: 100%;
    top: -6px;
    bottom: -6px;
    padding: 6px 12px;
    font-weight: 400;
    margin: 0 21px;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease;
  }
}

.name-input {
  font-family: CustomSansSerif, 'Lucida Grande', Arial, sans-serif;
  font-size: 15px;
  line-height: 18px;
  color: #79828B;
  margin: 12px 21px;
  font-style: normal;
}
</style>
