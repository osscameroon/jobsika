// @ts-check
import { defineStore, acceptHMRUpdate } from "pinia";
import { 
	IRating,
	GetRatingsResponse,
	GetAverageResponse,
	RootRatingsState,
	GetRatingsParams,
	GetAverageParams
} from "../types/ranting";
import axios from "axios";
const baseURL = "http://localhost:7000";

export const useRatingsStore = defineStore({
	id: "ratings",
	state: () => ({
		ratings: [],
		page: 1,
  		limit: 10,
		offset: 0,
		nbHits: 0,
		average: 0,
  		averageStars: 0,
		loading: false,
		error: null
	} as RootRatingsState),

	getters: {
		items: (state) => state.ratings,
	},

	actions: {
		async getItems(configs: GetRatingsParams) {
			this.loading = true;
      const params = {
				page: configs.page ? configs.page : 1,
				limit: configs.limit ? configs.limit: 10,
				company: configs.company ? configs.company : "",
				jobtitle: configs.jobtitle ? configs.jobtitle : "",
				seniority: configs.seniority ? configs.seniority : "",
				city: configs.city ? configs.city : ""
			};
			try {
				const { data, status } = await axios.get<GetRatingsResponse>(
					`${baseURL}/ratings`,
					{
						params: { ...params },
						headers: {
							"Content-Type": "application/json"
						},
					},
				);
				this.ratings = data.hits;
				this.nbHits = data.nbHits;
				this.offset = data.offset;
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

		async getAverage(configs: GetAverageParams) {
			this.loading = true;
      const params = {
				company: configs.company ? configs.company : "",
				jobtitle: configs.jobtitle ? configs.jobtitle : "",
				seniority: configs.seniority ? configs.seniority : "",
				city: configs.city ? configs.city : ""
			};
			try {
				const { data, status } = await axios.get<GetAverageResponse>(
					`${baseURL}/average-rating`,
					{
						params: { ...params },
						headers: {
							"Content-Type": "application/json"
						},
					},
				);
				this.average = data.salary;
				this.averageStars = data.rating;
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

		async getItem(id: string): Promise<IRating | null> {
			this.loading = true;
			try {
				const { data, status } = await axios.get<IRating>(
					`${baseURL}/ratings/${id}`,
					{
						headers: {
							"Content-Type": "application/json",
						},
					},
				);

				return data;
			} catch (error) {
				if (axios.isAxiosError(error)) {
					this.error = error.message;
				} else {
					this.error = "An unexpected error occurred";
				}
				return null;
			} finally {
				this.loading = false;
			}
		},

		async addItem(rating: IRating) {
			this.loading = true;

			try {
				const { data, status } = await axios.post<IRating>(
					`${baseURL}/ratings`,
					{...rating},
					{
						headers: {
							"Content-Type": "application/json",
						},

					},
				);
				if (status === 200) {
					this.ratings.push(data);
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
	import.meta.hot?.accept(acceptHMRUpdate(useRatingsStore, import.meta.hot))
}