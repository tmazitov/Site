import axios from 'axios'
import { readValue } from '../actions/jwt';

const client = axios.create({
    baseURL: process.env.VUE_APP_API_PATH,
    headers: {
      "Authorization": "Bearer "+ readValue(),
    },
    withCredentials: true 
  });

const auth = axios.create({
  baseURL: process.env.VUE_APP_API_PATH,
  withCredentials: true
})

export {auth}
export default client