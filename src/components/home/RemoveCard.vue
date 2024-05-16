<script setup>
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ref } from 'vue'
import ApiService from '@/api'
import checkIfUserIsAuthenticated from '@/views/auth/checkAuth'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'Vehicle successfully removed',
    life: 3000
  })
}

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Error removing vehicle`,
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

const removeVehicle = async () => {
  if (wincode.value && checkIfUserIsAuthenticated()) {
    try {
      await ApiService.removeVehicle(wincode.value, localStorage.getItem('token'))
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
      <template #title>Remove a vehicle</template>
      <template #content>
        <InputText type="text" placeholder="Enter wincode" v-model="wincode" />
      </template>
      <template #footer>
        <Button icon="pi pi-trash" label="Remove" severity="danger" @click="removeVehicle" />
      </template>
    </Card>
  </div>
</template>
