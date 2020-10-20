export interface QueryParams {
  from: number;
  size: number;
  query?: string;
  terms?: Map<string, string>;
  sort?: SortParam;
}
export type SortParam = {[key in string]: string};
