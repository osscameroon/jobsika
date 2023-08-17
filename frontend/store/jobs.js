export const state = () => ({
  page: 1,
  limit: 10,
  nbHits: 0,
  jobs: [],
  newjob: {
    company_name: '',
    company_email: '',
    job_title: '',
    is_remote: false,
    city: '',
    country: 'Cameroon',
    department: '',
    salary_range_min: 0,
    salary_range_max: 0,
    description: '',
    benefits: '',
    how_to_apply: '',
    application_url: '',
    application_email_address: '',
    application_phone_number: '',
    tags: '',
  },
  paymentlink: '',
})

export const getters = {
  page: (state) => state.page,
  limit: (state) => state.limit,
  nbHits: (state) => state.nbHits,
  jobs: (state) => state.jobs,
  newjob: (state) => state.newjob,
  paymentlink: (state) => state.paymentlink,
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
    state.newjob = { ...payload }
  },
  ADDJOB(state, payload) {
    state.jobs = [{ ...payload }, ...state.jobs]
  },
  SETPAYMENTLINK(state, payload) {
    state.paymentlink = payload.tier_url
  },
}
