<template>
  <q-card class="no-shadow bg-white">
    <q-item @click.native="open(data.post_id)">
      <q-item-side :avatar="data.avatar_url" />
      <q-item-main>
        <q-item-tile label>{{data.title}}</q-item-tile>
        <q-item-tile sublabel>{{data.created_at}}</q-item-tile>
      </q-item-main>
    </q-item>
    <q-card-main class="font-12 text-faded">
      {{data.description}}
      <div class="row justify-end">
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
          @click="openIPFSUrl"
        >
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
  props: ['data', 'index'],
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
      this.$router.push(`/post/${id}`)
    },
    openIPFSUrl() {
      let url = this.post && this.post.ipfs_id ? `${IPFS_GATEWAY}${this.post.ipfs_id}/` : null
      openURL(url)
    }
  },
  computed: {

  }
}
</script>

<style lang="stylus" scoped>
.icon-svg {
  width: 30px;
  height: 30px;
  border-radius: 15px;
  cursor: pointer;
}
</style>

