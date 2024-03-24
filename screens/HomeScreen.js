import React, { useEffect, useState, useContext } from 'react'
import { View, TextInput, SafeAreaView, StatusBar } from 'react-native'
import { AntDesign } from '@expo/vector-icons';
import UserLogo from '../components/UserLogo';
import Trending from '../components/Trending';
import { getAPIUrl } from '../serverAPI'
import axios from 'axios';
import { AuthContext } from '../AuthContextProvider';
import auth from '@react-native-firebase/auth';

export default function HomeScreen({ route, navigation }) {

    const { userUID, setUserUID, setFavourites } = useContext(AuthContext);

    // For loading and error
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    // For searching
    const [search, setSearch] = useState('')

    const searchSubmit = () => {
        if (!search || search === '') return
        navigation.navigate('Results', { product: search })
    }

    // For Popular Searches
    const API_URL = getAPIUrl()
    const [products, setProducts] = useState([])
    const [vendors, setVendors] = useState([])
    async function getTrending() {
        try {
            await axios.get(API_URL + "user/" + userUID + "/favorites").then(response => {
                setFavourites(response.data.products)
            }).catch(error => {
                console.log(error)
            })
            await axios.get(API_URL + "trending/" + "Electronics").then(response => {
                let vendorsDict = response.data.vendors.reduce((obj, item) => {
                    obj[item.name] = item.logo_url
                    return obj
                }, {})
                setProducts(response.data.products)
                setVendors(vendorsDict)
                setLoading(false)
                setError(false)
            }).catch(error => {
                setLoading(false)
                setError(true)
            })
        } catch (error) {
            setLoading(false)
            setError(true)
        }
    }

    useEffect(() => {
        if (userUID)
            getTrending()
        else
            setUserUID(auth().currentUser.uid)
    }, [userUID])

    return (
        <SafeAreaView
            className="h-screen"
            style={{ paddingTop: StatusBar.currentHeight + 5, paddingBottom: StatusBar.currentHeight * 8.5 }}
        >
            { /* User Image & ScriptBuy Logo */}
            <UserLogo navigation={navigation} />

            { /* Search Bar */}
            <View className="flex-row justify-center p-2">
                <View className="flex-row bg-gray-200 rounded-lg ml-3 mr-3 p-3 pl-5 pr-5 flex-1 items-center">
                    <AntDesign name="search1" size={20} color="gray" />
                    <TextInput
                        placeholder="Search..."
                        className="ml-2 font-l w-5/6"
                        onChangeText={setSearch}
                        onSubmitEditing={searchSubmit}
                    />
                </View>
            </View>

            { /* Trending */}
            <Trending
                products={products}
                vendors={vendors}
                loading={loading}
                error={error}
            />

        </SafeAreaView>
    )
}

HomeScreen.defaultProps = {
    postLogin: false
}