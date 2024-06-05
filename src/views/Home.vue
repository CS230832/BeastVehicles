<script setup>
import { onMounted, ref } from 'vue'
import ApiService from '@/api'
import checkIfUserIsRoot from './auth/checkRoot'
import ThreeDModel from '@/components/home/ThreeDModel.vue'

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
    console.log(error.response.data.data)
  }
}

onMounted(async () => {
  isRoot.value = await checkIfUserIsRoot()
  if (!isRoot.value) {
    await getParkingName()
  }
})
</script>

<template>
  <h1 class="text-xl font-semibold p-3">
    Welcome to Beast Vehicles!
    <span v-if="!isRoot">Your parking station is "{{ parkingName }}".</span>
  </h1>
  <div>
    <ThreeDModel />
  </div>
</template>

<style scoped></style>
