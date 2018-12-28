<template>
  <q-page padding>
    <div :class="maxWidthClass">
      <list-container
        :datas="posts"
        @loadMore="loadMore"
        :expand="true"
      >
        <template
          slot="item"
          slot-scope="props"
        >
          <separator />
          <post-item
            :data="props.data"
            :index="props.index"
          />
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
        limit: 10
      },
      posts: []
    }
  },
  mounted() {
    this.getPosts()
  },
  methods: {
    ...mapActions(['getAllPost', 'getDrafts', 'getMyPosts']),
    async getPosts() {
      let res = await this.getAllPost({
        offset: this.offset,
        limit: this.pagination.limit
      })
      if (res) {
        console.log(res)
        this.posts = res
      }
    },
    open(postId) {
      this.$router.push('/')
    },
    loadMore() {

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
