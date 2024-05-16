const checkIfUserIsAuthenticated = () => {
  if (localStorage.getItem('token')) {
    return true
  } else {
    return false
  }
}

export default checkIfUserIsAuthenticated
