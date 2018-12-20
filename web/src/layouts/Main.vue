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
          @click="showRight = !showRight"
          icon="menu"
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
  mapGetters
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
      locale: 'en'
    }
  },
  async created() {
    const code = this.$route.query.code
    if (code === undefined || code === '') {
      await this.fetchUser()
    } else {
      this.isLoading = true
      await this.authUser(code)
    }
  },
  mounted() {
    // let user = getCache('account')
    // console.log(user)
    let lang = getCache('lang') || 'en'
    this.setLang(lang)
  },
  methods: {
    ...mapActions(['getUserInfo', 'auth']),
    async authUser(code) {
      let res = await this.auth(code)
      if (res) {
        //  this.fetchUser()
        this.isLoading = false
        console.log(res)
      }
    },
    async fetchUser() {
      await this.getUserInfo()
      console.log(this.$store.state)
    },
    setLang(lang) {
      this.locale = this.$i18n.locale = lang
      import(`src/i18n/${lang}`).then(lang => {
        this.$q.i18n.set(lang.default)
      })
      setCache('lang', lang)
    }
  },
  computed: {
    ...mapGetters(['userInfo']),
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
