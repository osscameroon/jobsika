// @ts-check
import { defineStore, acceptHMRUpdate } from "pinia";
import { RootCompaniesState, ICompany } from "../types/company";
import axios from "axios";
const baseURL = "http://localhost:7000";

export const useCompaniesStore = defineStore({
	id: "companies",
	state: () => ({
		companies: [],
		filtercompany: "",
		loading: false,
		error: null
	} as RootCompaniesState),

	getters: {
		items: (state) => state.companies,
	},

	actions: {
		async getItems() {
			this.loading = true;
			try {
				const { data, status } = await axios.get<ICompany[]>(
					`${baseURL}/companies`,
					{
						headers: {
							"Content-Type": "application/json"
						},
					},
				);
				this.companies = data;
			} catch (error) {
				if (axios.isAxiosError(error)) {
					this.error = error.message;
				} else {
					this.error = "An unexpected error occurred";
				}
			} finally {
				this.loading = false;
			}
		},

		async addItem(company: ICompany) {
			this.loading = true;

			try {
				const { data, status } = await axios.post<ICompany>(
					`${baseURL}/companies`,
					{...company},
					{
						headers: {
							"Content-Type": "multipart/form-data",
						},

					},
				);
				if (status === 200) {
					this.companies.push(data);
				}
			} catch (error) {
				if (axios.isAxiosError(error)) {
					this.error = error.message;
				} else {
					this.error = "An unexpected error occurred";
				}
			} finally {
				this.loading = false;
			}
		}
	},
})

if (import.meta.hot) {
	import.meta.hot?.accept(acceptHMRUpdate(useCompaniesStore, import.meta.hot))
}