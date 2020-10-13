import {Component, OnInit, Output} from '@angular/core';
import {Store} from '@ngrx/store';
import {SearchParams} from '../../../ngrx/reducers/resource/set-search-parameters.reducer';
import {Observable} from 'rxjs';

@Component({
  selector: 'app-search-drawer',
  templateUrl: './search-drawer.component.html',
  styleUrls: ['./search-drawer.component.scss']
})
export class SearchDrawerComponent implements OnInit {

  @Output() private query: string;

  constructor() {
  }

  ngOnInit(): void {
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    const query = filterValue.trim().toLowerCase();
    if (query !== this.query) {
      this.query = query;
      // this.filterEmitter.emit();
    }
  }
}
