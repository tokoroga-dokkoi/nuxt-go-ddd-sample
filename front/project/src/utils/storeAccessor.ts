/* eslint-disable import/no-mutable-exports */
import { Store } from 'vuex'
import { getModule } from 'vuex-module-decorators'
import SignUpStore from '~/store/modules/user/SignUpStore'

let UserSignUpStore: SignUpStore

function initializeStores(store: Store<any>): void {
  UserSignUpStore = getModule(SignUpStore, store)
}

export { initializeStores, UserSignUpStore }
