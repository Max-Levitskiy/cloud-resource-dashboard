import {createFormGroupState, FormGroupState, onNgrxForms} from 'ngrx-forms';
import {ActionReducer, createReducer} from '@ngrx/store';
import {FormControlState} from 'ngrx-forms/src/state';

export interface SearchParams {
  service;
  region;
  projectId;
  query;
}

const formInitState: SearchParams = {
  service: null,
  region: null,
  projectId: null,
  query: null
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
