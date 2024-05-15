import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/auth/Login.vue'
import Signup from '@/views/auth/Signup.vue'
import NotFound from '@/views/NotFound.vue'

import FreeBlock from '@/views/FreeBlock.vue'
import FullBlock from '@/views/FullBlock.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/home'
    },

    {
      path: '/home',
      name: 'home',
      component: Home,
      meta: { requiresAuth: true }
    },

    {
      path: '/search',
      name: 'search',
      component: () => import('@/components/home/SearchCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/add',
      name: 'add',
      component: () => import('@/components/home/AddCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/remove',
      name: 'remove',
      component: () => import('@/components/home/RemoveCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/free',
      name: 'free',
      component: () => import('@/views/Free.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/free/:blockName',
      name: 'free-block',
      component: FreeBlock,
      props: true,
      meta: { requiresAuth: true }
    },

    {
      path: '/full',
      name: 'full',
      component: () => import('@/views/Full.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/full/:blockName',
      name: 'full-block',
      component: FullBlock,
      props: true,
      meta: { requiresAuth: true }
    },

    {
      path: '/login',
      name: 'login',
      component: Login,
      meta: { requiresGuest: true }
    },

    {
      path: '/signup',
      name: 'signup',
      component: Signup,
      meta: { requiresGuest: true }
    },

    {
      path: '/:catchAll(.*)',
      name: 'not-found',
      component: NotFound
    }
  ]
})

const checkIfUserIsAuthenticated = () => {
  return true
}

const isAuthenticated = checkIfUserIsAuthenticated()

router.beforeEach((to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!isAuthenticated) {
      next({ name: 'login' })
    } else {
      next()
    }
  } else if (to.matched.some((record) => record.meta.requiresGuest)) {
    if (isAuthenticated) {
      next({ name: 'home' })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
