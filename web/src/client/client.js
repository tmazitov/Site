import axios from 'axios'
import { refreshTokens } from '../actions/auth';
import { tokenIsValid, readValue } from '../actions/jwt';

const client = axios.create({
    baseURL: process.env.VUE_APP_API_PATH,
    withCredentials: true 
  });

client.interceptors.request.use(
  async function(request){
    if ( !tokenIsValid() ){
      await refreshTokens()
    }
    request.headers.Authorization = 'Bearer ' + readValue()
    return request;
  },
  (error) => {
    if (error.statusCode == 401){
      refreshTokens()
    }
  }
)

const auth = axios.create({
  baseURL: process.env.VUE_APP_API_PATH,
  withCredentials: true
})

export {auth}
export default client