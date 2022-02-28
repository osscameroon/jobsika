import Vue from "vue"

Vue.filter('price', (value = 1) => {
  return value.toString().replace(/(?<=\d)(?=(\d{3})+(?!\d))/g, ',') + " FCFA"
})
