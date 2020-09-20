import {AfterViewInit, Component, EventEmitter, OnInit, ViewChild} from '@angular/core';
import {ResourceService} from '../../../services/resource.service';
import {Resource} from '../../../model/resource';
import {SearchResponse} from 'elasticsearch';
import {MatPaginator} from '@angular/material/paginator';
import {MatSort} from '@angular/material/sort';
import {merge} from 'rxjs';
import {catchError, map, startWith, switchMap} from 'rxjs/operators';
import {EsHits} from '../../../model/es/es-hits';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements AfterViewInit {
  displayedColumns: string[] = [
    // 'Id',
    'CloudProvider',
    'ResourceType',
    // 'AccountId',
    'Name',
    'Region',
    'CreationDate',
    // 'Tags',
  ];
  @ViewChild(MatPaginator) paginator: MatPaginator;
  filterEmitter = new EventEmitter();

  resources: EsHits<Resource>[] = [];
  total = 0;
  isLoadingResults = true;
  private query: string;

  constructor(private resourceService: ResourceService) {
  }

  ngAfterViewInit(): void {

    merge(this.paginator.page, this.filterEmitter)
      .pipe(
        startWith({}),
        switchMap(() => {
          this.isLoadingResults = true;
          return this.resourceService.fetchResources(
            this.paginator.pageIndex * this.paginator.pageSize,
            this.query
          );
        }),
        map(data => {
          this.isLoadingResults = false;
          this.total = data.hits.total.value;
          return data;
        }),
        // catchError(() => {
        //   this.isLoadingResults = false;
        //   // Catch if the GitHub API has reached its rate limit. Return empty data.
        //   // this.isRateLimitReached = true;
        //   return observableOf([]);
        // })
      )
      .subscribe(searchResponse => this.resources = searchResponse.hits.hits);
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.query = filterValue.trim().toLowerCase();
    this.filterEmitter.emit();
  }
}
