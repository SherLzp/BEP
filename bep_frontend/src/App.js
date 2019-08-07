import React, { Component } from 'react';
import { Router, Route } from 'react-router-dom';
import { connect } from 'react-redux';
import 'semantic-ui-css/semantic.min.css'
import { history } from './_helpers'
import { HomePageLayout } from './components/HomePageLayout'
import { alertActions } from './_actions';
import { LoginLayout } from './components/LoginLayout'
import { SignupLayout } from './components/SignupLayout'

class App extends Component {
    constructor(props) {
        super(props);
        history.listen((location, action) => {
            // clear alert on location change
            this.props.clearAlerts();
        });
    }
    render() {
        const { alert } = this.props;
        return (
            <div className="jumbotron">
                <div className="container">
                    <div className="col-sm-8 col-sm-offset-2">
                        {alert.message &&
                            <div className={`alert ${alert.type}`}>{alert.message}</div>
                        }
                        <Router history={history}>
                            <Route exact path="/" component={HomePageLayout} />
                            <Route path="/login" component={LoginLayout} />
                            <Route path="/signup" component={SignupLayout} />
                        </Router>
                    </div>
                </div>
            </div>
        );
    }
}

function mapState(state) {
    const { alert } = state;
    return { alert };
}

const actionCreators = {
    clearAlerts: alertActions.clear
};

const connectedApp = connect(mapState, actionCreators)(App);
export { connectedApp as App };
