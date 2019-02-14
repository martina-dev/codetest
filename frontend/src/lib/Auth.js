import { setCookie, getCookie } from './Session'
import decode from 'jwt-decode'

const API_HOST = "http://localhost:8000"

export const authenticate = (formData) => {

    fetch(`${API_HOST}/login`, {
        method: "POST",
        body: formData,
        headers: { Authorization: "Bearer atoken"},
    }).then(res => res.json())
    .then(response =>  setCookie(response))
    .catch(error => console.error('Error:', error))

}
  
export const isAuthenticated = () => !!getCookie("jwt");

export const register = (formData) => {
    fetch(`${API_HOST}/create`, {
        method: "POST",
        body: formData,
        headers: { Authorization: "Bearer atoken"},
    }).then(res => res.json())
      .then(response => setCookie(response))
} 


export const getRole = () => {
    const cookie = getCookie()
    const decoded = decode(cookie)
    return decoded.Role
}