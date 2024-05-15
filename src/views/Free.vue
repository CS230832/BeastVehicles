<script setup>
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'
import Card from 'primevue/card'

import { onMounted, ref } from 'vue'
import ApiService from '@/api'
import { useRouter } from 'vue-router'

const router = useRouter()
const toast = useToast()

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Error fetching free slots`,
    life: 3000
  })
}

const data = ref(null)

const getFreeBlocks = async () => {
  try {
    data.value = await ApiService.getFreeBlocks('Test Station')
  } catch (error) {
    showErrorMessage()
  }
}

const navigateToBlock = (blockName) => {
  router.push({ name: 'free-block', params: { blockName } })
}

onMounted(() => {
  getFreeBlocks()
})
</script>

<template>
  <Toast />
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
