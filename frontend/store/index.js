import axios from "axios";
import { BASE_URL } from "../constants/api";

export const actions = {
  async getCompanies({ commit }, payload) {
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
    console.log("resp", resp.data.hits.length);
    if (resp) {
      commit("ratings/SETNBHITS", resp.data.nbHits);
      commit("ratings/SETCOMPANIES", resp.data.hits);
      commit("ratings/SETAVERAGE", resp.data)
      let average = 0;
      for (let item = 0; item < resp.data.hits.length; item++) {
        average += resp.data.hits[item].salary
        commit("ratings/SETAVERAGE", average / resp.data.hits.length);
      }
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
    const resp = await axios.post(BASE_URL + '/ratings', data)
    console.log("resp", resp.data);
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
  async fetchAverage({ commit }, payload) {
    const resp = await axios.get(
      BASE_URL + '/average-rating',
    )
    if (resp) {
      commit("ratings/SETAVERAGE", resp.data)
    }
  }
}
