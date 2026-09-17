import axios from "axios";
export default defineNuxtPlugin((nuxtApp) => {
  const api = axios.create({
    baseURL: "http://localhost:7000",
    headers: {
      common: {},
    },
  });
	return {
    provide: {
      axios: api,
    },
  };
});