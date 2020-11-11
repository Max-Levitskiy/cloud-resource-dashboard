import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {FormGroupState} from 'ngrx-forms';
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
  distinctServices$: Promise<string[]>;
  distinctRegions$: Promise<string[]>;

  constructor(private store: Store<State>,
              private resourceService: ResourceService
              ) {
  }

  ngOnInit(): void {
    this.form$ = this.store.select('searchForm');
    this.distinctServices$ = this.resourceService.fetchDistinctField('service');
    this.distinctRegions$ = this.resourceService.fetchDistinctField('region');
  }
}
