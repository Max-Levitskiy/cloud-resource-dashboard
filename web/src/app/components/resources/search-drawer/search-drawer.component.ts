import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs';
import {FormGroupState} from 'ngrx-forms';
import {Store} from '@ngrx/store';
import {State} from '../../../reducers';
import {SearchParams} from '../../../forms/search-drawer-form';

@Component({
  selector: 'app-search-drawer',
  templateUrl: './search-drawer.component.html',
  styleUrls: ['./search-drawer.component.scss']
})
export class SearchDrawerComponent implements OnInit {

  form$: Observable<FormGroupState<SearchParams>>;

  constructor(private store: Store<State>) {
    this.form$ = this.store.select('searchForm');
  }

  ngOnInit(): void {
  }

}
