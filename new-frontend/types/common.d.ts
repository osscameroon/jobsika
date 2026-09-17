export type RootCommonsState = {
  data: string[];
	filter: string;
  loading: boolean,
  error: null | string
};

export interface ITooltip {
  id: number;
  name: string;
}

export interface INav {
  id: string;
  link: string;
  name: string;
}