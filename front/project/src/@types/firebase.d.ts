import firebase from 'firebase/app'
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
