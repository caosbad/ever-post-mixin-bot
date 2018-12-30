/* eslint-disable */
import _ from 'lodash'
import moment from 'moment'

const filters = {
  t: (value, t) => {
    return t(value)
  },
  time: (value) => {
    return moment(value).format('YYYY-MM-DD  hh:mm:ss')
  }
}
export default filters
