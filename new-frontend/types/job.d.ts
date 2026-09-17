export interface IJob{
	id: number,
	createdat: string;
	updatedat: string;
	company_name: string;
	job_title: string;
	is_remote: boolean;
	city: string;
	country: string;
	department: string;
	salary_range_min: number;
	salary_range_max: number;
	description: string;
	benefits: string;
	how_to_apply: string;
	application_url: string;
	application_email_address: string;
	application_phone_number: string;
	tags?: string;
}

export type GetJobsResponse = {
  hits: IJob[];
	page: number;
  limit: number;
	nbHits: number;
	offset: number;
}

export type RootJobsState = {
  jobs: IJob[];
  page: number;
  limit: number;
	offset: number;
	nbHits: number;
  loading: boolean,
  error: null | string
};

export type GetJobsParams = {
  page: number;
  limit: number;
};