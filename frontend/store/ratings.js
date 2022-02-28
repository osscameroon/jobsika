
export const state = () => ({
  page: 1,
  limit: 10,
  nbHits: 0,
  ratings: [],
  average: 0,
})

export const getters = {
  page: (state) => state.page,
  limit: (state) => state.limit,
  nbHits: (state) => state.nbHits,
  ratings: (state) => state.ratings,
  average: (state) => state.average,
}

export const mutations = {
  SETPAGE(state, value) {
    state.page = value;
  },
  SETLIMIT(state, value) {
    state.limit = value;
  },
  SETNBHITS(state, value) {
    state.nbHits = value;
  },
  SETCOMPANIES(state, payload) {
    state.companies = [].concat(payload);
  },
  ADDCOMPANY(state, payload) {
    state.companies = [{ company_name: payload.company_name, salary: payload.salary, city: payload.city, seniority: payload.seniority, rating: payload.rating, comment: payload.comment, job_title: payload.job_title }, ...state.companies]
  },
  SELECTVALUECOMPANY(state, value) {
    state.selectvaluecompany = value
  },
  SELECTVALUEJOB(state, value) {
    state.selectvaluejob = value
  },
  SELECTVALUECITY(state, value) {
    state.selectvaluecity = value
  },
  SELECTVALUESALARY(state, value) {
    state.selectvaluesalary = value
  },
  SELECTVALUECOMMENT(state, value) {
    state.selectvaluecomment = value
  },
  SETFILTERJOB(state, value) {
    state.filterjob = value
  },
  SETFILTERCOMPANY(state, value) {
    state.filtercompany = value
  },
  SETAVERAGE(state, value) {
    state.average = value
  },
  SELECTVALUESTARS(state, value) {
    state.selectvaluestars = value
  },
}
