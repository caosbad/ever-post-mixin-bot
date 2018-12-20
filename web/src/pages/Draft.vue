<template>
  <q-page panding>
    <div class="row justify-center q-py-lg">
      <div class="col-lg-9 col-xs-11 col-md-10 form">
        <q-field
          class="col-8"
          :label="$t('TITLE')"
          :label-width="2"
          :error="$v.post.title.$error"
          :error-label="$t('VAL_TITLE_INVALID')"
          :count="256"
          :helper="$t('VAL_TITLE_POST')"
        >
          <q-input
            @blur="$v.post.title.$touch"
            v-model="post.title"
            n
            clearable
          />
        </q-field>
        <q-field
          class="col-8 q-mb-lg"
          :label="$t('NAME')"
          :label-width="2"
          :error="$v.post.name.$error"
          :count="20"
          :error-label="$t('VAL_URL_INVALID')"
          :helper="$t('VAL_URL_POST')"
        >
          <q-input
            @blur="$v.post.name.$touch"
            v-model="post.name"
            clearable
          />
        </q-field>
        <q-field
          :error="$v.post.content.$error"
          :error-label="$t('VAL_CONTENT_INVALID')"
        >
          <m-editor
            :language="locale"
            v-model="content"
            class="m-t-1"
            @change="change"
            :placeholder="$t('NEW_POST_PLACEHOLDER')"
          />
        </q-field>
        <div class="row justify-between">
          <q-btn
            :loading="loading"
            :disabled="$v.post.valid"
            class="col-auto m-t-1 self-center"
            color="primary"
            @click="post"
          >
            {{$t('SUBMIT')}}
          </q-btn>
          <!-- <q-btn
            rounded
            big
            class="col-auto m-t-1 self-center"
            @click="cancel"
          >
            {{$t('CANCEL')}}
          </q-btn> -->
        </div>
      </div>
    </div>
  </q-page>
</template>

<style>
</style>

<script>
import {
  QField,
  QInput,
  QRadio,
  QBtnToggle,
  QSlider,
  QPage,
  QBtn
} from 'quasar'
import {
  required,
  maxLength,
  minLength
} from 'vuelidate/lib/validators'
// import { mapActions, mapGetters } from 'vuex'
import { toastError } from '../utils/util'
import MEditor from '../components/MEditor'

export default {
  name: 'Draft',
  components: {
    QField,
    QInput,
    QRadio,
    QBtnToggle,
    QSlider,
    QPage,
    MEditor,
    QBtn
  },
  data() {
    return {
      post: {
        title: '',
        name: '',
        content: ''
      },
      postHtml: ''
    }
  },
  validations: {
    // TODO figure out why vuelidate can not bind vue 'this' in rule config
    post: {
      title: {
        required,
        maxLength: maxLength(256),
        minLength: minLength(5)
      },
      name: {
        maxLength: maxLength(20),
        minLength: minLength(3)
      },
      content: {
        required,
        minLength: minLength(10)
      }
    }
  },
  methods: {
     change(val, html) {
      this.post.content = val
      this.postHtml = html
      // this.$v.post.content.$touch()
    },
    post() {
      this.$v.post.touch()
      if (!this.$v.post.valid) {
        toastError(this.$t('ERROR'))
      }
      // let article = document.createElement('div')
      // article.innerHTML = this.postHtml
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill
    },
    locale() {
      let isZH = this.$i18n.locale === 'zh'
       return isZH ? 'zh-CN' : 'en'
    }
  }
}
</script>

<style lang="stylus">
.editor-container {
  font-family: CustomSerif, Georgia, Cambria, 'Times New Roman', serif;
  font-weight: 400;
  font-style: normal;
  font-size: 18px;
  line-height: 1.58;
  padding: 0;
  margin: 0;
  color: rgba(0, 0, 0, 0.8);
}

.title-input {
  font-family: CustomSansSerif, 'Lucida Grande', Arial, sans-serif;
  font-weight: 700;
  font-style: normal;
  font-size: 32px;
  margin: 21px 21px 12px;
  line-height: 34px;

  ::before {
    position: absolute;
    left: 0;
    right: 0;
    content: attr(data-placeholder);
    color: rgba(0, 0, 0, 0.44);
    font-weight: 400;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    pointer-events: none;
  }

  ::after {
    position: absolute;
    content: attr(data-label);
    color: rgba(0, 0, 0, 0.44);
    border-right: 1px solid rgba(0, 0, 0, 0.15);
    right: 100%;
    top: -6px;
    bottom: -6px;
    padding: 6px 12px;
    font-weight: 400;
    margin: 0 21px;
    opacity: 0;
    visibility: hidden;
    transition: opacity 0.15s ease;
  }
}

.name-input {
  font-family: CustomSansSerif, 'Lucida Grande', Arial, sans-serif;
  font-size: 15px;
  line-height: 18px;
  color: #79828B;
  margin: 12px 21px;
  font-style: normal;
}
</style>
