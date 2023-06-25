export const state = () => ({
  page: 1,
  limit: 10,
  nbHits: 0,
  jobs: [],
  newjob: {}
})

export const getters = {
  page: (state) => state.page,
  limit: (state) => state.limit,
  nbHits: (state) => state.nbHits,
  jobs: (state) => state.jobs,
  newjob: (state) => state.newjob
}

export const mutations = {
  SETPAGE(state, value) {
    state.page = value
  },
  SETLIMIT(state, value) {
    state.limit = value
  },
  SETNBHITS(state, value) {
    state.nbHits = value
  },
  SETJOBS(state, payload) {
    state.jobs = [].concat(payload)
  },
  SETNEWJOB(state, payload) {
    state.newjob = {...payload};
  },
  ADDJOB(state, payload) {
    state.jobs = [{ ...payload }, ...state.jobs]
  },
}
