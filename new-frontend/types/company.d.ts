export interface ICompany{
	id: number;
  name: string;
  rating: number;
  createdat: string;
  updatedat: string;
}

export type RootCompaniesState = {
  companies: ICompany[];
  filtercompany: string,
  loading: boolean,
  error: null | string
};