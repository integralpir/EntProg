import { createRouter, createWebHistory } from 'vue-router'
import ProductsView from '../views/Products.vue'
import AboutView from '../views/About.vue'
import HomeView from '../views/Home.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about', 
      name: 'about',
      component: AboutView
    },
    {
      path: '/products',
      name: 'products',
      component: ProductsView
    },
  ]
})

export default router