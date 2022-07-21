import Userchat from './user/Userchat';
import { Row , Col } from 'antd';
import React from  'react';

const Chatuser = () => {
    return (
        <>
            <Row>
                <Col span={24} style={{textAlign:'center'}}>
                   <Userchat/>
                   <Userchat/>
                   <Userchat/>
                   <Userchat/>
                   <Userchat/>
                </Col>
            </Row>
        </>
    )
}

export default Chatuser;