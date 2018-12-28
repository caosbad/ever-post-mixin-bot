<template>
  <div class="row separator">
  </div>
</template>

<script>
export default {
  name: 'Separator',
  props: [''],
  components: {},
  data() {
    return {}
  },
  methods: {
    goBack() {
      this.$router.go(-1)
    }
  }
}
</script>

<style lang="stylus" scoped>
.separator {
  height: 10px;
  background-color: #f2f5fb;
}
</style>