import Vue from "vue"

Vue.filter('price', (value) => {
  return value.toString().replace(/(?<=\d)(?=(\d{3})+(?!\d))/g, ',') + " FCFA"
})
