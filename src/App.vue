<script setup>
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import Navbar from './components/home/Navbar.vue'

const isAuthenticated = ref(true)
</script>

<template>
  <Navbar v-if="isAuthenticated" />
  <RouterView />
</template>

<style scoped></style>
