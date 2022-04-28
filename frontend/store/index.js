import axios from "axios";

export const actions = {
  async getRatings({ commit }, payload) {
    const params = {
      page: payload.page,
      limit: payload.limit,
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : "",
      seniority: payload.seniority ? payload.seniority : "",
    }
    const resp = await axios.get(
      this.$config.baseURL + '/ratings',
      {
        params: { ...params }
      }
    )
    if (resp) {
      commit("ratings/SETNBHITS", resp.data.nbHits);
      commit("ratings/SETRATINGS", resp.data.hits);
    }
  },
  async getCompanies({ commit }) {
    const resp = await axios.get(
      this.$config.baseURL+ '/companies')
    if(resp){
      commit("companies/SETCOMPANIES", resp.data);
    }
  },
  async getJobtitles({ commit }) {
    const resp = await axios.get(
      this.$config.baseURL + '/jobtitles')
    if(resp){
      commit("jobtitles/SETJOBTITLES", resp.data);
    }
  },
  async getCities({ commit }) {
    const resp = await axios.get(
      this.$config.baseURL + '/cities')
    if(resp){
      commit("cities/SETCITIES", resp.data);
    }
  },
  async getSeniorities({ commit }) {
    const resp = await axios.get(
      this.$config.baseURL + '/seniority')
    if(resp){
      commit("seniorities/SETSENIORITIES", resp.data);
    }
  },
  async postRating({ commit }, data) {
    data.salary = parseInt(data.salary);
    const resp = await axios.post(this.$config.baseURL + '/ratings', data)
    if(resp) {
      commit("ratings/ADDRATING", data)
    }
  },
  async fetchAverage({ commit }, payload) {
    const params = {
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : "",
      seniority: payload.seniority ? payload.seniority : "",
    }
    const resp = await axios.get(
      this.$config.baseURL + '/average-rating',
      {
        params: { ...params }
      }
    )
    if (resp) {
      commit("ratings/SETAVERAGE", resp.data.salary)
      commit("ratings/SETAVERAGESTARS", resp.data.rating)
    }
  },
  filterJob({ commit }, value){
    commit("jobtitles/SETFILTERJOB", value)
  },
  filterCompany({ commit }, value){
    commit("companies/SETFILTERCOMPANY", value)
  },
  filterSeniority({ commit }, value){
    commit("seniorities/SETFILTERSENIORITY", value)
  },
}
