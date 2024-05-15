<script setup>
import { onMounted, ref } from 'vue'
import Toast from 'primevue/toast'
import ApiService from '@/api'
import { useToast } from 'primevue/usetoast'
import Card from 'primevue/card'

const toast = useToast()

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Error fetching slots`,
    life: 3000
  })
}

const props = defineProps({
  blockName: String
})

const data = ref([])
const slots = ref(Array.from({ length: 50 }, (_, index) => index))

const getFullSlots = async () => {
  try {
    data.value = await ApiService.getFullSlots('Test Station', props.blockName)
  } catch (error) {
    showErrorMessage()
  }
}

onMounted(() => {
  getFullSlots()
})

const isSlotFull = (slot) => {
  return data.value.some((item) => item.slot === slot)
}

const getSlotWincode = (slot) => {
  const slotData = data.value.find((item) => item.slot === slot)
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
      class="w-24"
    >
      <template #content>
        <p class="text-center font-semibold">
          {{ isSlotFull(slot) ? getSlotWincode(slot) : 'free' }}
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
