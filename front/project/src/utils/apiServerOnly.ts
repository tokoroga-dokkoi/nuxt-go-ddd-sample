import { NuxtAxiosInstance } from '@nuxtjs/axios'

/* eslint no-mutable-exports: off */
let $axios: NuxtAxiosInstance

export function initializeAxios(axiosInstance: NuxtAxiosInstance): void {
  $axios = axiosInstance
  if (process.env.production) {
    $axios.defaults.baseURL = 'http://localhost:8000'
  } else {
    $axios.defaults.baseURL = 'http://proxy:8000/api'
  }
}

export { $axios }
