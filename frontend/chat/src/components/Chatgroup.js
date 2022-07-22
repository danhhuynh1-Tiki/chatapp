import Groupchat from './user/Groupchat';
import React from 'react';
import { Row, Col, Button} from 'antd';




const Chatgroup = () => {
    return (
        <>
            <Row style={{textAlign:'center',marginBottom : '1px'}}>
                <Col span={24}>
                    <Button>+</Button>
                </Col>
            </Row>
            <Row style={{ height : '65vh', overflow:'scroll'}}>
                <Col span={24} style={{textAlign:'center'}}>
                    <Groupchat/>
                </Col>
            </Row>
           
        </>
    )
}
export default Chatgroup;