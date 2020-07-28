import Vue from 'vue'
import App from './App.vue'
import VCalendar from 'v-calendar';


Vue.config.productionTip = false


new Vue({
  render: h => h(App),
}).$mount('#app')

Vue.use(VCalendar, {
  componentPrefix: 'vc',  // Use <vc-calendar /> instead of <v-calendar />
});