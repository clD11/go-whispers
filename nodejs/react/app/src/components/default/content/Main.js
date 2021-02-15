import React, { Component } from "react";
import {Container, Row, Col} from "react-bootstrap";
import LineGraph from "../charts/LineGraph";
import ResizableBox from "./ResizableBox";
import "./Main.css";

class Main extends Component {
    render() {
        return (
            <Container fluid className="content-main">
                <Row className="">
                    <Col sm={12}>
                        <h1>go-whispers</h1>>
                    </Col>
                </Row>
                <Row className="">
                    <Col sm={12}>
                        <ResizableBox>
                            <LineGraph />
                        </ResizableBox>
                    </Col>
                </Row>
            </Container>
        );
    }
}

export default Main;