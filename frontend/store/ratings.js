
export const state = () => ({
    page: 1,
    limit: 10,
    nbHits: 0,
    companies: []
})

export const getters = {
    page: (state) => state.page,
    limit: (state) => state.limit,
    nbHits: (state) => state.nbHits,
    companies: (state) => state.companies
}
  
export const mutations = {
    SETPAGE(state, value) {
        state.page = value;
    },
    SETLIMIT(state, value) {
        state.limit = value;
    },
    SETNBHITS(state, value) {
        state.nbHits = value;
    },
    SETCOMPANIES(state, payload) {
        state.companies = [].concat(payload);
    },
}