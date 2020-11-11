import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {FormGroupState, FormState} from 'ngrx-forms';
import {Store} from '@ngrx/store';
import {State} from '../../../reducers';
import {SearchParams} from '../../../forms/search-drawer-form';
import {ResourceService} from '../../../services/resource.service';

@Component({
  selector: 'app-search-drawer',
  templateUrl: './search-drawer.component.html',
  styleUrls: ['./search-drawer.component.scss'],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class SearchDrawerComponent implements OnInit {

  form$: Observable<FormGroupState<SearchParams>>;
  distinctCloudProviders$: Promise<string[]>;
  distinctServices$: Promise<string[]>;
  distinctRegions$: Promise<string[]>;
  distinctProjects$: Promise<string[]>;

  constructor(private store: Store<State>,
              private resourceService: ResourceService
              ) {
  }

  ngOnInit(): void {
    this.form$ = this.store.select('searchForm');
    this.distinctCloudProviders$ = this.resourceService.fetchDistinctField('cloudProvider');
    this.distinctServices$ = this.resourceService.fetchDistinctField('service');
    this.distinctRegions$ = this.resourceService.fetchDistinctField('region');
    this.distinctProjects$ = this.resourceService.fetchDistinctField('projectId');
  }
}
