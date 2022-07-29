import axios from "axios";

const CallCreateRoom = async (id) => {
    try{
        const response = await axios.get(`http://localhost:8080/api/room/${id}`)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const CallCreateGroup = async (name,email) =>{
    console.log(name)
    const ls = email.split(",")
    let listMemebers = []
    for(let i = 0 ; i < ls.length ; i++){
        const email = {"email" : ls[i]}
        listMemebers.push(email)
    }
    const data = JSON.stringify({"name" : name,"members" : listMemebers })
    // console.log(data)
    try{
        const response = await axios.post("http://localhost:8080/api/room",data,{headers: { 'Content-Type': 'application/json' }})
        // console.log(response.data)
        return response.data
    }catch(error){
        console.log(error)
    }
}
const GetGroup = async (email) => {
    try{
        const response = await axios.get(`http://localhost:8080/api/room/group/${email}`)
        return response.data
    }catch(error){
        console.log(error)
    }
}
export { CallCreateRoom,CallCreateGroup,GetGroup }