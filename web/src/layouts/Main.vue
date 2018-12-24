<template>
  <q-layout view="hHr Lpr lFf">
    <q-layout-header class="row justify-center">
      <q-toolbar
        color=""
        inverted
        class="col-10 "
      >
        <img
          class=""
          src=""
        >
        <!-- <q-btn
          class="mobile-only"
          flat
          round
          dense
          @click="showLeft = !showLeft"
          icon="menu"
        /> -->
        <q-toolbar-title>Everpost</q-toolbar-title>
        <img
          class="user-avatar"
          v-if="userInfo"
          :src="userAvatar"
          alt=""
        >

        <q-btn
          v-if="this.userInfo"
          class="q-mx-sm"
          flat
          round
          dense
          @click="$router.push('/draft')"
          icon="add_circle_outline"
        />
        <!-- <q-btn
          v-if="this.userInfo"
          class="q-mx-sm"
          flat
          round
          dense
          @click="showRight = !showRight"
          icon="menu"
        /> -->
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
    <q-modal v-model="paymentModal">
      <h4>Basic Modal</h4>
      <q-btn
        color="primary"
        @click="opened = false"
        label="Close"
      />
    </q-modal>
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
  QBtn
} from 'quasar'
import {
  getCache,
  setCache
} from '../utils/util'
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
    QBtn
  },
  data() {
    return {
      isLoading: false,
      leftDrawerOpen: this.$q.platform.is.desktop,
      showRight: false,
      showLeft: false,
      locale: 'en',
      paymentModal: false
    }
  },
  async created() {
    const code = this.$route.query.code
    if (code === undefined || code === '') {
      // await this.fetchUser()
      debugger
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
    ...mapActions(['getUserInfo', 'auth', 'fetch']),
    ...mapMutations(['setAccount']),
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
      this.$root.$on('pay', this.openPaymenModal)
    },
    cancelEvent() {
      this.$root.$off('pay', this.openPaymenModal)
    },
    async openPaymenModal(callback) {
      if (this.getAssets.length) {

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
    }
  }
}
</script>

<style lang="stylus" scoped>
.user-avatar {
  border-radius: 15px;
}
</style>
