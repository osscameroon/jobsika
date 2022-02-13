// import axios from "axios";
// import { BASE_URL } from "../constants/api";

export const state = () => ({
    page: 1,
    limit: 20
})

export const getters = () => ({
    page: (state) => state.page,
    limit: (state) => state.limit
})

export const actions = () => ({

})
  
export const mutations = {
    SETPAGE(state, value) {
        state.page = value;
    },
    SETLIMIT(state, value) {
        state.limit = value;
    },
}