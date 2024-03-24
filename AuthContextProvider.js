import React from 'react';

export const AuthContext = React.createContext();

export default function AuthContextProvider(props) {

    return (
        <AuthContext.Provider value={props.value}>
            {props.children}
        </AuthContext.Provider>
    )
}