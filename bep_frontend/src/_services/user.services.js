import { API_URLS } from '../_constants/api.url'
import { fetch_get_helper, fetch_post_helper } from './utils'

const getUserBalance = (userId) => {
    const url = API_URLS.USER_QUERY_USER_BALANCE_BY_USER_ID_URL
    const body = JSON.stringify({
        userId: userId,
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

export const requestServices = {
    getUserBalance,
    acceptResponse,
}