<script setup>
import Card from 'primevue/card'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

import { useToast } from 'primevue/usetoast'
import Toast from 'primevue/toast'

import { ref } from 'vue'
import ApiService from '@/api'

const toast = useToast()

const showErrorMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Error finding vehicle`,
    life: 3000
  })
}

const showNotFoundMessage = () => {
  toast.add({
    severity: 'error',
    summary: 'Error',
    detail: `Vehicle not found`,
    life: 3000
  })
}

const vehicleDetails = ref(null)
const wincode = ref('')

const findVehicle = async () => {
  try {
    const data = await ApiService.findVehicle(wincode.value)

    if (data) {
      vehicleDetails.value = data
    } else {
      showNotFoundMessage()
    }
  } catch (error) {
    showErrorMessage()
  }
}
</script>

<template>
  <Toast />
  <div class="p-4">
    <Card>
      <template #title>Search for a vehicle</template>
      <template #content>
        <form>
          <InputText placeholder="Enter wincode" v-model="wincode" />
        </form>
      </template>
      <template #footer>
        <Button icon="pi pi-search" label="Search" @click="findVehicle" />
      </template>
    </Card>
  </div>
  <p v-if="vehicleDetails">
    {{ vehicleDetails }}
  </p>
</template>
