import { Store } from 'vuex'
import { initializeStores } from '~/utils/storeAccessor'

const initializer = (store: Store<any>) => initializeStores(store)

export const plugins = [initializer]

export * from '@/utils/storeAccessor'
