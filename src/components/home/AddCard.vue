<script setup>
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ref } from 'vue'
import ApiService from '@/api'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'Vehicle successfully added',
    life: 3000
  })
}

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Error adding vehicle`,
    life: 3000
  })
}

const showEmptyWincode = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Wincode cannot be empty`,
    life: 3000
  })
}

const wincode = ref('')
const data = ref(null)

const addVehicle = async () => {
  if (wincode.value) {
    try {
      data.value = await ApiService.addVehicle(
        wincode.value,
        'Test Station',
        localStorage.getItem('token')
      )
      showSuccessMessage()
      console.log(data.value)
    } catch (error) {
      showErrorMessage()
    }
  } else {
    showEmptyWincode()
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
        <p class="font-semibold">Station: {{ data.parking }}</p>
        <p class="font-semibold">Block: {{ data.block }}</p>
        <p class="font-semibold">Slot: {{ data.slot }}</p>
      </template>
    </Card>
  </div>
</template>
