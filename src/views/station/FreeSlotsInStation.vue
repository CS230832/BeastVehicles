<script setup>
import { onMounted, ref, computed } from 'vue'
import ApiService from '@/api'
import { useRouter } from 'vue-router'

import checkIfUserIsRoot from '../auth/checkRoot'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import { useToast } from 'primevue/usetoast'

import RootInput from './RootInput.vue'
import { useStationStore } from '@/stores'

const store = useStationStore()
const station = computed(() => store.station)

const router = useRouter()
const toast = useToast()

const errorMessage = ref(null)
const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: errorMessage.value,
    life: 3000
  })
}

const getStationName = async () => {
  try {
    const response = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )

    return response.data.parking
  } catch (error) {
    console.log(`Error getting station name: ${error.response.data.data}`)
  }
}

const data = ref(null)

const getFreeBlocks = async () => {
  try {
    const response = await ApiService.getFreeBlocks(await getStationName())
    data.value = response.data
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const navigateToBlock = (blockName) => {
  router.push({ name: 'free-block', params: { blockName } })
}

const isRoot = ref(null)

onMounted(async () => {
  isRoot.value = await checkIfUserIsRoot()
  getFreeBlocks()
})
</script>

<template>
  <Toast />
  <div v-if="isRoot" class="px-10 mt-10">
    <RootInput />
    {{ station }}
  </div>
  <div class="flex gap-10 flex-wrap items-center p-10" v-if="data">
    <Card
      v-for="(slots, blockName) in data"
      :key="blockName"
      class="min-w-64 cursor-pointer"
      @click="navigateToBlock(blockName)"
    >
      <template #title
        ><h1 class="text-center text-9xl">
          {{ blockName }}
        </h1></template
      >
      <template #content>
        <p class="text-center font-semibold">{{ slots.length }}/50</p>
      </template>
    </Card>
  </div>
</template>
