import {Explanation} from 'elasticsearch';

export class EsHits<T> {
  // tslint:disable:variable-name
  _index: string;
  _type: string;
  _id: string;
  _score: number;
  _source: T;
  _version?: number;
  _explanation?: Explanation;
  fields?: any;
  highlight?: any;
  inner_hits?: any;
  matched_queries?: string[];
  sort?: string[];
}
