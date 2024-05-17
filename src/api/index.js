import axios from 'axios'

const API_URL = 'http://192.168.1.16:8080/api/v1'

const ApiService = {
  createStation: async (name, region, capacity, token) => {
    try {
      const response = await axios.post(
        `${API_URL}/parkings/register`,
        { name, region, capacity: parseInt(capacity) },
        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error creating station: ', error)
      throw error
    }
  },

  removeStation: async (name, token) => {
    try {
      const response = await axios.delete(
        `${API_URL}/parkings/delete?name=${name}`,

        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error removing station: ', error)
      throw error
    }
  },

  getStation: async (name) => {
    try {
      const response = await axios.get(`${API_URL}/parkings/info?name=${name}`)
      return response.data
    } catch (error) {
      console.log('Error getting station: ', error)
      throw error
    }
  },

  addVehicle: async (wincode, token) => {
    try {
      const response = await axios.post(
        `${API_URL}/vehicles/register`,
        { wincode },

        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error adding vehicle: ', error)
      throw error
    }
  },

  addMultipleVehicles: async (wincodes, token) => {
    try {
      const response = await axios.post(
        `${API_URL}/vehicles/set/register`,
        wincodes,

        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error adding vehicles: ', error)
      throw error
    }
  },

  findVehicle: async (wincode) => {
    try {
      const response = await axios.get(`${API_URL}/vehicles/info?wincode=${wincode}`)
      return response.data
    } catch (error) {
      console.error('Error finding vehicle: ', error)
      throw error
    }
  },

  findMultipleVehicles: async (wincodes) => {
    try {
      const response = await axios.post(`${API_URL}/vehicles/set/info`, wincodes)
      return response.data
    } catch (error) {
      console.error('Error finding vehicle:', error)
      throw error
    }
  },

  removeVehicle: async (wincode, token) => {
    try {
      const response = await axios.delete(
        `${API_URL}/vehicles/delete?wincode=${wincode}`,
        {},

        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error removing vehicle: ', error)
      throw error
    }
  },

  removeMultipleVehicles: async (wincodes, token) => {
    try {
      const response = await axios.delete(
        `${API_URL}/vehicles/set/delete`,
        wincodes,

        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.error('Error removing vehicle: ', error)
      throw error
    }
  },

  getAllBlocks: async (parking) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/all/info?parking=${parking}`)
      return response.data
    } catch (error) {
      console.error('Error getting all slots: ', error)
      throw error
    }
  },

  getFreeBlocks: async (parking) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/all/free?parking=${parking}`)
      return response.data
    } catch (error) {
      console.error('Error getting free slots: ', error)
      throw error
    }
  },

  getFullBlocks: async (parking) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/all/full?parking=${parking}`)
      return response.data
    } catch (error) {
      console.error('Error getting full slots: ', error)
      throw error
    }
  },

  getAllSlots: async (parking, name) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/info?parking=${parking}&name=${name}`)
      return response.data
    } catch (error) {
      console.error('Error getting all slots in specific station: ', error)
      throw error
    }
  },

  getFreeSlots: async (parking, name) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/free?parking=${parking}&name=${name}`)
      return response.data
    } catch (error) {
      console.error('Error getting free slots in specific station: ', error)
      throw error
    }
  },

  getFullSlots: async (parking, name) => {
    try {
      const response = await axios.get(`${API_URL}/blocks/full?parking=${parking}&name=${name}`)
      return response.data
    } catch (error) {
      console.error('Error getting full slots in specific station: ', error)
      throw error
    }
  },

  login: async (username, password) => {
    try {
      const response = await axios.post(`${API_URL}/users/login`, { username, password })
      return response.data
    } catch (error) {
      console.error('Error logging in:', error)
      throw error
    }
  },

  logout: async (token) => {
    try {
      await axios.post(
        `${API_URL}/users/logout`,
        {},
        {
          headers: {
            Authorization: token
          }
        }
      )
    } catch (error) {
      console.error('Error loggin out:', error)
      throw error
    }
  },

  register: async (username, password, role, firstName, lastName, parking, token) => {
    try {
      const response = await axios.post(
        `${API_URL}/users/register`,
        { username, password, role, firstName, lastName, parking },
        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.log('Error registering user: ', error)
      throw error
    }
  },

  getUser: async (username, token) => {
    try {
      const response = await axios.get(`${API_URL}/users/info?username=${username}`, {
        headers: {
          Authorization: token
        }
      })
      return response.data
    } catch (error) {
      console.log('Error getting user: ', error)
      throw error
    }
  },

  removeUser: async (username, token) => {
    try {
      const response = await axios.delete(
        `${API_URL}/users/delete?username=${username}`,
        {},
        {
          headers: {
            Authorization: token
          }
        }
      )
      return response.data
    } catch (error) {
      console.log('Error removing user: ', error)
      throw error
    }
  }
}

export default ApiService
