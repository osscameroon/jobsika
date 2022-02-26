import axios from "axios";
import { API_BASE_URL } from "../utils/config";

export const actions = {
  async getCompanies({ commit }, payload) {
    const params = {
      page: payload.page,
      limit: payload.limit,
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : ""
    }
    const resp = await axios.get(
      API_BASE_URL + '/ratings',
      {
        params: { ...params }
      }
    )
    if (resp) {
      commit("ratings/SETNBHITS", resp.data.nbHits);
      commit("ratings/SETCOMPANIES", resp.data.hits);
    }
  },
  async postCompany({ commit }, payload) {
    const data = {
      company_name: payload.company_name,
      salary: payload.salary,
      city: payload.city,
      seniority: payload.seniority,
      rating: payload.rating,
      comment: payload.comment,
      job_title: payload.job_title
    }
    const resp = await axios.post(API_BASE_URL + '/ratings', data)
    if (resp) {
      commit("ratings/ADDCOMPANY", resp.data)
    }
  },
  selectValueCompany({ commit }, payload) {
    commit('ratings/SELECTVALUECOMPANY', payload)
  },
  selectValueJob({ commit }, payload) {
    commit('ratings/SELECTVALUEJOB', payload)
  },
  selectValueCity({ commit }, payload) {
    commit('ratings/SELECTVALUECITY', payload)
  },
  selectValueSalary({ commit }, payload) {
    commit('ratings/SELECTVALUESALARY', payload)
  },
  selectValueComment({ commit }, payload) {
    commit('ratings/SELECTVALUECOMMENT', payload)
  },
  filterJob({ commit }, payload) {
    commit('ratings/SETFILTERJOB', payload)
  },
  filterCompany({ commit }, payload) {
    commit('ratings/SETFILTERCOMPANY', payload)
  },
  selectValueStars({ commit }, payload) {
    commit('ratings/SELECTVALUESTARS', payload)
  },
  async fetchAverage({ commit }, payload) {
    const params = {
      company: payload.company ? payload.company : "",
      jobtitle: payload.jobtitle ? payload.jobtitle : ""
    }
    const resp = await axios.get(
      API_BASE_URL + '/average-rating',
      {
        params: { ...params }
      }
    )
    if (resp) {
      commit("ratings/SETAVERAGE", resp.data.salary)
      commit("ratings/SELECTVALUESTARS", resp.data.rating)
    }
  },
}
