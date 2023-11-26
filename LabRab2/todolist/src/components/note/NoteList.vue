<template>
  <div>
    <ul>
      <NoteItem
          v-for="(note, i) of notes"
          v-bind:note="note"
          v-bind:index="i"
          v-on:remove-note="removeNote"
      />
    </ul>
  </div>
</template>

<script>
import NoteItem from "@/components/note/NoteItem.vue";
export default {
  props: ['notes'],
  components: {
    NoteItem,
  },
  methods: {
    removeNote(id) {
      this.$emit('remove-note', id)
    }
  }
}
</script>

<style scoped>
ul {
  list-style: none;
  margin: 0;
  padding: 0;
}
</style>