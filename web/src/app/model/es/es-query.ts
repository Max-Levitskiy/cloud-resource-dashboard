export interface QueryParams {
  from: number
  size: number
  query: string
  sort?: SortParam
}
export type SortParam = {[key in string]: string}
