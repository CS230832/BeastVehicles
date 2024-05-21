<script setup>
import { onMounted } from 'vue'
import ApiService from '@/api'

const getUser = async () => {
  try {
    const response = await ApiService.getUser('nazariy', localStorage.getItem('token'))
    console.log(response.data)
  } catch (error) {
    console.log(error.response.data.data)
  }
}

onMounted(() => {
  getUser()
})
</script>

<template>
  <h1 class="text-xl font-semibold p-3">Welcome to Beast Vehicles!</h1>
</template>
