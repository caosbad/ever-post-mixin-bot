<template>
  <div id="editorPro">
    <mavon-editor
      :value="value"
      :placeholder="placeholder"
      @change="change"
      @save="save"
      :subfield="false"
      :defaultOpen="defaultOpen"
      :toolbarsFlag="toolbarsFlag"
      :editable="editable"
      :toolbars="toolbars"
      :navigation="navigation"
      style="height: 100%"
    />
  </div>
</template>

<script>
// Local Registration
import { mavonEditor } from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
export default {
  name: 'MEditor',
  props: ['editable', 'placeholder', 'value', 'toolbarsFlag', 'defaultOpen', 'toolbarsConfig', 'navigation'],
  components: {
    mavonEditor
    // or 'mavon-editor': mavonEditor
  },
  data() {
    return {
      input: ''
    }
  },
  mounted() {},
  methods: {
    change(val, html) {
      this.$emit('change', val, html)
    },
    save(val, html) {
      this.$emit('save', val, html)
    }
  },
  computed: {
    toolbars() {
      let isMobile = this.$q.platform.is.mobile
      if (this.toolbarsConfig) return this.toolbarsConfig
      if (!isMobile) {
        return {
          bold: true, // 粗体
          italic: true, // 斜体
          header: true, // 标题
          // underline: true, // 下划线
          // strikethrough: true, // 中划线
          mark: true, // 标记
          // superscript: true, // 上角标
          // subscript: true, // 下角标
          quote: true, // 引用
          ol: true, // 有序列表
          ul: true, // 无序列表
          link: true, // 链接
          // imagelink: true, // 图片链接
          code: true, // code
          table: true, // 表格
          fullscreen: true, // 全屏编辑
          readmodel: true, // 沉浸式阅读
          // htmlcode: true, // 展示html源码
          // help: true, // 帮助
          // undo: true, // 上一步
          // redo: true, // 下一步
          // trash: true, // 清空
          save: true, // 保存（触发events中的save事件）
          // navigation: true, // 导航目录
          // alignleft: true, // 左对齐
          // aligncenter: true, // 居中
          // alignright: true, // 右对齐
          subfield: true, // 单双栏模式
          preview: true // 预览
        }
      } else {
        return {
          bold: true, // 粗体
          italic: true, // 斜体
          header: true // 标题
        }
      }
    }
  }
}
</script>

<style lang="stylus" >
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