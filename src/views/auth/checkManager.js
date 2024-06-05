import ApiService from '@/api'

let userData = null

const checkIfUserIsManager = async () => {
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
}

export default checkIfUserIsManager
