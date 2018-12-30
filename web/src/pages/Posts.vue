<template>
  <q-page
    padding
    class="justify-center"
  >
    <div :class="maxWidthClass">
      <!-- <q-page-sticky
        position="top"
        class="row"
        expand
      > -->
      <div class="tabs-container">
        <q-tabs
          v-model="selectedTab"
          align="center"
          color="gray"
          @select="tabChange"
          two-lines
          inverted
        >
          <q-tab
            slot="title"
            name="draft"
            icon="note"
            :label="$t('DRAFTS')"
          />
          <q-tab
            slot="title"
            name="telegraph"
            icon="title"
            :label="$t('TELEGRAPH')"
          />
          <q-tab
            slot="title"
            name="ipfs"
            icon="all_inclusive"
            :label="$t('IPFS')"
          />
        </q-tabs>
      </div>
      <!-- </q-page-sticky> -->
      <list-container
        :datas="posts"
        @loadMore="loadMore"
        :expand="true"
        :count="count"
      >
        <template
          slot="item"
          slot-scope="props"
        >
          <post-item
            :type="selectedTab"
            :data="props.data"
            :index="props.index"
          />
          <separator />
        </template>
      </list-container>
    </div>
  </q-page>
</template>

<style>
</style>

<script>
import {
  QPage,
  QList,
  QListHeader,
  QItem,
  QItemMain,
  QItemSeparator,
  QItemSide,
  QTabs,
  QTab
} from 'quasar'
import {
  mapActions
} from 'vuex'
import ListContainer from '../components/ListContainer'
import PostItem from '../components/PostItem'
import Separator from '../components/Separator'
import {
  toastError
} from '../utils/util'
import _ from 'lodash'

export default {
  name: 'Index',
  components: {
    QPage,
    QList,
    QListHeader,
    QItem,
    QItemMain,
    QItemSeparator,
    QItemSide,
    ListContainer,
    PostItem,
    Separator,
    QTabs,
    QTab
  },
  data() {
    return {
      selectedTab: 'draft',
      pagination: {
        page: 1,
        limit: 20
      },
      count: 0,
      posts: []
    }
  },
  mounted() {
    this.getDatas()
  },
  methods: {
    ...mapActions(['getMyPosts']),
    async getDatas(append = false) {
      try {
        let res = await this.getMyPosts({
        offset: this.offset,
        limit: this.pagination.limit,
        type: this.selectedTab
      })
      if (res) {
        console.log(res)
        if (append) {
          this.posts = this.posts.concat(res.posts)
        } else {
          this.posts = res.posts
        }
        this.count = res.total_count
      } else {
        toastError(this.$t('LOAD_ERROR'))
      }
      } catch (error) {
        toastError(this.$t('LOAD_ERROR'))
      }
    },
    open(postId) {
      this.$router.push('/')
    },
    async loadMore(index, done) {
      this.pagination.page++
      await this.getDatas(true)
      _.delay(() => done(), 1000)
    },
    tabChange() {
      this.pagination = {
        page: 1,
        limit: 20
      }
      this.getDatas()
    }
  },
  computed: {
    offset() {
      let {
        page,
        limit
      } = this.pagination
      return (page - 1) * limit
    },
    maxWidthClass() {
      return this.$q.platform.is.desktop ? 'col-10 ' : 'col-12 '
    },
    isPost() {
      return this.selectedTab === 'telegraph'
    },
    isDraft() {
      return this.selectedTab === 'draft'
    },
    isIPFS() {
      return this.selectedTab === 'ipfs'
    }
  }
}
</script>

<style lang="stylus">
.tabs-container {
}
</style>
