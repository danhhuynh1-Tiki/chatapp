import {createStore} from "redux"

let sizeUserChat = { size : 2}
let sizeGroupChat = {size : 2}
let sizeMessageChat = {size: 10}

const changeSize = (state,action) =>{
    state.size = state.size * 2
}

const UpdateSizeUserChat = createStore(changeSize,sizeUserChat)
const UpdateSizeGroupChat = createStore(changeSize,sizeGroupChat)
const updateSizeMessageChat = createStore(changeSize,sizeMessageChat)

export { UpdateSizeUserChat,UpdateSizeGroupChat,updateSizeMessageChat}