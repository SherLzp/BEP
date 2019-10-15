import { API_URLS } from '../_constants/api.url'
import { fetch_post_helper } from './utils'

const queryResponsesByUserId = (userId) => {
    const url = API_URLS.RESPONSE_QUERY_RESPONSES_BY_USER_ID_URL
    const body = JSON.stringify({
        userId: userId,
    })
    return fetch_post_helper(url, body)
}

const queryResponsesByRequestId = (requestId) => {
    const url = API_URLS.RESPONSE_QUERY_RESPONSES_BY_REQUEST_ID_URL
    const body = JSON.stringify({
        requestId: requestId,
    })
    return fetch_post_helper(url, body)
}

const queryReceivedResponsesByUserId = (userId) => {
    const url = API_URLS.RESPONSE_QUERY_RECEIVED_RESPONSES_BY_USER_ID_URL
    const body = JSON.stringify({
        userId: userId,
    })
    return fetch_post_helper(url, body)
}

const pushResponse = (requestId, userId, answer) => {
    const url = API_URLS.RESPONSE_PUSH_RESPONSE_URL
    const body = JSON.stringify({
        requestId: requestId,
        userId: userId,
        answer: answer,
    })
    return fetch_post_helper(url, body)
}

const acceptResponse = (userId, requestId, responseId) => {
    const url = API_URLS.USER_ACCPET_RESPONSE_URL
    const body = JSON.stringify({
        userId: userId,
        requestId: requestId,
        responseId: responseId,
    })
    return fetch_post_helper(url, body)
}


export const responseServices = {
    queryResponsesByUserId,
    queryResponsesByRequestId,
    queryReceivedResponsesByUserId,
    pushResponse,
    acceptResponse,
}