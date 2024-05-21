<script setup>
import { ref } from 'vue'

import InputText from 'primevue/inputtext'

import { useStationStore } from '@/stores'

const store = useStationStore()
const stationName = ref(null)

const sendStationName = () => {
  store.updateStation(stationName.value)
}
</script>

<template>
  <InputText
    type="text"
    placeholder="Enter station name"
    v-model="stationName"
    @input="sendStationName"
  />
</template>
