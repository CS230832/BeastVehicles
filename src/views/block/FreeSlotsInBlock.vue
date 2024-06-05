<script setup>
import { onMounted, ref, computed } from 'vue'
import ApiService from '@/api'
import checkIfUserIsRoot from '../auth/checkRoot'
import { useStationStore } from '@/stores'
import { useRouter } from 'vue-router'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import { useToast } from 'primevue/usetoast'

const isRoot = ref(null)
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

const validBlockNames = ref(null)

const getAllBlocks = async () => {
  try {
    const response = await ApiService.getAllBlocks(
      isRoot.value ? station.value : await getStationName()
    )
    validBlockNames.value = response.data
  } catch (error) {
    console.log(`Error getting all blocks: ${error.response.data.data}`)
  }
}

const props = defineProps({
  blockName: String
})

const isValidBlockName = () => {
  return validBlockNames.value.hasOwnProperty.call(validBlockNames.value, props.blockName)
}

const slots = ref(Array.from({ length: 50 }, (_, index) => index + 1))
const freeSlots = ref([])

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

const getFreeSlots = async () => {
  try {
    const response = await ApiService.getFreeSlots(
      isRoot.value ? station.value : await getStationName(),
      props.blockName
    )
    freeSlots.value = response.data[props.blockName] || []
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const slotIsFree = (slot) => {
  return freeSlots.value.some((item) => item.slot === slot)
}

onMounted(async () => {
  isRoot.value = await checkIfUserIsRoot()
  await getAllBlocks()
  if (isValidBlockName()) {
    await getFreeSlots()
  } else {
    router.replace('/not-found')
  }
})
</script>

<template>
  <Toast />
  <div class="flex flex-wrap items-center gap-5 p-5" v-if="validBlockNames">
    <Card
      v-for="slot in slots"
      :key="slot"
      :class="{ free: slotIsFree(slot), 'not-free': !slotIsFree(slot) }"
      class="min-w-20"
    >
      <template #content>
        <p class="text-center font-semibold">{{ slot }}</p>
      </template>
    </Card>
  </div>
</template>

<style scoped>
.free {
  background-color: white;
  color: #2d3748;
}

.not-free {
  background-color: #0f172a;
  color: #cfd1d4;
}
</style>
