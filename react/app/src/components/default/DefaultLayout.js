import React from "react";
import Container from "react-bootstrap/Container";
import DefaultNavbar from "./navbar/DefaultNavbar";
import "./content/Main.css"

const DefaultLayout = props => (
    <>
        <DefaultNavbar />
        <Container fluid className="default-layout-main">
            {props.children}
        </Container>
    </>
);

export default DefaultLayout;