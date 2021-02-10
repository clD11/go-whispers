import React, { Component } from "react"
import { Route } from "react-router-dom"

const AppRouter = ({ component: Component, layout: Layout, ...rest }) => (
    <Route {...rest} render={props => (
        <Layout>
            <Component {...props} />
        </Layout>
    )} />
);

export default AppRouter