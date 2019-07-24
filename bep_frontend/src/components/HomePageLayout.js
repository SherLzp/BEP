import PropTypes from 'prop-types'
import React, { Component } from 'react'
import {
  Button,
  Container,
  Header,
  Icon,
  Menu,
  Responsive,
  Segment,
  Visibility,
} from 'semantic-ui-react'
import { Route, Link } from 'react-router-dom'
import LoginForm from './LoginLayout';

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

/* eslint-disable react/no-multi-comp */
/* Heads up! HomepageHeading uses inline styling, however it's not the best practice. Use CSS or styled components for
 * such things.
 */
const HomepageHeading = ({ mobile }) => (
  <Container text>
    <Header
      as='h1'
      content='Blob Exchange Platform'
      inverted
      style={{
        fontSize: mobile ? '2em' : '4em',
        fontWeight: 'normal',
        marginBottom: 0,
        marginTop: mobile ? '1.5em' : '3em',
      }}
    />
    <Header
      as='h2'
      content='Exchange infomation anonymously'
      inverted
      style={{
        fontSize: mobile ? '1.5em' : '1.7em',
        fontWeight: 'normal',
        marginTop: mobile ? '0.5em' : '1.5em',
      }}
    />
    <LoginForm showBtn={<Button primary size='huge'>
      Get Started
      <Icon name='right arrow' />
    </Button>}
    />
  </Container>
)

HomepageHeading.propTypes = {
  mobile: PropTypes.bool,
}

/* Heads up!
 * Neither Semantic UI nor Semantic UI React offer a responsive navbar, however, it can be implemented easily.
 * It can be more complicated, but you can create really flexible markup.
 */
class DesktopContainer extends Component {
  state = {}

  hideFixedMenu = () => this.setState({ fixed: false })
  showFixedMenu = () => this.setState({ fixed: true })

  render() {
    const { children } = this.props
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
                  <Link to="/home">Home</Link>
                </Menu.Item>
                {/* <Menu.Item position='right'>
                    <Button as={Link} to="/login" display={"none"} inverted={!fixed} onClick={this.onLogin}>
                        Log in
                    </Button>
                    <Button as={Link} to="/signup" inverted={!fixed} onClick={this.onSignup} primary={fixed} style={{ marginLeft: '0.5em' }}>
                        Sign up
                    </Button>
                </Menu.Item> */}
              </Container>
            </Menu>
            <HomepageHeading />
          </Segment>
        </Visibility>
        {children}
      </Responsive>
    )
  }
}

DesktopContainer.propTypes = {
  children: PropTypes.node,
}

const ResponsiveContainer = ({ children }) => (
  <div>
    <DesktopContainer>{children}</DesktopContainer>
  </div>
)

ResponsiveContainer.propTypes = {
  children: PropTypes.node,
}

const HomepageLayout = () => (
  <ResponsiveContainer>
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
  </ResponsiveContainer>
)
export default HomepageLayout