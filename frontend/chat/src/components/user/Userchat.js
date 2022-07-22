import React from 'react';
import {Col,Row,Avatar,Button,Badge} from 'antd';
import { UserOutlined } from '@ant-design/icons';

const styleUserChat = {
    borderRadius : '3px 3px'
}
const UserChat = () => {
    return (    
        <>
             <Row className="Userchat" style={styleUserChat}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <span>
                <Badge dot={true}>
                    <Avatar size={50} shape="square" icon={<UserOutlined />} />
                </Badge>
                </span>
                </Col>
                <Col span={10}>
                   <h7>Danh Huynh</h7>
                </Col>
                <Col span={9}>
                <Button type="danger"></Button>
                </Col>
            </Row>
        </>
    )
}
export default UserChat;