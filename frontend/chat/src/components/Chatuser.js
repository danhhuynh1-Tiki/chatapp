import Userchat from './user/Userchat';
import { Row , Col } from 'antd';
import React from  'react';


const Chatuser = () => {
    const ChatUser = () => {
        console.log('Chat user');
    }
    return (
        <>
            <Row style={{height:'70vh',overflow:'scroll'}}>
                <Col span={24} style={{textAlign:'center'}}>
                   <Userchat onClick={ChatUser}/>
                   <Userchat/>
                   <Userchat/>
                </Col>
            </Row>
        </>
    )
}

export default Chatuser;