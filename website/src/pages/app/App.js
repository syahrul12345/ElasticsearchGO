import React from 'react';
import { Grid } from '@material-ui/core/'
import Login from '../login'
import DashBoard from '../dashboard'
import { BrowserRouter as Router,
        Switch ,
        Route} from 'react-router-dom'
function App() {
  return (
    <Router>
      <Switch>
        <Route path="/dashboard">
          <DashBoard/>
        </Route>
        <Route exact path="/">
          <Login/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
