// @ts-check
import { defineStore, acceptHMRUpdate } from "pinia";
import { IJob, GetJobsResponse, RootJobsState, GetJobsParams } from "../types/job";
import axios from "axios";
const baseURL = "http://localhost:7000";

export const useJobsStore = defineStore({
	id: "jobs",
	state: () => ({
		jobs: [],
		page: 1,
  	limit: 10,
		offset: 0,
		nbHits: 0,
		loading: false,
		error: null
	} as RootJobsState),

	getters: {
		items: (state) => state.jobs,
	},

	actions: {
		async getItems(configs: GetJobsParams) {
			this.loading = true;
      const params = {...configs};
			try {
				const { data, status } = await axios.get<GetJobsResponse>(
					`${baseURL}/jobs`,
					{
						params: { ...params },
						headers: {
							"Content-Type": "application/json"
						},
					},
				);
				this.jobs = data.hits;
				this.offset = data.offset;
				this.nbHits = data.nbHits;
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

		async getItem(id: string): Promise<IJob | null> {
			this.loading = true;
			try {
				const { data, status } = await axios.get<IJob>(
					`${baseURL}/jobs/${id}`,
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

		async addItem(job: IJob) {
			this.loading = true;
			const form = new FormData();
			form.append('company_name', job.company_name);
			form.append('job_title', job.job_title);
			form.append('is_remote', job.is_remote ? "1" : "0");
			form.append('city', job.city);
			form.append('country', job.country);
			form.append('department', job.department || "");
			form.append('salary_range_min', String(job.salary_range_min));
			form.append('salary_range_max', String(job.salary_range_max));
			form.append('description', job.description  || "");
			form.append('benefits', job.benefits || "");
			form.append('how_to_apply', job.how_to_apply || "");
			form.append('application_url', job.application_url || "");
			form.append('application_phone_number', job.application_phone_number || "");
			form.append('tags', job.tags || "");

			try {
				const { data, status } = await axios.post<IJob>(
					`${baseURL}/jobs`,
					form,
					{
						headers: {
							"Content-Type": "multipart/form-data",
						},

					},
				);
				if (status === 200) {
					this.jobs.push(data);
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
		},

		async updateItem(job: IJob) {
			this.loading = true;
			const form = new FormData();
			form.append('company_name', job.company_name);
			form.append('job_title', job.job_title);
			form.append('is_remote', job.is_remote ? "1" : "0");
			form.append('city', job.city);
			form.append('country', job.country);
			form.append('department', job.department || "");
			form.append('salary_range_min', String(job.salary_range_min));
			form.append('salary_range_max', String(job.salary_range_max));
			form.append('description', job.description  || "");
			form.append('benefits', job.benefits || "");
			form.append('how_to_apply', job.how_to_apply || "");
			form.append('application_url', job.application_url || "");
			form.append('application_phone_number', job.application_phone_number || "");
			form.append('tags', job.tags || "");

			try {
				const { data, status } = await axios.put<IJob>(
					`${baseURL}/products/${job.id}`,
					form,
					{
						headers: {
							"Content-Type": "multipart/form-data",
						},
					},
				);
				if (status === 200) {
					const idx = this.jobs.findIndex(
						elem => elem.id === job.id
					)
					this.jobs[idx] = { ...data };
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
		},

		async removeItem(id: number ) {
			this.loading = true;
			try {
				const { data, status } = await axios.delete<IJob>(
					`${baseURL}/products/${id}`,
					{
						headers: {
							"Content-Type": "application/json",
						},
					},
				);
				if (status === 200) {
					this.jobs = this.jobs.filter(
						elem => elem.id !== id
					);
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
	import.meta.hot?.accept(acceptHMRUpdate(useJobsStore, import.meta.hot))
}