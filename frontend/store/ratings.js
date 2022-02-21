
export const state = () => ({
  page: 1,
  limit: 10,
  nbHits: 0,
  companies: [],
  selectvaluecompany: '',
  selectvaluejob: '',
  selectvaluecity: '',
  selectvaluesalary: 0,
  selectvaluecomment: '',
  selectvaluestars: 0
})

export const getters = {
  page: (state) => state.page,
  limit: (state) => state.limit,
  nbHits: (state) => state.nbHits,
  companies: (state) => state.companies,
  selectvaluecompany: (state) => state.selectvaluecompany,
  selectvaluejob: (state) => state.selectvaluejob,
  selectvaluecity: (state) => state.selectvaluecity,
  selectvaluesalary: (state) => state.selectvaluesalary,
  selectvaluecomment: (state) => state.selectvaluecomment,
  selectvaluestars: (state) => state.selectvaluestars
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
  SELECTVALUESTARS(state, value) {
    state.selectvaluestars = value
  },
}
