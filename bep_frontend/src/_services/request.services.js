import { API_URLS } from '../_constants/api.url'
import { fetch_get_helper, fetch_post_helper } from './utils'

const showAllRequests = () => {
    const url = API_URLS.REQUEST_SHOW_ALL_REQUESTS_URL
    return fetch_get_helper(url)
}

const queryRequestsByUserId = (userId) => {
    const url = API_URLS.REQUEST_QUERY_REQUESTS_BY_USER_ID_URL
    const body = JSON.stringify({
        userId: userId,
    })
    return fetch_post_helper(url, body)
}

const pushRequest = (userId, requirement, reward, expiredTime) => {
    const url = API_URLS.REQUEST_PUSH_REQUEST_URL
    const body = JSON.stringify({
        userId: userId,
        requirement: requirement,
        reward: reward,
        expiredTime: expiredTime,
    })
    return fetch_post_helper(url, body)
}


export const requestServices = {
    showAllRequests,
    queryRequestsByUserId,
    pushRequest,
}