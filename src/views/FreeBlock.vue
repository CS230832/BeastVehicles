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

const slots = ref(Array.from({ length: 50 }, (_, index) => index))
const freeSlots = ref([])

const getFreeSlots = async () => {
  try {
    freeSlots.value = await ApiService.getFreeSlots('Test Station', props.blockName)
  } catch (error) {
    showErrorMessage()
  }
}

const slotIsFree = (slot) => {
  return freeSlots.value.some((item) => item.slot === slot)
}

onMounted(() => {
  getFreeSlots()
})
</script>

<template>
  <Toast />
  <div class="flex flex-wrap items-center gap-5 p-5">
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
