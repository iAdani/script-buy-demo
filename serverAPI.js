import axios from "axios";

const API_URL = "http://10.0.2.2:8080/";

//Server Auth
async function authUser(token, userUID, API_URL_DEF = API_URL) {
    try {
      axios.defaults.headers.common['Authorization'] = 'Bearer ' + token
      await axios.post(API_URL + 'auth', {
        uid: userUID,
      }).then(response => {
        return response
      }).catch(error => {
        console.log(error)
      })
    } catch (error) {
      console.log(error)
    }
  }


const getAPIUrl = () => {
    return API_URL
}
export { getAPIUrl, authUser }