<template>
  <q-layout view="hHr Lpr lFf">
    <q-layout-header class="row justify-center shadow-1">
      <q-toolbar
        color=""
        inverted
        :class="maxWidthClass"
      >

        <!-- <q-btn
          class="mobile-only"
          flat
          round
          dense
          @click="showLeft=!showLeft"
          icon="menu"
        /> -->
        <!-- <img
          class=""
          src=""
        > -->

        <q-toolbar-title>
          <!-- <q-btn-group class="desktop-only">
            <q-btn
              v-if="this.userInfo"
              class="q-mx-sm"
              flat
              @click="$router.push('/posts')"
              icon=""
            >{{$t('MY_POSTS')}}</q-btn>
            <q-btn
              v-if="this.userInfo"
              class="q-mx-sm"
              flat
              @click="$router.push('/')"
              icon=""
            >{{$t('ALL_POSTS')}}</q-btn>
          </q-btn-group> -->
          <span
            class="cursor-pointer text-lg "
            @click="$router.push('/')"
          >EverPost</span>
        </q-toolbar-title>
        <img
          class="user-avatar cursor-pointer"
          v-if="userInfo"
          :src="userAvatar"
          @click="$router.push('/posts')"
          alt=""
        >

        <q-btn
          v-if="this.userInfo"
          class="q-mx-sm"
          flat
          round
          dense
          @click="$router.push('/draft')"
          icon="add"
        />
        <q-btn
          v-else
          class="q-mx-sm"
          flat
          round
          dense
          @click="fetchUser"
          icon="account_circle"
        />

      </q-toolbar>

    </q-layout-header>

    <!-- optional -->
    <q-layout-drawer
      overlay
      v-model="showLeft"
      side="left"
    >
      <!-- content; any -->
    </q-layout-drawer>

    <!-- optional -->
    <q-layout-drawer
      overlay
      v-model="showRight"
      side="right"
    >
      <!-- content; any -->
    </q-layout-drawer>

    <q-page-container>
      <router-view />
    </q-page-container>
    <payment-modal
      :show="paymentModalShow"
      :assets="getAssets"
      @close="close"
      @pay="callPay"
      @reload="reloadAssets"
    />
  </q-layout>
</template>

<script>
import {
  QLayout,
  QPageContainer,
  QLayoutHeader,
  QLayoutDrawer,
  QToolbar,
  QToolbarTitle,
  QBtn,
  QBtnGroup,
  QBtnDropdown,
  QList,
  QItem,
  QItemMain
} from 'quasar'
import {
  getCache,
  setCache,
  toastError
} from '../utils/util'
import PaymentModal from '../components/PaymentModal'
// import {BOT} from '../utils/constants'
import {
  mapActions,
  mapGetters,
  mapMutations
} from 'vuex'
export default {
  name: 'MyLayout',
  components: {
    QLayout,
    QPageContainer,
    QLayoutHeader,
    QLayoutDrawer,
    QToolbar,
    QToolbarTitle,
    QBtn,
    QBtnGroup,
    QBtnDropdown,
    QList,
    QItem,
    QItemMain,
    PaymentModal
  },
  data() {
    return {
      isLoading: false,
      showRight: false,
      showLeft: false,
      locale: 'en',
      paymentModalShow: false,
      assets: []
    }
  },
  async created() {
    const code = this.$route.query.code
    if (code === undefined || code === '') {
      // await this.fetchUser()
      let user = getCache('account')
      if (user) this.setAccount(user)
    } else {
      this.isLoading = true
      await this.authUser(code)
    }
    this.initEvent()
  },
  mounted() {
    let lang = getCache('lang') || 'en'
    this.setLang(lang)
  },
  destroyed() {
    this.cancelEvent()
  },
  methods: {
    ...mapActions(['getUserInfo', 'auth', 'fetch', 'getUserAssets']),
    ...mapMutations(['setAccount', 'setAssets']),
    async authUser(code) {
      let res = await this.auth(code)
      if (res) {
        this.fetchUser()
        this.isLoading = false
      }
    },
    async fetchUser() {
      await this.getUserInfo()
    },
    setLang(lang) {
      this.locale = this.$i18n.locale = lang
      import(`src/i18n/${lang}`).then(lang => {
        this.$q.i18n.set(lang.default)
      })
      setCache('lang', lang)
    },
    initEvent() {
      this.$root.$on('payment', this.openPaymenModal)
    },
    cancelEvent() {
      this.$root.$off('payment', this.openPaymenModal)
    },
    async openPaymenModal(callback) {
      try {
      this.paymentCb = callback
      this.paymentModalShow = true
      if (this.getAssets.length === 0) {
        let res = await this.getUserAssets()
        if (res.success) {
          this.assets = res.assets
          this.setAssets(res.assets)
        } else {
          toastError(this.$t('NO_ASSETS'))
          return null
        }
      }
      } catch (e) {
        this.paymentCb({}, false)
      }
    },
    async reloadAssets() {
      this.assets = []
      try {
        let res = await this.getUserAssets()
        if (res.success) {
          this.assets = res.assets
          this.setAssets(res.assets)
        } else {
          toastError(this.$t('NO_ASSETS'))
          return null
        }
      } catch (e) {
        toastError(this.$t('NO_ASSETS'))
        return null
      }
    },
    close() {
      this.paymentModalShow = false
      if (this.paymentCb) {
        this.paymentCb('', false)
      }
    },
    callPay(asset) {
      this.paymentModalShow = false
      if (this.paymentCb) {
        this.paymentCb(asset, true)
      }
    }
  },
  computed: {
    ...mapGetters(['userInfo', 'getAssets']),
    userAvatar() {
      if (this.userInfo) {
        let avatarUrl = this.userInfo.avatar_url
        let len = avatarUrl.length
        return avatarUrl.substr(0, len - 3) + '26'
      }
    },
    maxWidthClass() {
      return this.$q.platform.is.desktop ? 'col-10' : 'col-12'
    }
  }
}
</script>

<style lang="stylus" scoped>
.user-avatar {
  border-radius: 15px;
}
</style>
