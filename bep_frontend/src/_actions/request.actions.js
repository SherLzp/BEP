import { FETCH_STATUS } from '../_constants'
import { RequestActionTypes as types } from '../_constants/actions.types'
import { requestServices } from '../_services'

const showAllRequests = (json) => {
    return {
        type: types.SHOW_ALL_REQUESTS,
        payload: json.data
    }
}

const showAllRequestsAsync = () => {
    return dispatch => {
        dispatch(fetchBegin())
        requestServices.showAllRequests().then(
            json => {
                if (json.status === 0) {
                    dispatch(fetchSuccess())
                }
                dispatch(showAllRequests(json))
            }
        )
    }
}

const queryRequestsByUserId = (json) => {
    return {
        type: types.QUERY_REQUESTS_BY_USER_ID,
        payload: json.data
    }
}

const queryRequestsByUserIdAsync = (userId) => {
    return dispatch => {
        dispatch(fetchBegin())
        requestServices.queryRequestsByUserId(userId).then(
            json => {
                if (json.status === 0) {
                    dispatch(fetchSuccess())
                }
                dispatch(queryRequestsByUserId(json))
            }
        )
    }
}

const pushRequest = (json) => {
    return {
        type: types.PUSH_REQUEST,
        payload: json.data
    }
}

const pushRequestAsync = (userId, requirement, reward, expiredTime) => {
    return dispatch => {
        dispatch(fetchBegin())
        requestServices.pushRequest(userId, requirement, reward, expiredTime).then(
            json => {
                if (json.status === 0) {
                    dispatch(fetchSuccess())
                }
                dispatch(pushRequest(json))
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

export const requestActions = {
    showAllRequestsAsync,
    queryRequestsByUserIdAsync,
    pushRequestAsync,
}