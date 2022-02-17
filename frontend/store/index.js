import axios from "axios";
import { BASE_URL } from "../constants/api";

export const actions = {
  async getCompanies({ commit }, payload) {
    const params = {
      page: payload.page,
      limit: payload.limit
    }
    const resp = await axios.get(
      BASE_URL + '/ratings',
      {
        params: { ...params }
      }
    );
    console.log("resp", resp);
    if (resp) {
      commit("ratings/SETNBHITS", resp.data.nbHits);
      commit("ratings/SETCOMPANIES", resp.data.hits);
    }
  },
  async postCompany({ commit }, payload) {
    const data = {
      companyname: payload.company_name,
      salary: payload.salary,
      city: payload.city,
      seniority: payload.seniority,
      rating: payload.rating,
      comment: payload.comment
    }
    const resp = await axios.post(BASE_URL + '/ratings', data)
    console.log("resp", resp.data);
    if (resp) {
      commit("ratings/ADDCOMPANY", resp.data)
    }
  }
}
