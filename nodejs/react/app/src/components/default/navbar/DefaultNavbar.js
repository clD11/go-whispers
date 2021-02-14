import React from 'react';
import { Nav, Navbar } from 'react-bootstrap';
import "./DefaultNavbar.css";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome/index'
import { faBars } from '@fortawesome/free-solid-svg-icons/index'

const DefaultNavbar = () => (
    <Navbar className="main-navbar" expand="md">
        <Navbar.Toggle aria-controls="basic-navbar-nav" />
        <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="navbar-nav ml-auto">
                <Nav.Link href="#home"></Nav.Link>
            </Nav>
        </Navbar.Collapse>
        <Navbar.Brand href="#/"><FontAwesomeIcon className="navbar-bars-icon" icon={faBars} /></Navbar.Brand>
    </Navbar>
);

export default DefaultNavbar;