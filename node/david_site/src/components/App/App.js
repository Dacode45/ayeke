/*! React Starter Kit | MIT License | http://www.reactstarterkit.com/ */

import React, { PropTypes } from 'react';
import styles from './App.less';
import withContext from '../../decorators/withContext';
import withStyles from '../../decorators/withStyles';
import AppActions from '../../actions/AppActions';
import AppStore from '../../stores/AppStore';
import Header from '../Header';
import ContentPage from '../ContentPage';
import ContactPage from '../ContactPage';
import LoginPage from '../LoginPage';
import RegisterPage from '../RegisterPage';
import NotFoundPage from '../NotFoundPage';
import Feedback from '../Feedback';
import Footer from '../Footer';

const pages = { ContentPage, ContactPage, LoginPage, RegisterPage, NotFoundPage };

@withContext
@withStyles(styles)
class App {

  static propTypes = {
    path: PropTypes.string.isRequired
  };

  ComponentDidMount() {
    window.addEventListener('popstate', this.handlePopState);
  }

  ComponentWillUnmount() {
    window.removeEventListener('popstate', this.handlePopState);
  }

  shouldComponentUpdate(nextProps) {
    return this.props.path !== nextProps.path;
  }

  render() {
    let Component;

    switch (this.props.path) {

      case '/':
      case '/about':
      case '/privacy':
        let page = AppStore.getPage(this.props.path);
        Component = React.createElement(pages[page.Component], page);
        break;

      case '/contact':
        Component = <ContactPage />;
        break;

      case '/login':
        Component = <LoginPage />;
        break;

      case '/register':
        Component = <RegisterPage />;
        break;
    }

    return Component ? (
      <div>
        <Header />
        {Component}
        <Feedback />
        <Footer />
      </div>
    ) : <NotFoundPage />;
  }

  handlePopState(event) {
    AppActions.navigateTo(window.location.pathname, {replace: !!event.state});
  }

}

export default App;
