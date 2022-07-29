import axios from "axios";

const CallCreateRoom = async (id) => {
    try{
        const response = await axios.get(`http://localhost:8080/api/room/${id}`)
        console.log("response create room",response)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const CallCreateGroup = async (name,admin,email) =>{
    console.log(name)
    const ls = email.split(",")
    let listMemebers = []
    for(let i = 0 ; i < ls.length ; i++){
        const email = {"email" : ls[i]}
        listMemebers.push(email)
    }
    const data = JSON.stringify({"name" : name,"admin":admin,"members" : listMemebers })
    // console.log(data)
    try{
        const response = await axios.post("http://localhost:8080/api/room",data,{headers: { 'Content-Type': 'application/json' }})
        console.log("response create group",response)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const GetGroup = async (email) => {
    try{
        const response = await axios.get(`http://localhost:8080/api/room/group/${email}`)
        // console.log("response get group",response)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const GetMembers = async (id) =>{
    try{
        const response = await axios.get(`http://localhost:8080/api/room/group/members/${id}`)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const CallRemoveMember = async (room_id,email) =>{
    try{
        const response = await axios.delete(`http://localhost:8080/api/room/group/members/${room_id}/${email}`)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const AddMember = async (room_id,email) => {
    try{
        const response = await axios.post(`http://localhost:8080/api/room/group/members/${room_id}/${email}`)
        console.log(response)
        return response.data
    }catch(error){
        console.log(error)
    }
}
export { CallCreateRoom,CallCreateGroup,GetGroup,GetMembers,CallRemoveMember,AddMember}