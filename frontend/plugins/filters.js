import Vue from "vue"

Vue.filter('price', (value) => {
  return new Intl.NumberFormat('fr-CM', {style: 'currency', currency: 'XAF'}).format(value);
})
