import React, { Component } from 'react'
import {
  Button,
  Container,
  Header,
  Icon,
  Menu,
  Responsive,
  Segment,
  Visibility
} from 'semantic-ui-react'
import { Link } from 'react-router-dom'
import { userActions } from '../_actions';
import { connect } from 'react-redux';

// Heads up!
// We using React Static to prerender our docs with server side rendering, this is a quite simple solution.
// For more advanced usage please check Responsive docs under the "Usage" section.
const getWidth = () => {
  const isSSR = typeof window === 'undefined'

  return isSSR ? Responsive.onlyTablet.minWidth : window.innerWidth
}

var footStyle = {
  padding: "15px"
}

class HomePageLayout extends Component {
  state = {}

  hideFixedMenu = () => this.setState({ fixed: false })
  showFixedMenu = () => this.setState({ fixed: true })

  render() {
    const { fixed } = this.state

    return (
      <Responsive getWidth={getWidth} minWidth={Responsive.onlyTablet.minWidth}>
        <Visibility
          once={false}
          onBottomPassed={this.showFixedMenu}
          onBottomPassedReverse={this.hideFixedMenu}
        >
          <Segment
            inverted
            textAlign='center'
            style={{ minHeight: 700, padding: '1em 0em' }}
            vertical
          >
            <Menu
              fixed={fixed ? 'top' : null}
              inverted={!fixed}
              pointing={!fixed}
              secondary={!fixed}
              size='large'
            >
              <Container>
                <Menu.Item key="1" active>
                  Home
                </Menu.Item>
                <Menu.Item position='right'>
                  {this.props.username !== 'Sher' ? <Link to="/login"><Button inverted={!fixed}>Login / Signup</Button></Link> : <p>{this.props.username}</p>}

                </Menu.Item>
              </Container>
            </Menu>
            <Container text>
              <Header
                as='h1'
                content='Blob Exchange Platform'
                inverted
                style={{
                  fontSize: '4em',
                  fontWeight: 'normal',
                  marginBottom: 0,
                  marginTop: '3em',
                }}
              />
              <Header
                as='h2'
                content='Exchange infomation anonymously'
                inverted
                style={{
                  fontSize: '1.7em',
                  fontWeight: 'normal',
                  marginTop: '1.5em',
                }}
              />
              <Button primary size='huge'>
                Get Started
          <Icon name='right arrow' />
              </Button>
            </Container>
          </Segment>
          <Segment inverted vertical style={{ padding: '5em 0em' }}>
            <Container>
              <Segment
                inverted
                textAlign='center'
                style={{ minHeight: 70, padding: '1em 0em' }}
                vertical
              >
                <Header as="h3" style={footStyle}>
                  Copyright © 2019.7 BEP Team.  All Rights Reserved.
            </Header>
              </Segment>
            </Container>
          </Segment>
        </Visibility>
      </Responsive>
    )
  }
}

function mapState(state) {
  const { username } = state.authentication
  return { username }
}

const connectHomePage = connect(mapState, null)(HomePageLayout)


export { connectHomePage as HomePageLayout }