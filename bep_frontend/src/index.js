import React from 'react';
import { render } from 'react-dom';
import './index.css';
import * as serviceWorker from './serviceWorker';
import { store } from './_helpers'
import { Provider } from 'react-redux'
import { configureFakeBackend } from './_helpers';
import { App } from './App'
configureFakeBackend();

render(
    <Provider store={store}>
        <App />
    </Provider>,
    document.getElementById('root')
)

serviceWorker.unregister();
