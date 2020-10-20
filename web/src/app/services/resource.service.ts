import {Injectable} from '@angular/core';
import {ElasticsearchService} from './elasticsearch.service';
import {environment} from '../../environments/environment';
import {SearchParams, SearchResponse} from 'elasticsearch';
import {Resource} from '../model/resource';
import {MatSort, MatSortable} from '@angular/material/sort/sort';
import {QueryParams} from '../model/es/es-query';

@Injectable({
  providedIn: 'root'
})
export class ResourceService {

  constructor(private es: ElasticsearchService) { }

  fetchResources(queryParams: QueryParams): Promise<SearchResponse<Resource>> {
    return this.es.getClient().then(client => {

      const params: SearchParams = {
        index: environment.es.index.resource.name,
        from: queryParams.from,
        size: queryParams.size,
      };
      params.body = {
        sort: queryParams.sort
      };

      if (queryParams.terms) {
        params.body.query = {term: {}};
        queryParams
          .terms
          .forEach((value, key) =>
            params.body.query.term[key] = {value});
      }

      if (queryParams.query) {
        params.q = queryParams.query;
      }
      return client.search(params);
    });
  }
}
