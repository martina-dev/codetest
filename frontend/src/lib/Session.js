import Cookies from 'universal-cookie';
import cookie from "js-cookie";

export const setCookie = (data) => {
    const cookies = new Cookies()
    cookies.set('jwt', data, {
        path: "/"
    })
}

export const getCookie = () => {
    return cookie.get('jwt')
}


