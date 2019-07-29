import { userConstants } from '../_constants';

export function users(state = {}, action) {
  switch (action.type) {
    case userConstants.GETUSERINFO_REQUEST:
      return {
        loading: true
      };
    case userConstants.GETUSERINFO_SUCCESS:
      return {
        items: action.users
      };
    case userConstants.GETUSERINFO_FAILURE:
      return {
        error: action.error
      };
    default:
      return state
  }
}