import React from 'react';
import {Row,Col,Avatar,Button, Space} from 'antd';
import { UserOutlined } from '@ant-design/icons';
// import '../../assets/css/style.css';
const User = () =>{
    return (
        <>
            <Row className="User" style={{marginTop : '10px'}}>
                <Col span={5}>
                {/* <Avatar src={<Image src={avatar} style={{ width: 32 }} />} /> */}
                <Avatar shape="square" size={50} icon={<UserOutlined />} />
                </Col>
                <Col span={10}>
                    <Space align="center">
                        <h5>Danh Huynh</h5>
                   </Space>
                </Col>
                <Col span={9}>
                <Button type="danger">Logout</Button>
                </Col>
            </Row>
        </>
    )
}

export default User;