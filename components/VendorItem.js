import React from 'react'
import { View, Text, TouchableOpacity, Image } from 'react-native'

export default function VendorItem(props) {

  const Filtered = props.valueVendorFilter.includes(props.title)

  // if the vendor is in the filters list, remove it.
  // if the vendor is not in the filters list, add it.
  const handlePress = () => {
    if (Filtered) {
      props.setValueVendorFilter(props.valueVendorFilter.filter(vendor => vendor !== props.title))
    } else {
      props.setValueVendorFilter([...props.valueVendorFilter, props.title])
    }
  }

  return (
    <TouchableOpacity
      className={Filtered ? "flex-row rounded-lg bg-green-300 mr-2 w-36" : "flex-row rounded-lg bg-gray-200 mr-2 w-36"}
      underlayColor="blue"
      onPress={handlePress}
    >
      <View className="justify-between flex-1 items-center flex-row p-2 pr-4">
        <View className="rounded-lg overflow-hidden">
          <Image
            source={{ uri: props.image }}
            resizeMode='contain'
            className="h-12 w-14 bg-white"
            onError={(e) => { setImage("https://www.warnersstellian.com/Content/images/product_image_not_available.png") }}
            />
            </View>
        <Text className="font-bold text-l">
          {props.title}
        </Text>
      </View>
    </TouchableOpacity>
  )
}
