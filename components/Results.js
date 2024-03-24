import React, { useRef } from 'react'
import { View, Text, ScrollView, TouchableOpacity, SafeAreaView, ActivityIndicator, Image } from 'react-native'
import ListItem from './ListItem'
import { MaterialIcons } from '@expo/vector-icons';

export default function Results(props) {
  // For scrolling to top
  const resultsScroll = useRef()
  const scrollUp = () => {
    resultsScroll.current.scrollTo({ y: 0, animated: true })
  }

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
        props.products = props.products.filter(product => {
          return props.filterVendor.includes(product.vendor)
        })
      }
    } else {
      return <View className="h-full items-center m-5">
        <Image source={require('../assets/no_results.png')} resizeMode='contain' className="w-5/6" />
      </View>

    }

    // Sort
    if (props.sortBy === 'price') {
      props.products?.sort((a, b) => { return parseFloat(a.price) - parseFloat(b.price) })
    } else {
      props.products?.sort((a, b) => { return a.vendor.localeCompare(b.vendor) })
    }

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
      <View className={props.filtersOpen ? "h-screen bg-slate-200 opacity-50 absolute w-full z-10" : "hidden"}></View>
      <TouchableOpacity
        onPress={scrollUp}
        className="bg-slate-200 border border-slate-500 opacity-50 rounded-full w-14 h-14 items-center justify-center absolute bottom-12 right-3 z-40"
        style={props.loading || props.error ? { display: 'none' } : {}}
      >
        <MaterialIcons name="arrow-upward" size={24} color="black" />
      </TouchableOpacity>
      <ScrollView ref={resultsScroll} className="mb-5 w-full" scrollEnabled={props.filtersOpen ? props.filtersOpen : true}>
        <View className="flex-row flex-wrap p-3 pt-0 pb-0 items-center justify-center">
          {productsList()}
        </View>
        <View className="justify-center items-center m-5">
          <Text
            className="text-center text-slate-500"
            style={props.loading || props.error ? { display: 'none' } : {}}
          >
            End of results.
          </Text>
        </View>
      </ScrollView>
    </SafeAreaView>
  )
}
