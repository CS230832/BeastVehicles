<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'Vehicle successfully added',
    life: 3000
  })
}

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

const addVehicle = async () => {
  try {
    data.value = await ApiService.addVehicle(wincode.value, localStorage.getItem('token'))
    showSuccessMessage()
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
      <template #title>Add a vehicle</template>
      <template #content>
        <InputText type="text" placeholder="Enter wincode" v-model="wincode" />
      </template>
      <template #footer>
        <Button icon="pi pi-plus" label="Add" @click="addVehicle" />
      </template>
    </Card>
    <Card v-if="data">
      <template #title> Your vehicle's location </template>

      <template #content>
        <p class="font-semibold">Station: {{ data.data.parking }}</p>
        <p class="font-semibold">Block: {{ data.data.block }}</p>
        <p class="font-semibold">Slot: {{ data.data.slot }}</p>
      </template>
    </Card>
  </div>
</template>
