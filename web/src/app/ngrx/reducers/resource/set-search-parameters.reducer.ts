import { Action } from '@ngrx/store';


export const setSearchParametersFeatureKey = 'setSearchParameters';

export interface SearchParams {
  query: string;
}

export const initialState: SearchParams = {
  query: ''
};

export function searchParamsReducer(state = initialState, action: Action): SearchParams {
  switch (action.type) {

    default:
      return state;
  }
}
