import { RequestActionTypes as types } from '../_constants/actions.types'
import { FETCH_STATUS } from '../_constants'

export function showAllRequestsReducer(state = [], action) {
    switch (action.type) {
        case types.SHOW_ALL_REQUESTS:
            return [
                ...state,
                ...action.payload
            ]
        default:
            return state
    }
}

export function queryRequestsByUserIdReducer(state = [], action) {
    switch (action.type) {
        case types.QUERY_REQUESTS_BY_USER_ID:
            return [
                ...state,
                ...action.payload
            ]
        default:
            return state
    }
}

export function pushRequest(state = [], action) {
    switch (action.type) {
        case types.PUSH_REQUEST:
            return 1
        default:
            return state
    }
}

export const requestReducer = (state = {}, action) => {
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
        case types.SHOW_ALL_REQUESTS:
            return {
                ...state,
                allRequestRecords: showAllRequestsReducer([], action)
            }
        case types.QUERY_REQUESTS_BY_USER_ID:
            return {
                ...state,
                yourRequestRecords: queryRequestsByUserIdReducer([], action)
            }
        case types.PUSH_REQUEST:
            return {
                ...state,
            }
        default:
            return state
    }
}