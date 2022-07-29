
import axios from "axios"

const CallUsersApi = async () => {
    try{

        const response = await axios.get("http://localhost:8080/api/users?size=10")
        // console.log("response login user",response)
        return response.data
    }catch(error){
        console.log("loi ne",error)
    }
}

const LogoutUsersApi = async () => {
    try{
        const response = await axios.get("http://localhost:8080/api/auth/logout")
        return response.data
    }catch(error){
        console.log("loi ne",error)
    }
} 
const GetUserApi = async () => {
    try{
        const response = await axios.get("http://localhost:8080/api/users/me")
        console.log(response.data)
        return response.data
    }catch(error){
        console.log("loi ne",error)
    }
}

export { CallUsersApi, LogoutUsersApi , GetUserApi};