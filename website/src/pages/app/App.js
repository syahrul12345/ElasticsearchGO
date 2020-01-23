import React,{useState} from 'react'
import DashBoard from '../dashboard'
import Create from '../createaccount'
import Login from '../login'
import { BrowserRouter as Router,
        Switch ,
        Route} from 'react-router-dom'
function App() {
  // Set the global state
  const [state,setState] = useState({
    token:''
  })
  
  return (
    <Router>
      <Switch>
        <Route exact path="/create">
          <Create globalState={state} setState={setState}/>
        </Route>
        <Route exact path="/dashboard">
          <DashBoard globalState={state} setState={setState}/>
        </Route>
        <Route exact path="/">
          <Login globalState={state} setState={setState}/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
