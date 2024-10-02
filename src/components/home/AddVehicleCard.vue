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
    detail: 'Vehicle(s) successfully added',
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

const addMultipleVehicles = async () => {
  try {
    data.value = await ApiService.addMultipleVehicles(wincode.value, localStorage.getItem('token'))
    console.log(data.value)
    showSuccessMessage()
  } catch (error) {
    errorMessage.value = error.response.data.data
    showErrorMessage()
  }
}

const add = () => {
  if (typeof wincode.value === 'string') {
    addVehicle()
  } else if (Array.isArray(wincode.value)) {
    addMultipleVehicles()
  }

  wincode.value = null
}
</script>

<template>
  <Toast />
  <div class="p-4 flex flex-col gap-5">
    <Card>
      <template #title>Add a vehicle</template>
      <template #content>
        <Chips v-if="checked" v-model="wincode" separator="," />
        <InputText v-else type="text" placeholder="Enter wincode" v-model="wincode" />
        <div class="flex items-center gap-1 mt-4">
          <p class="text-sm">Multiple</p>
          <InputSwitch v-model="checked" @click="wincode = null" />
        </div>
      </template>
      <template #footer>
        <Button icon="pi pi-plus" label="Add" @click="add" />
      </template>
    </Card>
    <Card v-if="data">
      <template #title> Your vehicle's location </template>

      <template #content>
        <div class="flex items-center gap-8 flex-wrap">
          <div v-if="Array.isArray(data.data)" v-for="vehicle in data.data" :key="vehicle.wincode">
            <p class="font-bold">Wincode: {{ vehicle.wincode }}</p>
            <p class="font-semibold">Station: {{ vehicle.parking }}</p>
            <p class="font-semibold">Block: {{ vehicle.block }}</p>
            <p class="font-semibold">Slot: {{ vehicle.slot }}</p>
          </div>

          <div v-else>
            <p class="font-semibold">Station: {{ data.data.parking }}</p>
            <p class="font-semibold">Block: {{ data.data.block }}</p>
            <p class="font-semibold">Slot: {{ data.data.slot }}</p>
          </div>
        </div>
      </template>
    </Card>
  </div>
</template>
