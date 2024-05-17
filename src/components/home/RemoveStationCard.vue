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
    detail: 'Station successfully removed',
    life: 3000
  })
}

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: errorMessage.value,
    life: 3000
  })
}

const name = ref(null)
const errorMessage = ref(null)

const removeStation = async () => {
  try {
    await ApiService.removeStation(name.value, localStorage.getItem('token'))
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error
    showErrorMessage()
  }
}
</script>

<template>
  <Toast />
  <div class="p-4 flex flex-col gap-5">
    <Card>
      <template #title>Remove station</template>
      <template #content>
        <InputText type="text" placeholder="Enter station name" v-model="name" />
      </template>
      <template #footer>
        <Button icon="pi pi-minus-circle" label="Remove" @click="removeStation" />
      </template>
    </Card>
  </div>
</template>
