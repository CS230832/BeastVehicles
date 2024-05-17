import ApiService from '@/api'

var userData = {}

const checkIfUserIsRoot = async () => {
  try {
    userData = await ApiService.getUser(
      localStorage.getItem('username'),
      localStorage.getItem('token')
    )
    // console.log(userData.data.role == 'root')
    return userData.data.role === 'root'
  } catch (error) {
    console.log('Error getting user data: ', error)
    return false
  }
}

export default checkIfUserIsRoot
