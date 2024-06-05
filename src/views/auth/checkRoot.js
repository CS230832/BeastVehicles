import ApiService from '@/api'

let userData = null

const checkIfUserIsRoot = async () => {
  try {
    userData = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )
    return userData.data.role === 'root'
  } catch (error) {
    console.log('Error getting user data: ', error.response.data.data)
    return false
  }
}

export default checkIfUserIsRoot
