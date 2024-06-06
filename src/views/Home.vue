<script setup>
import { onMounted, ref } from 'vue'
import ApiService from '@/api'
import checkIfUserIsAuthenticated from './auth/checkAuth'
import checkIfUserIsRoot from './auth/checkRoot'

const isAuth = ref(true)
const isRoot = ref(true)
const parkingName = ref(null)

const getParkingName = async () => {
  try {
    const response = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )
    parkingName.value = response.data.parking
  } catch (error) {
    console.log(`Error getting parking name: ${error.response.data.data}`)
  }
}

onMounted(async () => {
  isAuth.value = checkIfUserIsAuthenticated()
  isRoot.value = await checkIfUserIsRoot()
  if (isAuth.value && !isRoot.value) {
    await getParkingName()
  }
})
</script>

<template>
  <h1 class="text-xl font-semibold p-3">
    Welcome to Beast Vehicles!
    <span v-if="isAuth && !isRoot">Your parking station is "{{ parkingName }}".</span>
  </h1>
</template>

<style scoped></style>
