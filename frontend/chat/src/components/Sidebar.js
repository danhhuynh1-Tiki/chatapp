import React from 'react';
import User from './user/User';
import { Row, Col, Button,Tabs } from 'antd';
import Chatuser from '../components/Chatuser';
import Chatgroup from '../components/Chatgroup';
import '../assets/css/style.css';

const { TabPane } = Tabs
let key = "1"
const Message = (key) => {
    if( key === "1" || key === ""){
        
        console.log("message user")
    }else{
        key = "2"
        console.log("Message group")
    }
    
}
const LoadMore = () => {
   if(key === "1"){
        console.log("Load more user chat")
   }else{
        console.log("load more group chat")
   }
}
const Memebers = () => {
    console.log("Memeber")
}
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
                    <Tabs defaultActiveKey="1" onChange={Message} centered>
                        <TabPane tab="Members" key="1" onClick={Memebers}>
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
                    <Button onClick={LoadMore}>Load More</Button>
                </Col>
            </Row>
        </>
    )
}
export default Sidebar;