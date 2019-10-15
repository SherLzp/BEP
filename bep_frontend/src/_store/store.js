import { createStore, combineReducers, applyMiddleware, compose } from 'redux'
import thunkMiddleware from 'redux-thunk'
import { requestReducer, responseReducer, userReducer } from '../_reducers'

let initialState = {
    "request": {
        "allRequestRecords": [],
        "yourRequestRecords": [],
    },
    "response": {
        "yourResponseRecords": [],
        "requestResponseRecords": [],
        "requestAndResponsesRecords":[],
    },
    "user": {
        "userBalance": 0,
    },
}

const rootReducer = combineReducers({
    request: requestReducer,
    response: responseReducer,
    user: userReducer,
})

const middlewares = [thunkMiddleware]

// 配置Redux-devtools
const composeEnhancers =
    typeof window === 'object' &&
        window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ ?
        window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__({
            // Specify extension’s options like name, actionsBlacklist, actionsCreators, serialize...
        }) : compose;

const enhancer = composeEnhancers(
    applyMiddleware(...middlewares),
    // other store enhancers if any
);

export const store = createStore(
    rootReducer,
    initialState,
    enhancer
)
