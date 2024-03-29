import { Row,Col} from 'antd';
import React from 'react';

const Message = (props) => {
    return (
            <>
                <Row>
                    <Col span={12}><p style={{color:'#ced6e0',fontSize:'15px'}}>{props.message.email}</p></Col>
                </Row>
                <Row style={{marginBottom : '10px',borderRadius : '5px 5px'}}   >
                    
                    <Col span={12} style={{borderRadius : '5px 5px',backgroundColor:'#ecf0f1',padding : '10px 10px'}}>{props.message.content}</Col>
                </Row>
            </>

    )
}

export default Message;