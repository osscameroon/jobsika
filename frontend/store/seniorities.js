export const state = () => ({
  seniorities: [],
  filterseniority: ''
})
    
export const getters = {
  seniorities: (state) => state.seniorities,
  filterseniority: (state) => state.filterseniority,
}
    
export const mutations = {
  SETSENIORITIES(state, payload) {
    state.seniorities = [].concat(payload);
  },
  SETFILTERSENIORITY(state, value) {
    state.filterseniority = value;
  },
}