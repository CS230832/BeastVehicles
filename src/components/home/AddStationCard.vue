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
    detail: 'Station successfully created',
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
const region = ref(null)
const capacity = ref(null)
const errorMessage = ref(null)

const createStation = async () => {
  try {
    await ApiService.createStation(
      name.value,
      region.value,
      capacity.value,
      localStorage.getItem('token')
    )
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
      <template #title>Create a station</template>
      <template #content>
        <div class="flex gap-2">
          <InputText type="text" placeholder="Enter station name" v-model="name" />
          <InputText type="text" placeholder="Enter region" v-model="region" />
          <InputText type="number" placeholder="Enter capacity" v-model="capacity" />
        </div>
      </template>
      <template #footer>
        <Button icon="pi pi-car" label="Create" @click="createStation" />
      </template>
    </Card>
  </div>
</template>
