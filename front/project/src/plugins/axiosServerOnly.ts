import { Plugin } from '@nuxt/types'
import { initializeAxios } from '@/utils/apiServerOnly'

export const accessor: Plugin = ({ $axios }): void => {
  initializeAxios($axios)
}

export default accessor
