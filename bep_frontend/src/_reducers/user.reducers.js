import { UserActionTypes as types } from '../_constants/actions.types'
import { FETCH_STATUS } from '../_constants'

export function queryUserBalanceReducer(state = [], action) {
    switch (action.type) {
        case types.QUERY_USER_BALANCE_BY_USER_ID:
            return action.payload
        default:
            return state
    }
}

export const userReducer = (state = {}, action) => {
    switch (action.type) {
        case FETCH_STATUS.FETCH_BEGIN:
            return {
                ...state,
                fetchStatus: action.payload
            }
        case FETCH_STATUS.FETCH_SUCCESS:
            return {
                ...state,
                fetchStatus: action.payload
            }
        case FETCH_STATUS.FETCH_FAIL:
            return {
                ...state,
                fetchStatus: action.payload
            }
        case types.QUERY_USER_BALANCE_BY_USER_ID:
            return {
                ...state,
                userBalance: queryUserBalanceReducer([], action)
            }
        default:
            return state
    }
}