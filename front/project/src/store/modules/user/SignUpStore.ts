import { $firebase } from '@/utils/firebaseAccessor'
import { Action, Module, VuexModule } from 'vuex-module-decorators'
import { $axios } from '~/utils/apiClientOnly'

type SignUpType = 'email' | 'google' | 'facebook'
type AuthParams = {
  email?: string
  password?: string
  // signUpType: SignUpType
}
@Module({
  name: 'modules/user/SignUpStore',
  stateFactory: true,
  namespaced: true,
})
export default class SignUpStore extends VuexModule {
  // state
  private signUpResult: boolean = false
  private currentUser: User = {}
}
