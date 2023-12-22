import { createRouter, createWebHistory } from 'vue-router'
import Notes from "@/views/Notes.vue";
import AboutView from '../views/About.vue'
import RateView from '../views/Rate.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'notes',
      component: Notes
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView
    },
    {
      path: '/rate',
      name: 'rate',
      component: RateView
    },
  ]
})

export default router
