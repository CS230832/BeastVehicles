import { defineStore } from 'pinia'

export const useStationStore = defineStore('station', {
  state: () => ({
    station: ''
  }),
  actions: {
    updateStation(newStation) {
      this.station = newStation
    }
  }
})
