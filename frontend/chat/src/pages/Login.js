import React, { useRef } from 'react';
import { useNavigate } from 'react-router-dom';
// import '../../App.css';
// import 'antd/dist/antd.css';

// import { Button } from 'antd';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import { Col, Container,Row } from 'react-bootstrap';
import axios from 'axios';

function Login(){
    const username = useRef(null)
    const password = useRef(null)
    let navigate = useNavigate()
    const SendData = async() => {
        
        const data = JSON.stringify({ email: username.current.value,password:password.current.value })
        try{
            // Cannot use VPN ip : 172.20.10.3
            // Use VPN : 172.26.18.116
            const res = await axios.post('http://localhost:8080/api/auth/login',data,
                        {headers: { 'Content-Type': 'application/json' }})
            // alert(res)
            console.log(res)
            navigate("/chat")
        }catch(error){
            alert(`response : ${error}`)
        }
        
    }
    return  (
        <>
        <Container fluid>
            <Row>
                <Col xs={3}></Col>
            <Col xs={6}>
                <Form>
                    <Form.Group className="mb-3" controlId="formBasicEmail">
                        <Form.Label>Email address</Form.Label>
                        <Form.Control ref={username} type="email" placeholder="Enter email" />
                        {/* <Form.Text className="text-muted">
                        We'll never share your email with anyone else.
                        </Form.Text> */}
                    </Form.Group>

                    <Form.Group className="mb-3" controlId="formBasicPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control ref={password} type="password" placeholder="Password" />
                    </Form.Group>
                    {/* <Form.Group className="mb-3" controlId="formBasicCheckbox">
                        <Form.Check type="checkbox" label="Check me out" />
                    </Form.Group> */}
                    <Button variant="primary" onClick={SendData} type="button">
                        Submit
                    </Button>
                </Form>
            </Col>
            <Col xs={3}></Col>
            </Row>
            </Container>
        </>
    )
}

export default Login;