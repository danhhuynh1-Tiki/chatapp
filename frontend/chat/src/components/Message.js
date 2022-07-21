import React from 'react';
import { Row,Col,Input,Button } from 'antd';
import Message from './chat/Message';
import Usermessage from './chat/Usermessage';


const message = () =>{
    return (
        <>
                <Row style={{height : '90vh'}}>
                    <Col span={24}>
                            <Message/>
                            <Usermessage/>
                    </Col>
                </Row>

                
                <Row style={{width : '100%',textAlign:'center'}}>
                    <Col span={20}>
                    <Input placeholder="Basic usage" style={{height:'50px'}}/>
                    </Col>
                    <Col span={4}>
                        <Button type="primary" style={{height:'50px'}}>Send</Button>
                    </Col>
                </Row>
        </>
    )
}

export default message;