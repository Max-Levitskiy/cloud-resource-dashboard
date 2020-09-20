import { Injectable } from '@angular/core';
import {Client} from 'elasticsearch';
import * as elasticsearch from 'elasticsearch-browser';
import {environment} from '../../environments/environment';

@Injectable({
  providedIn: 'root'
})
export class ElasticsearchService {

  private client: Client;

  constructor() {
    if (!this.client) {
      this._connect();
    }
  }

  private _connect() {
    this.client = new elasticsearch.Client({
      host: environment.es.host + ':' + environment.es.port,
      log: environment.es.logLevel
    });
  }

  isAvailable(): Promise<void> {
    return this.client.ping({});
  }

  getClient(): Promise<Client> {
    return this.isAvailable().then(() => this.client);
  }
}
