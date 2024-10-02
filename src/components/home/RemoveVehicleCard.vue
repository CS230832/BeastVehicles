<script setup>
import { ref } from 'vue'
import ApiService from '@/api'

import Toast from 'primevue/toast'
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Chips from 'primevue/chips'
import InputSwitch from 'primevue/inputswitch'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'

const toast = useToast()

const showSuccessMessage = () => {
  toast.add({
    severity: 'success',
    summary: 'Success',
    detail: 'Vehicle successfully removed',
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

const checked = ref(false)
const wincode = ref(null)

const removeVehicle = async () => {
  try {
    await ApiService.removeVehicle(wincode.value, localStorage.getItem('token'))
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const removeMultipleVehicles = async () => {
  try {
    await ApiService.removeMultipleVehicles(wincode.value, localStorage.getItem('token'))
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const remove = () => {
  if (typeof wincode.value === 'string') {
    removeVehicle()
  } else if (Array.isArray(wincode.value)) {
    removeMultipleVehicles()
  }

  wincode.value = null
}
</script>

<template>
  <Toast />
  <div class="p-4">
    <Card>
      <template #title>Remove a vehicle</template>
      <template #content>
        <Chips v-if="checked" v-model="wincode" separator="," />
        <InputText v-else type="text" placeholder="Enter wincode" v-model="wincode" />
        <div class="flex items-center gap-1 mt-4">
          <p class="text-sm">Multiple</p>
          <InputSwitch v-model="checked" @click="wincode = null" />
        </div>
      </template>
      <template #footer>
        <Button icon="pi pi-trash" label="Remove" severity="danger" @click="remove" />
      </template>
    </Card>
  </div>
</template>
