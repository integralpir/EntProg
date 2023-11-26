<template>
  <div class="row">
    <div class="col-md-4">
      <h2>Список дел</h2>
      <hr>
      <AddTodo
          @add-todo="addTodo"
      />
      <select v-model="filter">
        <option value="all">All</option>
        <option value="completed">Completed</option>
        <option value="not-completed">Not Completed</option>
      </select>
      <hr>
      <TodoList
          v-bind:todos="filteredTodos"
          @remove-todo="removeTodo"
      />
    </div>

    <div class="col-md-4">
      <h2>Заметки</h2>
      <hr>
      <AddNote
          @add-note="addNote"
      />
      <hr>
      <NoteList
          @remove-note="removeNote"
      />
    </div>
  </div>
</template>

<script>
import TodoList from '@/components/todo/TodoList.vue'
import AddTodo from "@/components/todo/AddTodo.vue";
import NoteList from "@/components/note/NoteList.vue";
import AddNote from "@/components/note/AddNote.vue";
  export default {
    name: 'app',
    data() {
      return {
        todos: [],
        notes: [],
        loading: true,
        filter: 'all'
      }
    },
    computed: {
      filteredTodos() {
        if (this.filter === 'all') {
          return this.todos
        }

        if (this.filter === 'completed') {
          return this.todos.filter(t => t.completed)
        }

        if (this.filter === 'not-completed') {
          return this.todos.filter(t => !t.completed)
        }
      }
    },
    methods: {
      removeTodo(id) {
        this.todos = this.todos.filter(t => t.id !== id)
      },
      addTodo(todo) {
        this.todos.push(todo)
      },
      removeNote(id) {
        this.notes = this.notes.filter(n => n.id !== id)
      },
      addNote(note) {
        this.notes.push(note)
      }
    },
    components: {
      NoteList, AddNote,
      TodoList, AddTodo
    }
  }
</script>

<style>

</style>
