<template>
  <q-infinite-scroll
    :inline="inline"
    v-if="getDatas.length"
    :handler="loadMore"
    :offset="10"
  >
    <q-list
      separator
      no-border
    >
      <slot
        name="item"
        v-for="(data, idx) in getDatas "
        :data="data"
        :index="idx"
      />
    </q-list>
    <div
      v-if="needLoad"
      class="row justify-center"
      style="margin-bottom: 50px;"
    >
      <q-spinner-dots
        slot="message"
        :size="40"
      />
    </div>
  </q-infinite-scroll>

  <div
    v-else
    class="no-comment-contaner row"
  >
    <!-- <img
      class="img q-mt-lg justify-center"
      :src="noCommentImg"
    > -->
    <p class="col-12 text-center q-mt-md">{{$t('table.noData')}}</p>
  </div>
</template>

<script>
import { QList, QListHeader, QItem, QSpinnerDots } from 'quasar'

export default {
  name: 'ListContainer',
  props: ['datas', 'expand', 'count', 'inline'],
  components: {
    QList,
    QListHeader,
    QItem,
    QSpinnerDots
  },
  data() {
    return {
    }
  },
  mounted() {},
  methods: {
    loadMore(index, done) {
      if (!this.needLoad) return
      this.$emit('loadMore', index, done)
    }
  },
  computed: {
    getDatas() {
      let datas = this.datas || []
      return datas
    },
    needLoad() {
      if (!this.expand) return false
      return this.getDatas.length < this.count
    }
  }
}
</script>

<style lang="stylus" scoped>
.scroll-container {
  height: 90vh;
}

.no-comment-contaner {
  margin-top: 0px;

  .img {
    width: 150px;
    height: 120px;
    margin: 20px auto;
  }

  p {
    font-size: 12px;
    color: #B3BFD0;
    letter-spacing: 0.5px;
  }
}
</style>