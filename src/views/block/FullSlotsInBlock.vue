<script setup>
import { onMounted, ref } from 'vue'
import Toast from 'primevue/toast'
import ApiService from '@/api'
import { useToast } from 'primevue/usetoast'
import Card from 'primevue/card'

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

const props = defineProps({
  blockName: String
})

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
    const response = await ApiService.getFullSlots(await getStationName(), props.blockName)
    data.value = response.data
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

onMounted(() => {
  getFullSlots()
})

const isSlotFull = (slot) => {
  const slotsInBlock = data.value[props.blockName]
  return slotsInBlock && slotsInBlock.some((item) => item.slot === slot)
}

const getSlotWincode = (slot) => {
  const slotsInBlock = data.value[props.blockName]
  const slotData = slotsInBlock.find((item) => item.slot === slot)
  return slotData ? slotData.wincode : ''
}
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
