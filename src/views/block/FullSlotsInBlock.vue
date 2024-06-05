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

const data = ref({})
const slots = ref(Array.from({ length: 50 }, (_, index) => index + 1))

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

const getFullSlots = async () => {
  try {
    const response = await ApiService.getFullSlots(
      isRoot.value ? station.value : await getStationName(),
      props.blockName
    )
    data.value = response.data
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const isSlotFull = (slot) => {
  const slotsInBlock = data.value[props.blockName]
  return slotsInBlock && slotsInBlock.some((item) => item.slot === slot)
}

const getSlotWincode = (slot) => {
  const slotsInBlock = data.value[props.blockName]
  const slotData = slotsInBlock.find((item) => item.slot === slot)
  return slotData ? slotData.wincode : ''
}

onMounted(async () => {
  isRoot.value = await checkIfUserIsRoot()
  await getAllBlocks()
  if (isValidBlockName()) {
    await getFullSlots()
  } else {
    router.replace('/not-found')
  }
})
</script>

<template>
  <Toast />
  <div class="flex flex-wrap items-center gap-5 p-5">
    <Card
      v-for="slot in slots"
      :key="slot"
      :class="{ full: isSlotFull(slot), 'not-full': !isSlotFull(slot) }"
      class="w-40"
    >
      <template #content>
        <p class="text-center font-semibold">
          {{ isSlotFull(slot) ? getSlotWincode(slot) : slot }}
        </p>
      </template>
    </Card>
  </div>
</template>

<style scoped>
.full {
  background-color: white;
  color: #2d3748;
}

.not-full {
  background-color: #0f172a;
  color: #cfd1d4;
}
</style>
