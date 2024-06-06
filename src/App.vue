<script setup>
import { RouterView, useRoute } from 'vue-router'
import Navbar from './components/home/Navbar.vue'

const route = useRoute()
const isRouteActive = (path) => {
  return route.path === path
}
</script>

<template>
  <Navbar v-if="!isRouteActive('/login')" />
  <RouterView />
</template>

<style scoped></style>
