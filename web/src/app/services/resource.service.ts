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
  ) {
  }

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
      if (queryParams.terms.size > 0) {
        params.body.query = {bool: {must: [] }};
        queryParams
          .terms
          .forEach((value, key) => {
            const term: any = {};
            term[key] = value;
            params.body.query.bool.must.push({term});
          });

      }

      if (queryParams.query) {
        params.q = queryParams.query;
      }
      let search: Promise<SearchResponse<Resource>>;
      search = client.search(params);
      return search;
    });
  }

  fetchCount(): Promise<ResourceCount> {
    return this.http
      .get<ResourceCount>(this.getApiUrlFor('/resource/count'))
      .toPromise();
  }

  fetchDistinctField(fieldName: string): Promise<string[]> {
    return this.http
      .get<string[]>(this.getApiUrlFor('/resource/distinct/' + fieldName))
      .toPromise();
  }

  private getApiUrlFor(path: string) {
    return `${environment.api.host}:${environment.api.port}${path}`;
  }
}

