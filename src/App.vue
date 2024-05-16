<script setup>
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import Navbar from './components/home/Navbar.vue'

import checkIfUserIsAuthenticated from './views/auth/checkAuth'

const isAuthenticated = ref(checkIfUserIsAuthenticated())
</script>

<template>
  <Navbar v-if="isAuthenticated" />
  <RouterView />
</template>

<style scoped></style>
