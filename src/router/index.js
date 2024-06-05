import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import Login from '@/views/auth/Login.vue'
import NotFound from '@/views/NotFound.vue'
import FreeSlotsInBlock from '@/views/block/FreeSlotsInBlock.vue'
import FullSlotsInBlock from '@/views/block/FullSlotsInBlock.vue'
import checkIfUserIsAuthenticated from '@/views/auth/checkAuth'
import checkIfUserIsRoot from '@/views/auth/checkRoot'
import checkIfUserIsManager from '@/views/auth/checkManager'

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
      component: () => import('@/components/home/SearchVehicleCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/add',
      name: 'add',
      component: () => import('@/components/home/AddVehicleCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/add-station',
      name: 'add-station',
      component: () => import('@/components/root/AddStationCard.vue'),
      meta: { requiresAuth: true, requiresRoot: true }
    },

    {
      path: '/remove-station',
      name: 'remove-station',
      component: () => import('@/components/root/RemoveStationCard.vue'),
      meta: { requiresAuth: true, requiresRoot: true }
    },

    {
      path: '/add-user',
      name: 'add-user',
      component: () => import('@/components/root/AddUserCard.vue'),
      meta: { requiresAuth: true, requiresManagerOrRoot: true }
    },

    {
      path: '/remove-user',
      name: 'remove-user',
      component: () => import('@/components/root/RemoveUserCard.vue'),
      meta: { requiresAuth: true, requiresManagerOrRoot: true }
    },

    {
      path: '/remove',
      name: 'remove',
      component: () => import('@/components/home/RemoveVehicleCard.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/free',
      name: 'free',
      component: () => import('@/views/station/FreeSlotsInStation.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/free/:blockName',
      name: 'free-block',
      component: FreeSlotsInBlock,
      props: true,
      meta: { requiresAuth: true }
    },

    {
      path: '/full',
      name: 'full',
      component: () => import('@/views/station/FullSlotsInStation.vue'),
      meta: { requiresAuth: true }
    },

    {
      path: '/full/:blockName',
      name: 'full-block',
      component: FullSlotsInBlock,
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
      path: '/:catchAll(.*)',
      name: 'not-found',
      component: NotFound
    }
  ]
})

const isAuthenticated = checkIfUserIsAuthenticated()

router.beforeEach(async (to, from, next) => {
  if (to.matched.some((record) => record.meta.requiresAuth)) {
    if (!isAuthenticated) {
      next({ name: 'login' })
    } else {
      const isRoot = await checkIfUserIsRoot()
      if (to.matched.some((record) => record.meta.requiresRoot)) {
        if (isRoot) {
          next()
        } else {
          next({ name: 'home' })
        }
      } else if (to.matched.some((record) => record.meta.requiresManagerOrRoot)) {
        const isManager = await checkIfUserIsManager()
        if (isRoot || isManager) {
          next()
        } else {
          next({ name: 'home' })
        }
      } else {
        next()
      }
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
