import firebase from 'firebase/app'
import 'firebase/auth'
import { FirebaseOptions } from '@firebase/app-types'
import { Plugin } from '@nuxt/types'

const firebaseConfig: FirebaseOptions = {
  apiKey: process.env.FIREBASE_API_KEY,
  authDomain: process.env.FIREBASE_AUTH_DOMAIN,
  projectId: process.env.FIREBASE_PROJECT_ID,
}

if (!firebase.apps.length) {
  firebase.initializeApp(firebaseConfig)
}

// declare types
declare module 'vue/types/vue' {
  interface Vue {
    $firebase: any
    $auth: firebase.auth.Auth
  }
}

declare module '@nuxt/types' {
  interface NuxtAppOptions {
    $firebase: any
    $auth: firebase.auth.Auth
  }
}

declare module 'vuex/types/index' {
  interface Store<S> {
    $firebase: any
    $auth: firebase.auth.Auth
  }
}

const firebasePlugin: Plugin = (context, inject) => {
  inject('firebase', firebase)
  inject('auth', firebase.auth())
}

export default firebasePlugin
