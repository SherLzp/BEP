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

export function queryReceivedResponsesByUserIdReducer(state = [], action) {
    switch (action.type) {
        case types.QUERY_RECEIVED_RESPONSES_BY_USER_ID:
            return [
                ...state,
                ...action.payload
            ]
        default:
            return state
    }
}

export function pushResponseReducer(state = [], action) {
    switch (action.type) {
        case types.PUSH_RESPONSE:
            return 1
        default:
            return state
    }
}

export function acceptResponseReducer(state = [], action) {
    switch (action.type) {
        case types.ACCPET_RESPONSE:
            return 1
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
        case types.QUERY_RECEIVED_RESPONSES_BY_USER_ID:
            return {
                ...state,
                requestAndResponsesRecords: queryReceivedResponsesByUserIdReducer([], action)
            }
        case types.PUSH_RESPONSE:
            return {
                ...state,
            }
        case types.ACCPET_RESPONSE:
            return {
                ...state,
            }
        default:
            return state
    }
}