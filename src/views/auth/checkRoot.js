import ApiService from '@/api'

var userData = {}

const checkIfUserIsRoot = async () => {
  try {
    userData = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )
    return userData.data.role === 'root'
  } catch (error) {
    console.log('Error getting user data: ', error)
    return false
  }
}

export default checkIfUserIsRoot
