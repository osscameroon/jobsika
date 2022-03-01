export const state = () => ({
  cities: [],
})
    
export const getters = {
  cities: (state) => state.cities,
}
    
export const mutations = {
  SETCITIES(state, payload) {
    state.cities = [].concat(payload);
  },
}