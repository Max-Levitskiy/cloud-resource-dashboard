import {AfterViewInit, ChangeDetectionStrategy, Component, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {ResourceService} from '../../../services/resource.service';
import {Resource} from '../../../model/resource';
import {MatPaginator} from '@angular/material/paginator';
import {merge, Observable, of, Subscription} from 'rxjs';
import {catchError, map, startWith, switchMap, tap} from 'rxjs/operators';
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
  styleUrls: ['./list.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ListComponent implements OnInit, OnDestroy, AfterViewInit {
  displayedColumns: string[] = [
    'CloudProvider',
    'Service',
    'ProjectId',
    'ResourceId',
    'Region',
    'CreationDate',
    'Tags',
  ];
  @ViewChild(MatPaginator) paginator: MatPaginator;
  @ViewChild(MatSort) sort: MatSort;

  resources$: Observable<EsHits<Resource>[]>;
  total = 0;
  isLoadingResults = true;
  form$: Observable<FormGroupState<SearchParams>>;
  private controls: FormGroupControls<SearchParams>;
  private formSubscription: Subscription;

  constructor(private resourceService: ResourceService, private store: Store<State>) {
  }

  ngOnInit(): void {
    this.form$ = this.store.select('searchForm');
  }

  ngOnDestroy(): void {
    this.formSubscription.unsubscribe();
  }

  ngAfterViewInit(): void {
    this.resources$ = merge(this.paginator.page, this.sort.sortChange, this.form$)
      .pipe(
        startWith({}),
        tap(() => this.isLoadingResults = true),
        map(() => {
          const queryParams: QueryParams = {
            query: this.controls.query.value,
            size: this.paginator.pageSize,
            from: this.paginator.pageIndex * this.paginator.pageSize,
          };
          queryParams.terms = new Map<string, string>();
          if (this.controls.service.value) {
            queryParams.terms.set('Service', this.controls.service.value);
          }
          if (this.controls.region.value) {
            queryParams.terms.set('Region', this.controls.region.value);
          }
          if (this.controls.projectId.value) {
            queryParams.terms.set('ProjectId', this.controls.projectId.value);
          }
          if (this.sort.active && this.sort.direction) {
            const sortParam: SortParam = {};
            sortParam[this.sort.active] = this.sort.direction;
            queryParams.sort = sortParam;
          }
          return queryParams;
        }),
        switchMap((queryParams) => this.resourceService.fetchResources(queryParams)),
        map(data => {
          this.isLoadingResults = false;
          // TODO: Remove ts-ignore after https://github.com/DefinitelyTyped/DefinitelyTyped/pull/47812 will be merged
          // @ts-ignore
          this.total = data.hits.total.value;
          return data.hits.hits;
        }),
        catchError(() => {
          this.isLoadingResults = false;
          return of([]);
        })
      );
    this.formSubscription = this.form$.subscribe(form => this.controls = form.controls);
  }

}
