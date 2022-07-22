import React from 'react';
import { Row, Col, Avatar,Button,Badge} from 'antd';
import { UserOutlined } from '@ant-design/icons';

const styleGroup ={
    marginTop : '5px',
    marginBottom : '5px',
    borderRadius : '3px 3px'
}

const ChatGroup = () => {
    console.log("chat group")
}
const Groupmessage = () => {
    return (
            <>
             <Row style={styleGroup} className="GroupChat" onClick={ChatGroup}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <span>
                {/* <Badge dot> */}
                    <Avatar size={50} shape="square" icon={<UserOutlined />} />
                {/* </Badge> */}
                </span>
                </Col>
                <Col span={10}>
                   <h7>Chem gio</h7>
                </Col>
                <Col span={9}>
                <Button type="danger"></Button>
                </Col>
            </Row>
            </>

    )
}

export default Groupmessage;