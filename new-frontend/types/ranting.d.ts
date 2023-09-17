export interface IRating{
	salary_id: number;
  company_id: number;
  company_rating_id: number;
  rating: number;
  salary: number;
  company_name: string;
  seniority: string;
  comment: string;
  job_title: string;
  country: string;
  city: string;
  createdat: string;
}

export type GetAverageResponse = {
  rating: number;
	salary: number;
}

export type GetRatingsResponse = {
  hits: IRating[];
	page: number;
  limit: number;
	nbHits: number;
	offset: number;
}

export type RootRatingsState = {
  ratings: IJob[];
  page: number;
  limit: number;
	offset: number;
	nbHits: number;
  average: number;
  averageStars: number;
  loading: boolean,
  error: null | string
};

export type GetRatingsParams = {
  page?: number;
  limit?: number;
  company?: string | null;
  jobtitle?: string | null;
  seniority?: string | null;
  city?: string | null;
};

export type GetAverageParams = {
  company?: string | null;
  jobtitle?: string | null;
  seniority?: string | null;
  city?: string | null;
};