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

const addVehicle = async () => {
  if (wincode.value) {
    try {
      await ApiService.addVehicle(wincode.value, 'Test Station')
      showSuccessMessage()
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
  <div class="p-4">
    <Card>
      <template #title>Add a vehicle</template>
      <template #content>
        <InputText type="text" placeholder="Enter wincode" v-model="wincode" />
      </template>
      <template #footer>
        <Button icon="pi pi-plus" label="Add" @click="addVehicle" />
      </template>
    </Card>
  </div>
</template>
