<template>
  <q-page
    padding
    class="row justify-center"
  >
    <div :class="maxWidthClass">
      <list-container
        :datas="posts"
        @loadMore="loadMore"
        :expand="true"
        :count="count"
      >
        <template
          slot="header"
          slot-scope="props"
        >
          <q-list-header>{{$t('RECENT_POSTS')}}</q-list-header>
        </template>
        <template
          slot="item"
          slot-scope="props"
        >
          <post-item
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
  QItemTile,
  QBtn,
  QSpinnerDots
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
    QItemTile,
    QBtn,
    QSpinnerDots,
    ListContainer,
    PostItem,
    Separator
  },
  data() {
    return {
      pagination: {
        page: 1,
        limit: 20
      },
      count: 0,
      posts: []
    }
  },
  mounted() {
    this.getPosts()
  },
  methods: {
    ...mapActions(['getAllPost']),
    async getPosts(append = false) {
      try {
        let res = await this.getAllPost({
        offset: this.offset,
        limit: this.pagination.limit
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
      await this.getPosts(true)
      _.delay(() => done(), 1000)
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
      return this.$q.platform.is.desktop ? 'col-10' : 'col-12'
    }
  }
}
</script>
