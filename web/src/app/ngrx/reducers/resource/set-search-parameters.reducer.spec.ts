import { searchParamsReducer, initialState } from './set-search-parameters.reducer';

describe('SetSearchParameters Reducer', () => {
  describe('an unknown action', () => {
    it('should return the previous state', () => {
      const action = {} as any;

      const result = searchParamsReducer(initialState, action);

      expect(result).toBe(initialState);
    });
  });
});
