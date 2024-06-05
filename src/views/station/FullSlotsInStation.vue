<script setup>
import { onMounted, ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import ApiService from '@/api'
import checkIfUserIsRoot from '../auth/checkRoot'
import { useStationStore } from '@/stores'
import Toast from 'primevue/toast'
import RootInput from '@/components/root/RootInput.vue'
import Card from 'primevue/card'
import { useToast } from 'primevue/usetoast'

const router = useRouter()
const isRoot = ref(null)
const store = useStationStore()
const station = computed(() => store.station)
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

const getFullBlocks = async () => {
  try {
    const response = await ApiService.getFullBlocks(
      isRoot.value ? station.value : await getStationName()
    )
    data.value = response.data
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
    data.value = null
  }
}

const navigateToBlock = (blockName) => {
  router.push({ name: 'full-block', params: { blockName } })
}

onMounted(async () => {
  isRoot.value = await checkIfUserIsRoot()
  if (isRoot.value) {
    watch(
      station,
      async (newStation) => {
        if (newStation) {
          await getFullBlocks()
        }
      },
      { immediate: true }
    )
  } else {
    await getFullBlocks()
  }
})
</script>

<template>
  <Toast />
  <div v-if="isRoot" class="px-10 mt-10">
    <RootInput />
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
        <p class="text-center font-semibold">{{ slots ? slots.length : 0 }}/50</p>
      </template>
    </Card>
  </div>
</template>
