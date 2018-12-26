<template>
  <q-page padding>
    <div :class="maxWidthClass">
      <!-- <div>
        {{$t()}}
      </div> -->
      <q-list
        class=""
        no-border
        inset-separator
      >
        <q-list-header>{{$t('RECENT_POSTS')}}</q-list-header>
        <q-item
          v-if="posts.length"
          v-for="(post,idx) in posts"
          :key="idx"
          :to="'/post/'+ post.post_id"
        >
          <q-item-side>
            <q-item-tile avatar>
              <!-- <img :src="" /> -->
            </q-item-tile>
          </q-item-side>
          <q-item-main>
            <q-item-tile label>{{post.title}}</q-item-tile>
            <q-item-tile sublabel>{{post.description}}</q-item-tile>
          </q-item-main>
          <q-item-side right>
            <q-item-tile
              icon="chat_bubble"
              color="green"
            />
          </q-item-side>
        </q-item>
        <q-item v-else>
          {{$t('NO_DATA')}}
          <q-btn
            label="New"
            @click="$router.push('/draft')"
          />
        </q-item>
      </q-list>
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
  QBtn
} from 'quasar'
import {
  mapActions
} from 'vuex'

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
    QBtn
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
