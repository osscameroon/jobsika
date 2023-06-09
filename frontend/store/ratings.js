export const state = () => ({
  page: 1,
  limit: 10,
  nbHits: 0,
  ratings: [],
  average: 0,
  averageStars: 0,
})

export const getters = {
  page: (state) => state.page,
  limit: (state) => state.limit,
  nbHits: (state) => state.nbHits,
  ratings: (state) => state.ratings,
  average: (state) => state.average,
  averageStars: (state) => state.averageStars,
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
  SETRATINGS(state, payload) {
    state.ratings = [].concat(payload)
  },
  ADDRATING(state, payload) {
    state.ratings = [
      {
        company_name: payload.company_name,
        salary: payload.salary,
        city: payload.city,
        seniority: payload.seniority,
        rating: payload.rating,
        comment: payload.comment,
        job_title: payload.job_title,
      },
      ...state.ratings,
    ]
  },
  SETAVERAGE(state, value) {
    state.average = value
  },
  SETAVERAGESTARS(state, value) {
    state.averageStars = value
  },
}
