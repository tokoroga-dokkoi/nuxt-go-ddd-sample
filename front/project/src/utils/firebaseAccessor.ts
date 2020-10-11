import firebase from 'firebase/app'
import 'firebase/auth'
import { FirebaseOptions } from '@firebase/app-types'

/* eslint-disable import/no-mutable-exports */
let $firebase: any

export function initializeFirebase() {
  $firebase = firebase
  const firebaseConfig: FirebaseOptions = {
    apiKey: process.env.FIREBASE_API_KEY,
    authDomain: process.env.FIREBASE_AUTH_DOMAIN,
    projectId: process.env.FIREBASE_PROJECT_ID,
  }
  if (!$firebase.apps.length) {
    $firebase.initializeApp(firebaseConfig)
  }
  return $firebase
}

export { $firebase }
