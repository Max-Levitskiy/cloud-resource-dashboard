import {ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {ElasticsearchService} from '../../services/elasticsearch.service';
import {ApiService} from '../../services/api.service';

@Component({
  selector: 'app-status',
  templateUrl: './status.component.html',
  styleUrls: ['./status.component.scss']
})
export class StatusComponent implements OnInit {
  isEsConnected = false;
  esStatus: Status;
  apiStatus: Status;

  esStatusLoading = true;
  apiStatusLoading = true;

  constructor(
    private es: ElasticsearchService,
    private api: ApiService,
    private cd: ChangeDetectorRef) {
  }

  ngOnInit() {
    this.requestEsStatus();
    this.requestApiStatus();
  }

  private requestEsStatus() {
    this.es.isAvailable().then(() => {
      this.esStatus = Status.OK;
      this.isEsConnected = true;
    }, error => {
      this.esStatus = Status.ERROR;
      this.isEsConnected = false;
      console.error('Server is down', error);
    }).then(() => {
      this.cd.detectChanges();
    }).finally(() => this.esStatusLoading = false);
  }

  private requestApiStatus() {
    this.api
      .ping()
      .then(ok => ok ? this.apiStatus = Status.OK : this.apiStatus = Status.ERROR)
      .then((value) => console.log(value))
      .finally(() => this.apiStatusLoading = false);
  }

}

export enum Status {
  OK = 'OK', ERROR = 'ERR'
}
