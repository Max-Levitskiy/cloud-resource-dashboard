import * as fromSetSearchParameters from './set-search-parameters.actions';

describe('loadSetSearchParameters', () => {
  it('should return an action', () => {
    expect(fromSetSearchParameters.setSearchParameters().type).toBe('[SetSearchParameters] Set SearchParameters');
  });
});
