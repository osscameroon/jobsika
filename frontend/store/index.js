import axios from 'axios'

export const actions = {
  async getRatings({ commit }, payload) {
    const params = {
      page: payload.page ? payload.page : 1,
      limit: payload.limit ? payload.limit : 10,
      company: payload.company ? payload.company : '',
      jobtitle: payload.jobtitle ? payload.jobtitle : '',
      seniority: payload.seniority ? payload.seniority : '',
      city: payload.city ? payload.city : '',
    }
    const resp = await axios.get(this.$config.baseURL + '/ratings', {
      params: { ...params },
    })
    if (resp) {
      commit('ratings/SETNBHITS', resp.data.nbHits)
      commit('ratings/SETRATINGS', resp.data.hits)
    }
  },
  async getCompanies({ commit }) {
    const resp = await axios.get(this.$config.baseURL + '/companies')
    if (resp) {
      commit('companies/SETCOMPANIES', resp.data)
    }
  },
  async getJobtitles({ commit }) {
    const resp = await axios.get(this.$config.baseURL + '/jobtitles')
    if (resp) {
      commit('jobtitles/SETJOBTITLES', resp.data)
    }
  },
  async getCities({ commit }) {
    const resp = await axios.get(this.$config.baseURL + '/cities')
    if (resp) {
      commit('cities/SETCITIES', resp.data)
    }
  },
  async getSeniorities({ commit }) {
    const resp = await axios.get(this.$config.baseURL + '/seniority')
    if (resp) {
      commit('seniorities/SETSENIORITIES', resp.data)
    }
  },
  async postRating({ commit }, data) {
    data.salary = parseInt(data.salary)
    const resp = await axios.post(this.$config.baseURL + '/ratings', data)
    if (resp) {
      commit('ratings/ADDRATING', data)
    }
  },
  async fetchAverage({ commit }, payload) {
    const params = {
      company: payload.company ? payload.company : '',
      jobtitle: payload.jobtitle ? payload.jobtitle : '',
      seniority: payload.seniority ? payload.seniority : '',
      city: payload.city ? payload.city : '',
    }
    const resp = await axios.get(this.$config.baseURL + '/average-rating', {
      params: { ...params },
    })
    if (resp) {
      commit('ratings/SETAVERAGE', resp.data.salary)
      commit('ratings/SETAVERAGESTARS', resp.data.rating)
    }
  },
  async getJobs({ commit }, payload) {
    const params = {
      page: payload.page ? payload.page : 1,
      limit: payload.limit ? payload.limit : 10,
    }
    const resp = await axios.get(this.$config.baseURL + '/jobs', {
      params: { ...params },
    })
    if (resp) {
      commit('jobs/SETNBHITS', resp.data.nbHits)
      commit('jobs/SETJOBS', resp.data.hits)
    }
  },
  async postJob({ commit }, data) {
    try {
      const resp = await axios.post(this.$config.baseURL + '/jobs', data)
      if (resp) {
        return { status: true, data: resp.data }
      }
    } catch (error) {
      return { status: false, data: {} }
    }
  },
  async postSubscriber(_, data) {
    try {
      const resp = await axios.post(this.$config.baseURL + '/subscribers', data)
      return resp
    } catch (error) {
      return undefined
    }
  },
  async getJobPaymentLink({ commit }, data) {
    try {
      const resp = await axios.post(this.$config.baseURL + '/pay', data)
      commit('jobs/SETPAYMENTLINK', resp.data)
      return resp
    } catch (error) {
      return undefined
    }
  },
  filterJob({ commit }, value) {
    commit('jobtitles/SETFILTERJOB', value)
  },
  filterCompany({ commit }, value) {
    commit('companies/SETFILTERCOMPANY', value)
  },
  filterSeniority({ commit }, value) {
    commit('seniorities/SETFILTERSENIORITY', value)
  },
  filterCity({ commit }, value) {
    commit('cities/SETFILTERCITY', value)
  },
  setNewJob({ commit }, value) {
    commit('jobs/SETNEWJOB', value)
  },
}
