export const state = () => ({
  cities: [],
  filtercity: '',
})

export const getters = {
  cities: (state) => state.cities,
  filtercity: (state) => state.filtercity,
}

export const mutations = {
  SETCITIES(state, payload) {
    state.cities = [].concat(payload)
  },
  SETFILTERCITY(state, value) {
    state.filtercity = value
  },
}
