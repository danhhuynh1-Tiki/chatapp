import {createStore} from "redux"

let sizeUserChat = { size : 2}
let sizeGroupChat = {size : 2}
let sizeMessageChat = {size: 10}

const changeSize = (state,action) =>{
    console.log("state",state.size)
    console.log("action",action.size)
    state.size = state.size + action.size
    return state
}

const UpdateSizeUserChat = createStore(changeSize,sizeUserChat)
const UpdateSizeGroupChat = createStore(changeSize,sizeGroupChat)
const UpdateSizeMessageChat = createStore(changeSize,sizeMessageChat)

export { UpdateSizeUserChat,UpdateSizeGroupChat,UpdateSizeMessageChat}