import { ResponseActionTypes as types } from '../_constants/actions.types'
import { FETCH_STATUS } from '../_constants'

export function queryUserResponseReducer(state = [], action) {
    switch (action.type) {
        case types.QUERY_RESPONSES_BY_USER_ID:
            return [
                ...state,
                ...action.payload
            ]
        default:
            return state
    }
}

export function queryRequestResponseReducer(state = [], action) {
    switch (action.type) {
        case types.QUERY_RESPONSES_BY_REQUEST_ID:
            return [
                ...state,
                ...action.payload
            ]
        default:
            return state
    }
}

export const responseReducer = (state = {}, action) => {
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
        case types.QUERY_RESPONSES_BY_USER_ID:
            return {
                ...state,
                yourResponseRecords: queryUserResponseReducer([], action)
            }
        case types.QUERY_RESPONSES_BY_REQUEST_ID:
            return {
                ...state,
                requestResponseRecords: queryRequestResponseReducer([], action)
            }
        default:
            return state
    }
}