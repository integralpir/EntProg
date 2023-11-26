<template>
  <form @submit.prevent="onSubmit">
    <input type="text" v-model="title">
    <button type="submit">Create</button>
  </form>
</template>

<script>
export default {
  data() {
    return {
      title: ''
    }
  },
  methods: {
    onSubmit() {
      if (this.title.trim()) {
        const newNote = {
          id: Date.now(),
          title: this.title
        }

        this.$emit('add-note', newNote)
        this.title = ''
      }
    }
  }
}
</script>

<style scoped>

</style>