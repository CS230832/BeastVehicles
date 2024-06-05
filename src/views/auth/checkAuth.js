const checkIfUserIsAuthenticated = () => {
  if (localStorage.getItem('token') && localStorage.getItem('username')) {
    return true
  } else {
    return false
  }
}

export default checkIfUserIsAuthenticated
