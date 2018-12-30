<template>
  <q-card class="no-shadow bg-white">
    <q-item
      class="q-mt-sm"
      multiline
      @click.native="open(data.post_id)"
    >
      <q-item-side
        v-if="data.avatar_url"
        :avatar="data.avatar_url"
      />
      <q-item-main
        :label="data.title"
        label-lines="1"
        :sublabel="data.description"
        sublabel-lines="2"
      />
      <q-item-side
        class="desktop-only"
        right
      >
        <q-item-tile stamp>{{isDraft?data.created_at:data.updated_at | time}}</q-item-tile>
        <q-item-tile>
          <img
            v-if="data.path"
            class="icon-svg q-mx-xs"
            @click="openURL(data.telegraph_url)"
            :src="telegraphIcon"
          >
          <img
            v-if="data.ipfs_id"
            class="icon-svg q-mx-xs"
            :src="ipfsIcon"
            @click="openIPFSUrl(data)"
          >
        </q-item-tile>
      </q-item-side>
    </q-item>
    <q-card-main class="font-12 mobile-only">
      <div class="row justify-between">
        <div>
          <img
            v-if="data.path"
            class="icon-svg q-mx-xs"
            @click="openURL(data.telegraph_url)"
            :src="telegraphIcon"
          >
          <img
            v-if="data.ipfs_id"
            class="icon-svg q-mx-xs"
            :src="ipfsIcon"
            @click="openIPFSUrl(data)"
          >
        </div>
        <div>
          {{isDraft?data.created_at:data.updated_at | time}}
        </div>
      </div>

    </q-card-main>
  </q-card>
</template>

<script>
import { QCard, QCardMain, QItem, QItemMain, QItemSide, QCardActions, QItemTile, QCardSeparator, openURL } from 'quasar'
import { mapActions } from 'vuex'
// import { toast, translateErrMsg } from '../utils/util'
import telegraphIcon from '../assets/telegraph.svg'
import ipfsIcon from '../assets/ipfs.svg'
import {
  IPFS_GATEWAY
} from '../utils/constants'
export default {
  name: 'PostItem',
  props: ['data', 'index', 'type'],
  components: {
    QCard,
    QCardActions,
    QCardMain,
    QItem,
    QItemTile,
    QItemSide,
    QItemMain,
    QCardSeparator
  },
  data() {
    return {
      telegraphIcon,
      ipfsIcon
    }
  },
  methods: {
    ...mapActions(['dealShare']),
    openURL,
    open(id) {
      if (this.type === 'draft') {
        this.$router.push(`/post/${id}/edit`)
      } else {
        this.$router.push(`/post/${id}`)
      }
    },
    openIPFSUrl(post) {
      let url = post && post.ipfs_id ? `${IPFS_GATEWAY}${post.ipfs_id}/` : null
      openURL(url)
    },
    isPost() {
      return this.type === 'telegraph'
    },
    isDraft() {
      return this.type === 'draft'
    },
    isIPFS() {
      return this.type === 'ipfs'
    }
  },
  computed: {

  }
}
</script>

<style lang="stylus" scoped>
</style>

