import React, { useEffect, useContext } from 'react'
import { SafeAreaView, ActivityIndicator } from 'react-native'
import Logo from '../components/Logo'
import { AuthContext } from '../AuthContextProvider'

export default function Loading() {

    const { setLoading, setUserUID } = useContext(AuthContext)

    useEffect(() => {
        setUserUID(undefined)
        setTimeout(() => {
            setLoading(false)
        }, 3000)
    })

    return (
        <SafeAreaView className="h-screen justify-center items-center">
            <Logo />
            <ActivityIndicator size="large" color='#999999' />
        </SafeAreaView>
    )
}

