import Userchat from './user/Userchat';
import React from 'react';
import { Row, Col} from 'antd';

const Chatgroup = () => {
    return (
        <>
            <Row>
                <Col span={24} style={{textAlign:'center'}}>
                    <Userchat/>
                </Col>
            </Row>
        </>
    )
}
export default Chatgroup;