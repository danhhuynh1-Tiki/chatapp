import { createStore } from "redux";

let user = {
    email : ""
}

const SetEmail = (state,action) => {
    state.email = action.email
    return state
}

const EmailUser = createStore(SetEmail,user)

export {EmailUser}