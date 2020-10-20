import {createFormGroupState, FormGroupState, onNgrxForms} from 'ngrx-forms';
import {ActionReducer, createReducer} from '@ngrx/store';

export interface SearchParams {
  service;
  query;
}

const formInitState: SearchParams = {
  service: '',
  query: ''
};

const SEARCH_DRAWER_FORM_ID = 'search-drawer-form';
export const initialSearchFormState: FormGroupState<SearchParams> = createFormGroupState<SearchParams>(
  SEARCH_DRAWER_FORM_ID,
  formInitState
);

export const searchFormReducer: ActionReducer<FormGroupState<SearchParams>> = createReducer(
  initialSearchFormState,
  onNgrxForms()
);
