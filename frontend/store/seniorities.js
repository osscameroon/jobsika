export const state = () => ({
  seniorities: [],
})
    
export const getters = {
  seniorities: (state) => state.seniorities,
}
    
export const mutations = {
  SETSENIORITIES(state, payload) {
    state.seniorities = [].concat(payload);
  },
}