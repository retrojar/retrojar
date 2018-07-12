import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VList,
  VBtn,
  VIcon,
  VGrid,
  VToolbar,
  VCard,
  VSnackbar,
  VBottomNav,
  transitions
} from 'vuetify'
import 'vuetify/src/stylus/app.styl'

Vue.use(Vuetify, {
  components: {
    VApp,
    VNavigationDrawer,
    VList,
    VBtn,
    VIcon,
    VGrid,
    VToolbar,
    VCard,
    VSnackbar,
    VBottomNav,
    transitions
  }
})
