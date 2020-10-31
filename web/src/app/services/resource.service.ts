import {Injectable} from '@angular/core';
import {ElasticsearchService} from './elasticsearch.service';
import {environment} from '../../environments/environment';
import {SearchParams, SearchResponse} from 'elasticsearch';
import {Resource, ResourceCount} from '../model/resource';
import {QueryParams} from '../model/es/es-query';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ResourceService {

  constructor(private es: ElasticsearchService,
              private http: HttpClient
  ) { }

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
      let search: Promise<SearchResponse<Resource>>;
      search = client.search(params);
      return search;
    });
  }

  fetchResourceCount(): Promise<ResourceCount> {
    return this.http
      .get<ResourceCount>(`${environment.api.host}:${environment.api.port}/resource/count`)
      .toPromise();
  }
}
