import axios from "axios";
import { BASE_URL } from "../constants/api";

export const actions = {
    async getCompanies({ commit }, payload){
        const params = {
            page: payload.page,
            limit: payload.limit
        }
        const resp = await axios.get(
            BASE_URL + '/ratings',
            {
              params: {...params}
            }
        );
        console.log("resp", resp);
        if(resp){
            commit("ratings/SETNBHITS", resp.data.nbHits);
            commit("ratings/SETCOMPANIES", resp.data.hits);
        }
    }
}