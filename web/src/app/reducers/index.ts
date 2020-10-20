import {ActionReducerMap, MetaReducer} from '@ngrx/store';
import {environment} from '../../environments/environment';
import {FormGroupState} from 'ngrx-forms';
import {searchFormReducer, SearchParams} from '../forms/search-drawer-form';


export interface State {
  searchForm: FormGroupState<SearchParams>;
}

export const reducers: ActionReducerMap<State> = {
  searchForm: searchFormReducer
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];
