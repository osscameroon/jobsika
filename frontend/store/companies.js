export const state = () => ({
  companies: [],
  filtercompany: '',
})

export const getters = {
  companies: (state) => state.companies,
  filtercompany: (state) => state.filtercompany,
}

export const mutations = {
  SETCOMPANIES(state, payload) {
    state.companies = [].concat(payload)
  },
  SETFILTERCOMPANY(state, value) {
    state.filtercompany = value
  },
}
