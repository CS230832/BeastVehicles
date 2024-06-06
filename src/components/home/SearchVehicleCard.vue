<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

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

const data = ref(null)
const wincode = ref(null)

const findVehicle = async () => {
  try {
    data.value = await ApiService.findVehicle(wincode.value)
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}
</script>

<template>
  <Toast />
  <div class="p-4 flex flex-col gap-5">
    <Card>
      <template #title>Search for a vehicle</template>
      <template #content>
        <InputText placeholder="Enter wincode" v-model="wincode" />
      </template>
      <template #footer>
        <Button icon="pi pi-search" label="Search" @click="findVehicle" />
      </template>
    </Card>

    <Card v-if="data">
      <template #title> Vehicle's location </template>

      <template #content>
        <p class="font-semibold">Station: {{ data.data.parking }}</p>
        <p class="font-semibold">Block: {{ data.data.block }}</p>
        <p class="font-semibold">Slot: {{ data.data.slot }}</p>
      </template>
    </Card>
  </div>
</template>
