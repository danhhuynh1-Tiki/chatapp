import React from 'react';
import {Col,Row,Avatar,Button,Badge} from 'antd';
import { UserOutlined } from '@ant-design/icons';
// import UserService from '../../services/UsersService';

const styleUserChat = {
    borderRadius : '3px 3px'
}
const UserChat = (props) => {
    let status = (props.user.status === 1 ) ? true : false 
    return (    
        <>
             <Row className="Userchat" style={styleUserChat}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <span>
                <Badge dot={status}>
                    <Avatar size={50} shape="square" icon={<UserOutlined />} />
                </Badge>
                </span>
                </Col>
                <Col span={10}>
                   <h7>{props.user.email}</h7>
                </Col>
                <Col span={9}>
                <Button type="danger"></Button>
                </Col>
            </Row>
        </>
    )
}
export default UserChat;