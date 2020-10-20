import {AfterViewInit, Component, EventEmitter, ViewChild} from '@angular/core';
import {ResourceService} from '../../../services/resource.service';
import {Resource} from '../../../model/resource';
import {MatPaginator} from '@angular/material/paginator';
import {merge, Observable} from 'rxjs';
import {map, startWith, switchMap} from 'rxjs/operators';
import {EsHits} from '../../../model/es/es-hits';
import {MatSort} from '@angular/material/sort';
import {QueryParams, SortParam} from '../../../model/es/es-query';
import {Store} from '@ngrx/store';
import {FormGroupControls, FormGroupState} from 'ngrx-forms';
import {State} from '../../../reducers';
import {SearchParams} from '../../../forms/search-drawer-form';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.scss']
})
export class ListComponent implements AfterViewInit {
  displayedColumns: string[] = [
    'CloudProvider',
    'Service',
    // 'AccountId',
    'ResourceId',
    'Region',
    'CreationDate',
    'Tags',
  ];
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;
  filterEmitter = new EventEmitter();

  resources: EsHits<Resource>[] = [];
  total = 0;
  isLoadingResults = true;
  private query: string;
  form$: Observable<FormGroupState<SearchParams>>;
  private controls: FormGroupControls<SearchParams>;

  constructor(private resourceService: ResourceService, private store: Store<State>) {
    this.form$ = this.store.select('searchForm');
  }

  ngAfterViewInit(): void {
    this.form$.subscribe(form => this.controls = form.controls);
    merge(this.paginator.page, this.sort.sortChange, this.form$)
      .pipe(
        startWith({}),
        switchMap((value) => {
          this.isLoadingResults = true;
          const queryParams: QueryParams = {
            query: this.controls.query.value,
            size: this.paginator.pageSize,
            from: this.paginator.pageIndex * this.paginator.pageSize,
          };
          if (this.controls.service.value) {
            queryParams.terms = new Map<string, string>();
            queryParams.terms.set('Service', this.controls.service.value);
          }
          if (this.sort.active && this.sort.direction) {
            const sortParam: SortParam = {};
            sortParam[this.sort.active] = this.sort.direction;
            queryParams.sort = sortParam;
          }
          return this.resourceService.fetchResources(
            queryParams
          );
        }),
        map(data => {
          this.isLoadingResults = false;
          // TODO: Remove ts-ignore after https://github.com/DefinitelyTyped/DefinitelyTyped/pull/47812 will be merged
          // @ts-ignore
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
    // const filterValue = (event.target as HTMLInputElement).value;
    // const query = filterValue.trim().toLowerCase();
    // if (query !== this.query) {
    //   this.query = query;
    //   this.filterEmitter.emit();
    // }
  }
}
