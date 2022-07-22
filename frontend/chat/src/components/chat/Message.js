import { Row,Col} from 'antd';
import React from 'react';

const Message = () => {
    return (
            <>
                <Row style={{marginBottom : '10px',borderRadius : '5px 5px'}}   >
                    <Col span={12} style={{borderRadius : '5px 5px',backgroundColor:'#ecf0f1',padding : '10px 10px'}}>Hello</Col>
                </Row>
            </>

    )
}

export default Message;