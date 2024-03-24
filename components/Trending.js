import React from "react";
import { View, Text, ScrollView, SafeAreaView, ActivityIndicator, Image } from 'react-native'
import ListItem from './ListItem'

export default function Trending(props) {


    const productsList = () => {
        if (props.error)
            return <Image
                source={require('../assets/error.png')}
                resizeMode='contain'
                className="w-5/6"
            />

        if (props.loading)
            return <View className="h-full items-center m-5">
                <ActivityIndicator size="large" color='#999999' />
            </View>

        return props.products.map((product, index) => {
            return (
                <ListItem
                    key={index}
                    title={product.title}
                    price={product.price}
                    image={product.img}
                    view={props.view}
                    link={product.link}
                    vendor={product.vendor}
                    vendorImage={props.vendors[product.vendor]}
                />
            )
        })
    }

    return (
        <SafeAreaView className="relative h-full">
            <Text className="font-bold text-xl m-3">Trending</Text>
            <ScrollView className="mb-5 w-full">
                <View className="flex-row flex-wrap p-3 pt-0 pb-0 items-center justify-center">
                    {productsList()}
                </View>
            </ScrollView>
        </SafeAreaView>
    )
}