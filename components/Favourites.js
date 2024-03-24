import React from 'react'
import { View, Text, ScrollView, SafeAreaView, ActivityIndicator, Image } from 'react-native'
import ListItem from './ListItem'

export default function Favourites(props) {

    const productsList = () => {

        if (props.loading)
            return (
                <View className="h-full items-center m-5">
                    <ActivityIndicator size="large" color='#999999' />
                </View>
            )

        if (props.noResults || props.products.length === 0)
            return (
                <Image
                    source={require('../assets/no_results.png')}
                    resizeMode='contain'
                    className="w-5/6"
                />
            )

        if (props.error)
            return (
                <View className="h-full items-center m-5">
                    <Image
                        source={require('../assets/error.png')}
                        resizeMode='contain'
                        className="w-5/6" />
                </View>
            )

        // Filter
        if (props.products) {
            // Price filter
            props.products = props.products.filter(product => {
                return product.price >= props.filterPrice[0] && product.price <= props.filterPrice[1]
            })

            // Word filter
            if (props.filterWord && props.filterWord !== '') {
                props.products = props.products.filter(product => {
                    return product.title.toLowerCase().includes(props.filterWord.toLowerCase())
                })
            }

            //Vendor filter
            if (props.filterVendor && props.filterVendor.length !== 0) {
                // keep pruducts that have a vendor in the filter list
                props.products = props.products.filter(product => {
                    return props.filterVendor.includes(product.vendor)
                })
            }
        }

        // Sort
        if (props.sortBy === 'price') {
            props.products?.sort((a, b) => { return parseFloat(a.price) - parseFloat(b.price) })
        } else {
            props.products?.sort((a, b) => { return a.vendor.localeCompare(b.vendor) })
        }

        return (
            props.products.map((product, index) => {
                return (
                    <ListItem
                        key={index}
                        view={props.view}
                        title={product.title}
                        price={product.price}
                        image={product.img}
                        link={product.link}
                        vendor={product.vendor}
                        vendorImage={props.vendors[product.vendor]}
                    />
                )
            }))
    }

    return (
        <SafeAreaView className="relative h-full">
            <View className={props.filtersOpen ? "h-screen bg-slate-200 opacity-50 absolute w-full z-10" : "hidden"}></View>
            <ScrollView className="mb-5 w-full" scrollEnabled={props.filtersOpen ? props.filtersOpen : true}>
                <View className="flex-row flex-wrap p-3 pt-0 pb-0 items-center justify-center">
                    {productsList()}
                </View>
                <View className="justify-center items-center m-5">
                    <Text
                        className={props.loading ? "hidden" : "text-center text-slate-500"}
                        style={props.loading || props.error || props.noResults ? { display: 'none' } : {}}
                    >
                        End of results.
                    </Text>
                </View>
            </ScrollView>
        </SafeAreaView>
    )
}