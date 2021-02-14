import React from "react";
import {HashRouter, Switch} from "react-router-dom";
import AppRouter from "./AppRouter";
import DefaultLayout from "../default/DefaultLayout"
import Main from "../default/content/Main"

const App = () => (
    <HashRouter>
      <Switch>
        <AppRouter exact path="/" layout={DefaultLayout} component={Main}/>
      </Switch>
    </HashRouter>
);

export default App;
