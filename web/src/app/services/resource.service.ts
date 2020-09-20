import {Injectable} from '@angular/core';
import {ElasticsearchService} from './elasticsearch.service';
import {environment} from '../../environments/environment';
import {SearchParams, SearchResponse} from 'elasticsearch';
import {Resource} from '../model/resource';

@Injectable({
  providedIn: 'root'
})
export class ResourceService {

  constructor(private es: ElasticsearchService) { }

  fetchResources(from: number = 0, query?: string): Promise<SearchResponse<Resource>> {
    return this.es.getClient().then(client => {
      const params: SearchParams = {
        index: environment.es.index.resource.name,
        from,
      };
      if (query) {
        params.q = query;
      }
      return client.search(params);
    });
  }
}
