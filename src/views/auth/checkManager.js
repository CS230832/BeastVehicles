import ApiService from '@/api'
import checkIfUserIsAuthenticated from './checkAuth'

let userData = null

const checkIfUserIsManager = async () => {
  if (checkIfUserIsAuthenticated()) {
    try {
      userData = await ApiService.getUser(
        localStorage.getItem('username'),
        localStorage.getItem('token')
      )
      return userData.data.role === 'manager'
    } catch (error) {
      console.log('Error getting user data: ', error.response.data.data)
      return false
    }
  } else {
    return
  }
}

export default checkIfUserIsManager
