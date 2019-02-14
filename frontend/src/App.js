import React, { Component } from 'react';
import './App.css';
import Login from './Login';
import { Route, Redirect } from 'react-router-dom';
import Header from './components/Header';
import Student from './components/student/home';
import Classes from './components/student/subjects';
import Kcpe from './components/student/kcpe';
import Register from './Register';
import { isAuthenticated } from './lib/Auth'
import Subjects from './components/student/subjects';

const StudentRoute = ({ component: Component, ...rest }) => (
  <Route
    {...rest}
    render={props =>
      isAuthenticated() ? (
        <Component {...props} />
      ) : (
        <Redirect
          to={{
            pathname: "/login",
            state: { from: props.location }
          }}
        />
      )
    }
  />
);

class App extends Component {
  render() {
    return (
      <div className="App">
       <Header />

          <StudentRoute path="/student" component={Student} />
          <StudentRoute path="/subjects" component={Subjects} />
          <StudentRoute path="/kcpe" component={Kcpe} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Register} />
      </div>
    );
  }
}

export default App;
