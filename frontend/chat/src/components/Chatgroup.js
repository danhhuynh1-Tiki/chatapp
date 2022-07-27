import Groupchat from './user/Groupchat';
import React, {useRef, useState} from 'react';
import { Row, Col, Button,Modal,Input} from 'antd';




const Chatgroup = () => {
    
    const Name = useRef(null)
    const Memebers = useRef(null)

    const [isModalVisible, setIsModalVisible] = useState(false);

    const showModal = () => {
        setIsModalVisible(true);
    };
  
    const handleOk = () => {
        console.log(Name.current.value)
        console.log(Memebers.current.value)
        setIsModalVisible(false);
    };
  
    const handleCancel = () => {
        setIsModalVisible(false);
    };
    return (
        <>
            <Row style={{textAlign:'center',marginBottom : '1px'}}>
                <Col span={24}>
                    <Button onClick={showModal}>+</Button>
                </Col>
            </Row>
            <Modal title="Basic Modal" visible={isModalVisible} onOk={handleOk} onCancel={handleCancel}>
                <Input placeholder="Name" ref={Name}/>
                <Input placeholder="Memebers" ref={Memebers}/>
            </Modal>
            <Row style={{ height : '65vh', overflow:'scroll'}}>
                <Col span={24} style={{textAlign:'center'}}>
                    <Groupchat/>
                </Col>
            </Row>
           
        </>
    )
}
export default Chatgroup;