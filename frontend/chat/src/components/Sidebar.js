import React from 'react';
import User from './user/User';
import { Row, Col, Button,Tabs } from 'antd';
import Chatuser from '../components/Chatuser';
import Chatgroup from '../components/Chatgroup';
import '../assets/css/style.css';

const { TabPane } = Tabs;
const Sidebar = () => {
    return (
        <>
            <Row>
                <Col span={24} >
                    <User/>
                </Col>
            </Row>
            <Row style={{height : '80vh'}}>
                <Col span={24}>
                    <Tabs defaultActiveKey="1" centered>
                        <TabPane tab="Members" key="1">
                            <Chatuser/>
                        </TabPane>
                        <TabPane tab="Groups" key="2">
                            <Chatgroup/>
                        </TabPane>
                    </Tabs>
                </Col>
            </Row>
            <Row style ={{textAlign:'Center'}}>
                <Col span={24}>
                    <Button>Default Button</Button>
                </Col>
            </Row>
        </>
    )
}
export default Sidebar;