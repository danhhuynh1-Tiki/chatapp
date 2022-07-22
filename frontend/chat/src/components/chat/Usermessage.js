import { Row,Col} from 'antd';
import React from 'react';

const Usermessage = () => {
    return (
            <>
                <Row style = {{borderRadius : '5px 5px',marginBottom : '5px'}}>
                        <Col span={12} style={{height:'50px'}}></Col>
                        <Col span={12} style={{backgroundColor:'skyblue',borderRadius : '5px 5px',padding:'5px 5px'}}> versions of Lorem Ipsum.</Col>
                </Row>
            </>

    )
}

export default Usermessage;