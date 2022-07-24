import React from 'react';
import Sidebar from '../components/Sidebar';
import Message from '../components/Message';
import {Row,Col,Layout} from 'antd';
import '../assets/css/style.css';
// const { Sider,Header,Content,Footer } = Layout;
const Chat = () => {
    return (
        <>
            <Layout style={{ height: '100vh' }}>
                <Row style={{height: '100vh'}}>
                    <Col span={6} style={{backgroundColor:'white'}}>
                        <Sidebar />
                    </Col>
                    <Col span={18} style={{backgroundColor : 'White'}}>
                        <Message />
                    </Col>
                </Row>
            </Layout>
        </>
        
    )
}

export default Chat;