import { Plugin } from '@nuxt/types'
import { initializeFirebase } from '@/utils/firebaseAccessor'

const firebasePlugin: Plugin = (context, inject) => {
  const firebase = initializeFirebase()
  inject('firebase', firebase)
  inject('auth', firebase.auth())
}

export default firebasePlugin
