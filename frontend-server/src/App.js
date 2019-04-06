import React, {Component} from 'react';
import PrivateRoute from './components/PrivateRoute'
import {Route,Switch} from 'react-router-dom'
import Login from './containers/Login/index'
// import Login from './containers/Login2/index'
import Index from './containers/Index/index'
import Manage from './containers/Manage/index'
import './styles/App.css'
import './assets/font/iconfont.css'


class App extends Component {
  render() {
    return (
      <Switch>
        <PrivateRoute path='/manage' component={Manage}/>
        <Route path='/login' component={Login}/>
        <Route path='/' component={Index}/>
      </Switch>
    )
  }
}

export default App;
