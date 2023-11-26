<template>
  <li>
    <strong>{{index + 1}}</strong>
    {{ note.title }}
    <button class="rm"
            v-on:click="$emit('remove-note', note.id)"
    >&times;</button>
  </li>
</template>

<script>
export default {
  props: {
    note: {
      type: Object,
      required: true
    },
    index: Number,
  },
}
</script>

<style scoped>
li {
  border: 1px solid #ccc;
  display: flex;
  justify-content: space-between;
  padding: .5rem 2rem;
  margin-bottom: 1rem;
}

.rm {
  background: red;
  color: #fff;
  border-radius: 50%;
  font-weight: bold;
}

input {
  margin-right: 1rem;
}
</style>