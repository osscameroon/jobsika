// @ts-check
import { defineStore, acceptHMRUpdate } from "pinia";
import { RootCommonsState } from "../types/common";
import axios from "axios";
const baseURL = "http://localhost:7000";

export const useSenioritiesStore = defineStore({
	id: "seniorities",
	state: () => ({
		data: [],
		filter: "",
		loading: false,
		error: null
	} as RootCommonsState),

	getters: {
		items: (state) => state.data,
	},

	actions: {
		async getItems() {
			this.loading = true;
			try {
				const { data, status } = await axios.get<string[]>(
					`${baseURL}/seniority`,
					{
						headers: {
							"Content-Type": "application/json"
						},
					},
				);
				this.data = data;
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

		async addItem(jobtitle: string) {
			this.loading = true;

			try {
				const { data, status } = await axios.post<string>(
					`${baseURL}/seniority`,
					{jobtitle: jobtitle},
					{
						headers: {
							"Content-Type": "multipart/form-data",
						},

					},
				);
				if (status === 200) {
					this.data.push(data);
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
	import.meta.hot?.accept(acceptHMRUpdate(useSenioritiesStore, import.meta.hot))
}