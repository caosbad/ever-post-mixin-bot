<template>
  <q-modal
    v-model="show"
    maximized
    no-esc-dismiss
    no-backdrop-dismiss
  >
    <q-modal-layout>
      <q-toolbar
        class="justify-start"
        slot="header"
        color="white"
        text-color="black"
      >
        <q-btn
          v-if="pageState==1"
          class="q-mx-sm"
          flat
          round
          dense
          icon="close"
          @click="$emit('close')"
        />
        <q-btn
          v-if="pageState==2"
          class="q-mx-sm"
          flat
          round
          dense
          icon="keyboard_backspace"
          @click="pageState=1"
        />
        <div
          v-show="pageState==1"
          class="col-11 justify-center"
        >
          <q-search
            class="col-11"
            v-model="search"
            color="blue-grey-1"
            hide-underline
            inverted-light
          />

        </div>
        <div
          v-show="pageState==2"
          class="col-11"
        >

        </div>
        <q-btn
          class="q-mx-sm"
          flat
          round
          dense
          icon="refresh"
          @click="refresh"
        />
      </q-toolbar>

      <!-- <q-toolbar
        v-show="pageState==1"
        slot="footer"
      >
        <q-toolbar-title>
          {{$t('TIPS')}}
        </q-toolbar-title>
      </q-toolbar> -->

      <div class="layout-padding">

        <q-list
          v-if="getAssets.length"
          highlight
          v-show="pageState==1"
        >
          <q-list-header>{{$t('CHOOSE_ASSET')}}</q-list-header>
          <q-item
            @click.native="()=>chooseAsset(asset)"
            v-for="asset in getAssets"
            :key="asset.asset_id"
          >
            <q-item-side>
              <q-item-tile avatar>
                <img :src="asset.icon_url">
              </q-item-tile>
            </q-item-side>
            <q-item-main
              :label="asset.balance "
              :sublabel="asset.symbol"
            />
            <!-- <q-item-side right>
              <q-btn icon />
            </q-item-side> -->
          </q-item>
        </q-list>
        <div v-else>
          <div class="row justify-center q-my-lg">
            <q-inner-loading
              :visible="!this.assets.length"
              color="info"
            />
          </div>
          <div class="row justify-center">
            <!-- <q-btn
              icon="refresh"
              @click="$emit('reload')"
              :label="$t('RELOAD_ASSETS')"
            /> -->
          </div>

        </div>
        <div v-show="pageState==2">
          <q-jumbotron
            v-if="selectAsset"
            class=""
          >
            <div class="">
              <div class="row justify-center">
                <img :src="selectAsset.icon_url">
              </div>
              <div class="q-display-3 row justify-center ">{{selectAsset.symbol}}</div>

              <hr class="q-hr q-my-lg">
              <div class="col-12">
                <q-field
                  :helper="$t('BALANCE')+' : ' + selectAsset.balance"
                  :error="valid"
                >
                  <q-input
                    @input="change"
                    color="white"
                    v-model="amount"
                    class="no-margin"
                  />
                </q-field>
              </div>
              <div class="row justify-between q-mt-lg">
                <q-btn
                  color="positive"
                  class="q-py-sm  co-3"
                  outline
                  :label="$t('PAY')"
                  @click="pay"
                />
                <q-btn
                  color="info"
                  class="q-py-sm  co-3"
                  outline
                  :label="$t('BACK')"
                  @click="()=>{pageState=1;reset()}"
                />
                <q-btn
                  color="negative"
                  class="q-py-sm  co-3"
                  outline
                  :label="$t('GAVEUP')"
                  @click="()=>{$emit('close');reset()}"
                />

              </div>

            </div>
          </q-jumbotron>
        </div>
      </div>
    </q-modal-layout>

  </q-modal>
</template>

<script>
import {
  QModalLayout,
  QToolbar,
  QBtn,
  QToolbarTitle,
  QSearch,
  QList,
  QListHeader,
  QItem,
  QItemMain,
  QItemSide,
  QItemTile,
  QJumbotron,
  QInput,
  QField,
  QInnerLoading
} from 'quasar'
const amountStrReg = /^\d+(\.\d+)?$/
export default {
  name: 'PaymentModal',
  props: ['show', 'assets'],
  components: {
    QModalLayout,
    QToolbar,
    QBtn,
    QToolbarTitle,
    QSearch,
    QList,
    QListHeader,
    QItem,
    QItemMain,
    QItemSide,
    QItemTile,
    QJumbotron,
    QInput,
    QField,
    QInnerLoading
  },
  data() {
    return {
      amountStrReg,
      selectAsset: undefined,
      pageState: 1,
      search: '',
      amount: '',
      valid: false
    }
  },
  mounted() {},
  methods: {
    chooseAsset(asset) {
      this.selectAsset = asset
      this.pageState = 2
    },
    change(amount) {
      this.valid = !amountStrReg.test(amount) || Number(amount) > this.selectAsset.balance
      console.log(this.valid)
    },
    pay() {
      if (this.valid) {
          return null
      }
      let asset = Object.assign({}, this.selectAsset, {amount: this.amount})
      this.$emit('pay', asset)
      this.reset()
    },
    reset() {
      this.selectAsset = undefined
      this.pageState = 1
      this.amount = ''
    },
    refresh() {
      this.$emit('reload')
    }
  },
  computed: {
    toolbars() {},
    getAssets() {
      if (!this.search) return this.assets
      return this.assets.filter(asset => {
        let symbol = asset.symbol.toLowerCase()
        let search = this.search.toLowerCase()
        return symbol.indexOf(search) >= 0
      })
    }
  }
}
</script>

<style lang="stylus">
#editorPro {
  /* margin: auto;
  height: 580px; */
  .markdown-body {
    min-width: 100px !important;
    width: 100% !important;
  }
}

.v-note-wrapper {
  min-width: 250px;
}
</style>
