<script setup>
import { ref, computed } from 'vue'

import InputText from 'primevue/inputtext'
import Button from 'primevue/button'

import { useStationStore } from '@/stores'

const store = useStationStore()
const stationName = ref(null)
const currentStation = computed(() => store.station)

const sendStationName = () => {
  store.updateStation(stationName.value)
}
</script>

<template>
  <div class="flex flex-wrap gap-2 items-center">
    <InputText type="text" placeholder="Enter station name" v-model="stationName" />
    <Button label="Submit" @click="sendStationName" />
    <p>Current Station: "{{ currentStation }}"</p>
  </div>
</template>
