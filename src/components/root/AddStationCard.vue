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
    detail: 'Station successfully added',
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

const name = ref(null)
const region = ref(null)
const capacity = ref(null)

const addStation = async () => {
  try {
    await ApiService.addStation(
      name.value,
      region.value,
      capacity.value,
      localStorage.getItem('token')
    )
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  } finally {
    name.value = null
    region.value = null
    capacity.value = null
  }
}
</script>

<template>
  <Toast />
  <div class="p-4">
    <Card>
      <template #title>Add a station</template>
      <template #content>
        <div class="flex flex-wrap gap-2">
          <InputText type="text" placeholder="Enter station name" v-model="name" />
          <InputText type="text" placeholder="Enter region" v-model="region" />
          <InputText type="number" placeholder="Enter capacity" v-model="capacity" />
        </div>
      </template>
      <template #footer>
        <Button icon="pi pi-plus" label="Add" @click="addStation" />
      </template>
    </Card>
  </div>
</template>
