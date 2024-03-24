import React, { useState } from 'react'
import MultiSlider from '@ptomasroos/react-native-multi-slider';
import { AntDesign } from '@expo/vector-icons';
import { TextInput, View, Text, ScrollView, TouchableOpacity } from 'react-native';
import Vendors from './Vendors';


export default function Filters(props) {

    const [textValue, setTextValue] = useState('')
    const [textFilterEdited, setTextFilterEdited] = useState(false)


    const handleSliderValuesChange = (values) => {
        props.setPriceValues(values);
    };

    return (
        <>
            {/* {Price Filter} */}
            <Text className="font-semibold pb-2 mt-5">Price Between</Text>
            <View className="pt-9">
                <MultiSlider
                    values={[props.priceValues[0], props.priceValues[1]]}
                    sliderLength={200}
                    onValuesChangeFinish={handleSliderValuesChange}
                    isMarkersSeparated
                    enableLabel
                    min={props.minPrice}
                    max={props.maxPrice}
                    step={5}
                />
            </View>

            {/* Title Filter */}
            <Text className="pb-2 mt-3 font-semibold">Title Contains</Text>
            <View className="flex-row bg-gray-200 rounded-lg ml-3 mr-3 p-2 pl-3 pr-2">
                <AntDesign name="search1" size={22} color="gray" />
                <TextInput
                    className="flex-row ml-2 font-l w-5/6"
                    editable={props.modalVisible}
                    value={!textValue && !textFilterEdited ? props.valueWordFilter : textValue}
                    placeholder='Search...'
                    caretHidden={true}
                    onChangeText={(value) => {
                        setTextFilterEdited(true)
                        setTextValue(value)
                    }}
                    onEndEditing={(value) => {
                        props.setValueWordFilter(value.nativeEvent.text)
                    }}
                />
            </View>

            {/* Vendors Filter */}
            <Text className="pb-2 mt-6 font-semibold">Chosen Brands</Text>
            <View className="h-14">
                <ScrollView horizontal
                    showsHorizontalScrollIndicator={false}
                    className="ml-3 flex-row">
                    <Vendors
                        vendors={props.vendors}
                        valueVendorFilter={props.valueVendorFilter}
                        setValueVendorFilter={props.setValueVendorFilter} />
                </ScrollView>
            </View>


            <TouchableOpacity
                className="bg-blue-200 rounded-full p-6 pt-2 pb-2 mt-8"
                activeOpacity={0.8}
                onPress={() => {
                    props.setModalVisible(false)
                }}>
                <Text className="font-semibold">Close</Text>
            </TouchableOpacity>
        </>
    );
}