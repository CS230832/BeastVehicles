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

const data = ref(null)
const wincode = ref('')

const findVehicle = async () => {
  try {
    data.value = await ApiService.findVehicle(wincode.value) // I am here
  } catch (error) {
    showErrorMessage()
  }
}
</script>

<template>
  <Toast />
  <div class="p-4 flex flex-col gap-5">
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

    <Card v-if="data">
      <template #title> Vehicle's location </template>

      <template #content>
        <p class="font-semibold">Station: {{ data.parking }}</p>
        <p class="font-semibold">Block: {{ data.block }}</p>
        <p class="font-semibold">Slot: {{ data.slot }}</p>
      </template>
    </Card>
  </div>
</template>
