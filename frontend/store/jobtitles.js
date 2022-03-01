export const state = () => ({
  jobtitles: [],
	filterjob: ''
})
    
export const getters = {
  jobtitles: (state) => state.jobtitles,
	filterjob: (state) => state.filterjob,
}
    
export const mutations = {
  SETJOBTITLES(state, payload) {
    state.jobtitles = [].concat(payload);
  },
	SETFILTERJOB(state, value) {
    state.filterjob = value
  },
}