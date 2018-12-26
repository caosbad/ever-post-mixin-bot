// This is just an example,
// so you can safely delete all default props below

const lang = {
  lang: 'en',
  label: {
    clear: 'Clear',
    ok: 'Confirm',
    cancel: 'Cancel',
    close: 'Close',
    set: 'Set',
    select: 'Select',
    reset: 'Reset',
    remove: 'Remove',
    update: 'Update',
    create: 'Create',
    search: 'Search',
    filter: 'Filter',
    refresh: 'Refresh'
  },
  date: {
    days: 'SUNDAY_MONDAY_TUESDAY_WEDNESDAY_THURSDAY_FRIDAY_SATURDAY'.split('_'),
    daysShort: 'SUN_MON_TUE_WED_THU_FRI_SAT'.split('_'),
    months: 'JAN_FEB_MAR_APR_MAY_JUN_JUL_AUG_SEP_OCT_NOV_DEC'.split('_'),
    monthsShort: '1_2_3_4_5_6_7_8_9_10_11_12'.split('_'),
    firstDayOfWeek: 0, // 0-6, 0 - Sunday, 1 Monday, ...
    format24h: false
  },
  pullToRefresh: {
    pull: 'Pull to refresh',
    release: 'Release',
    refresh: 'Refreshing...'
  },
  table: {
    noData: 'No data',
    noResults: 'Not marched',
    loading: 'loading...',
    selectedRows: rows =>
      rows > 1 ? `${rows} selected row(s).` : `${rows === 0 ? 'No' : '1'} selected rows.`,
    rowsPerPage: '每页的行数:',
    allRows: '全部',
    pagination: (start, end, total) => `${start}-${end} of ${total}`,
    columns: '列'
  },
  editor: {
    url: '地址',
    bold: '粗体',
    italic: '斜体',
    strikethrough: '删除线',
    underline: '下划线',
    unorderedList: '无序列表',
    orderedList: '有序列表',
    subscript: '子脚本',
    superscript: '超级脚本',
    hyperlink: '超链接',
    toggleFullscreen: '全屏切换',
    quote: '引号',
    left: '左对齐',
    center: '居中对齐',
    right: '右对齐',
    justify: '两端对齐',
    print: '打印',
    outdent: '减少缩进',
    indent: '增加缩进',
    removeFormat: '清除样式',
    formatting: '格式化',
    fontSize: '字体大小',
    align: '对齐',
    hr: '插入水平线',
    undo: '撤消',
    redo: '重做',
    header1: '标题一',
    header2: '标题二',
    header3: '标题三',
    header4: '标题四',
    header5: '标题五',
    header6: '标题六',
    paragraph: '段落',
    code: '代码',
    size1: '非常小',
    size2: '比较小',
    size3: '正常',
    size4: '中等偏大',
    size5: '大',
    size6: '非常大',
    size7: '超级大',
    defaultFont: '默认字体'
  },

  // ====
  'EDITOR_PLACEHOLDER': 'Compose an epic...',
  'TITLE': 'Title',
  'YOURNAME': 'Your name'
}
export default lang
