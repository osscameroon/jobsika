import axios from "axios";
import { BASE_URL } from "../constants/api";

export const actions = {
  async getRatings({ commit }, payload) {
    const params = {
      page: payload.page,
      limit: payload.limit,
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : ""
    }
    const resp = await axios.get(
      BASE_URL + '/ratings',
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
      BASE_URL + '/companies')
    if(resp){
      commit("companies/SETCOMPANIES", resp.data);
    }
  },
  async getJobtitles({ commit }) {
    const resp = await axios.get(
      BASE_URL + '/jobtitles')
    if(resp){
      commit("jobtitles/SETJOBTITLES", resp.data);
    }
  },
  async getCities({ commit }) {
    const resp = await axios.get(
      BASE_URL + '/cities')
    if(resp){
      commit("cities/SETCITIES", resp.data);
    }
  },
  async getSeniorities({ commit }) {
    const resp = await axios.get(
      BASE_URL + '/seniority')
    if(resp){
      commit("seniorities/SETSENIORITIES", resp.data);
    }
  },
  async postRating({ commit }, data) {
    data.salary = parseInt(data.salary);
    const resp = await axios.post(BASE_URL + '/ratings', data)
    if (resp) {
      commit("ratings/ADDCOMPANY", resp.data)
    }
  },
  async fetchAverage({ commit }, payload) {
    const params = {
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : ""
    }
    const resp = await axios.get(
      BASE_URL + '/average-rating',
      {
        params: { ...params }
      }
    )
    if (resp) {
      commit("ratings/SETAVERAGE", resp.data.salary)
      // commit("ratings/SELECTVALUESTARS", resp.data.rating)
    }
  },
  filterJob({ commit }, value){
    commit("jobtitles/SETFILTERJOB", value)
  },
  filterCompany({ commit }, value){
    commit("companies/SETFILTERCOMPANY", value)
  },
}
