import React from 'react';
import Sidebar from '../components/Sidebar';
import Message from '../components/Message';
import {Row,Col,Layout} from 'antd';
import '../assets/css/style.css';
const { Sider,Header,Content,Footer } = Layout;
const Chat = () => {
    return (
        <>
            <Layout style={{ height: '100vh' }}>
                <Row>
                    <Col span={6} style={{backgroundColor:'white',
                    height: '100vh', overflow: "hidden" }}>
                        <Sidebar />
                    </Col>
                    <Col span={18}>
                        <Message />
                    </Col>
                </Row>
            </Layout>
        </>
        
    )
}

export default Chat;