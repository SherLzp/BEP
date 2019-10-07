import React, { Component } from 'react'
import { BrowserRouter, Route, Switch, Redirect } from 'react-router-dom'
import { Layout } from './_pages/layout'
import { HomeContent } from './_pages/homeContent'
import { PushRequestContent, PushResponseContent } from './_pages/request'
import { YourResponsesContent, ReceivedResponsesContent } from './_pages/response'
import { RecordsContent } from './_pages/records'
import { NotificationsContent } from './_pages/notifications'

import {
  AllRequestsContentContainer,
  YourRequestsContentContainer,
  YourResponsesContentContainer,
}
  from './_containers'


class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <Layout>
          <Switch>
            {/* Home */}
            <Route path="/home" component={HomeContent}></Route>
            {/* Request Part */}
            <Route path="/requests/allRequests" component={AllRequestsContentContainer}></Route>
            <Route path="/requests/yourRequests" component={YourRequestsContentContainer}></Route>
            <Route path="/requests/pushRequest" component={PushRequestContent}></Route>
            {/* Response Part */}
            <Route path="/responses/yourResponses" component={YourResponsesContentContainer}></Route>
            <Route path="/responses/receivedResponses" component={ReceivedResponsesContent}></Route>
            <Route path="/requests/pushResponse" component={PushResponseContent}></Route>
            {/* User Records Part*/}
            <Route path="/records" component={RecordsContent}></Route>
            {/* Notifications Part */}
            <Route path="/notifications" component={NotificationsContent}></Route>
            {/* User Info Part */}
            <Route path="/profile" component={HomeContent}></Route>
          </Switch >
        </Layout>
      </BrowserRouter >
    )
  }
}

export default App
