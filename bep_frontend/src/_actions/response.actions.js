import { FETCH_STATUS } from '../_constants'
import { ResponseActionTypes as types } from '../_constants/actions.types'
import { responseServices } from '../_services'

const queryResponsesByUserId = (json) => {
    return {
        type: types.QUERY_RESPONSES_BY_USER_ID,
        payload: json.data
    }
}

const queryResponsesByUserIdAsync = (userId) => {
    return dispatch => {
        dispatch(fetchBegin())
        responseServices.queryResponsesByUserId(userId).then(
            json => {
                if (json.status === 0) {
                    dispatch(fetchSuccess())
                }
                dispatch(queryResponsesByUserId(json))
            }
        )
    }
}

const queryResponsesByRequestId = (json) => {
    return {
        type: types.QUERY_RESPONSES_BY_REQUEST_ID,
        payload: json.data
    }
}

const queryResponsesByRequestIdAsync = (requestId) => {
    return dispatch => {
        dispatch(fetchBegin())
        responseServices.queryResponsesByRequestId(requestId).then(
            json => {
                if (json.status === 0) {
                    dispatch(fetchSuccess())
                }
                dispatch(queryResponsesByRequestId(json))
            }
        )
    }
}

const fetchBegin = () => {
    return {
        type: FETCH_STATUS.FETCH_BEGIN,
        payload: "FETCH_BEGIN",
    }
}
const fetchSuccess = () => {
    return {
        type: FETCH_STATUS.FETCH_SUCCESS,
        payload: "FETCH_SUCCESS",
    }
}
const fetchFail = () => {
    return {
        type: FETCH_STATUS.FETCH_FAIL,
        payload: "FETCH_FAIL"
    }
}

export const responseActions = {
    queryResponsesByUserIdAsync,
    queryResponsesByRequestIdAsync,
}