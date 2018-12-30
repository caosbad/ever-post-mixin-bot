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
  // ====
  'NO_DATA': 'no data',
  'EDITOR_PLACEHOLDER': 'Write your story',
  'TITLE': 'Title',
  'YOURNAME': 'Your name',
  'TIPS': 'Confirm your asset and amount befor your pay.',
  'CHOOSE_ASSET': 'Choose your asset',
  'BALANCE': 'balance',
  'PAY': 'Pay',
  'BACK': 'Back',
  'GAVEUP': 'Giveup',
  'NO_ASSETS': 'no assets avaliable',
  'NAME': 'Name',
  'DESC': 'Descriptionn',
  'NEW_POST_PLACEHOLDER': 'Write your story',
  'SAVE': 'Save',
  'REPUBLISH': 'Publish again',
  'CANCEL': 'Cancel',
  'DELETE': 'Delete',
  'CANNOT_SAVE_IPFS_POST': 'Can not save to IPFS, try it later.',
  'IPFS_ERROR': 'Send page to IPFS fail,try it later.',
  'SUCCESS': 'Success',
  'ERROR': 'error',
  'FORM_INVALID': 'Post after fill content',
  'RECENT_POSTS': 'Recen posts',
  'LOAD_ERROR': 'Load error, try it later.',
  'EDIT': 'Edit',
  'PUBLISH': 'Publish',
  'VIEW_ON_TELEGRAPH': 'View on Telegraph',
  'IPFS_IT': 'Send to IPFS',
  'VIEW_ON_IPFS': 'View on IPFS',
  'UNFOLLOW': 'unFollow',
  'FOLLOW': 'Follow this author',
  'DONATE_AUTHOR': 'Donate author',
  'DRAFTS': 'Drafts',
  'TELEGRAPH': 'Pots',
  'IPFS': 'Saved',
  'PREVIEW': 'Preview & Publish'
}
export default lang
