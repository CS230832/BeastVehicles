import axios from 'axios'

const API_URL = 'http://192.168.1.16:8080/api/v1'

const ApiService = {
  createStation: async (name, region, max) => {
    try {
      const response = await axios.post(`${API_URL}/parking`, { name, region, max })
      return response.data
    } catch (error) {
      console.error('Error creating station:', error)
      throw error
    }
  },

  addVehicle: async (wincode, parking) => {
    try {
      const response = await axios.post(`${API_URL}/vehicle`, { wincode, parking })
      return response.data
    } catch (error) {
      console.error('Error adding vehicle:', error)
      throw error
    }
  },

  findVehicle: async (wincode) => {
    try {
      const response = await axios.get(`${API_URL}/vehicle/${wincode}`)
      return response.data
    } catch (error) {
      console.error('Error finding vehicle:', error)
      throw error
    }
  },

  removeVehicle: async (wincode) => {
    try {
      const response = await axios.delete(`${API_URL}/vehicle/${wincode}`)
      return response.data
    } catch (error) {
      console.error('Error removing vehicle:', error)
      throw error
    }
  },

  getFreeBlocks: async (parking) => {
    try {
      const response = await axios.get(`${API_URL}/parking/${parking}/free`)
      return response.data
    } catch (error) {
      console.error('Error getting free slots:', error)
      throw error
    }
  },

  getFullBlocks: async (parking) => {
    try {
      const response = await axios.get(`${API_URL}/parking/${parking}/full`)
      return response.data
    } catch (error) {
      console.error('Error getting full slots:', error)
      throw error
    }
  },

  getFreeSlots: async (parking, name) => {
    try {
      const response = await axios.get(`${API_URL}/block/free?parking=${parking}&name=${name}`)
      return response.data
    } catch (error) {
      console.error('Error fetching slots:', error)
      throw error
    }
  },

  getFullSlots: async (parking, name) => {
    try {
      const response = await axios.get(`${API_URL}/block/full?parking=${parking}&name=${name}`)
      return response.data
    } catch (error) {
      console.error('Error fetching slots:', error)
      throw error
    }
  }
}

export default ApiService
