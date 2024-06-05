import { defineStore } from 'pinia'

export const useStationStore = defineStore('station', {
  state: () => ({
    station: null
  }),
  actions: {
    updateStation(newStation) {
      this.station = newStation
    }
  }
})
