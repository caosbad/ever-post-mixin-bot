/* eslint-disable */
import _ from 'lodash'

const filters = {
  t: (value, t) => {
    return t(value)
  },
  time: (value) => {
    return value
  }
}
export default filters
