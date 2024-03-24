import React, { useState, useRef, useContext, useEffect } from "react";
import {
    Pressable,
    Text,
    Modal,
    View,
    SafeAreaView,
    Keyboard,
    StatusBar
} from 'react-native'
import { MaterialCommunityIcons } from '@expo/vector-icons';
import DropDownPicker from 'react-native-dropdown-picker';
import Filters from "../components/Filters";
import Favourites from "../components/Favourites";
import { AuthContext } from "../AuthContextProvider";
import { getAPIUrl } from '../serverAPI'
import axios from 'axios'

const API_URL = getAPIUrl()

export default function FavouritesScreen({ route, navigation }) {

    const { userUID } = useContext(AuthContext);

    const [products, setProducts] = useState([])
    const [vendors, setVendors] = useState([])

    // For loading and error
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)
    const [noResults, setNoResults] = useState(false)

    // Drop Down Pickers
    const [openSort, setOpenSort] = useState(false);
    const [valueSort, setValueSort] = useState('');
    const [itemsSort, setItemsSort] = useState([
        { label: 'By Price', value: 'price' },
        { label: 'By Store', value: 'store' },
    ]);

    const [openView, setOpenView] = useState(false);
    const [valueView, setValueView] = useState('');
    const [itemsView, setItemsView] = useState([
        { label: 'Big View', value: 'big' },
        { label: 'Grid View', value: 'grid' },
        { label: 'List View', value: 'list' },
    ]);

    // For filters
    const [priceRange, setPriceRange] = useState([0, 0]);
    const [valuePriceFilter, setValuePriceFilter] = useState([0, 0]);
    const [modalVisible, setModalVisible] = useState(false);
    const [valueWordFilter, setValueWordFilter] = useState('');
    const wordFilterRef = useRef('')
    const lastKeyword = useRef(wordFilterRef.current)
    const [valueVendorFilter, setValueVendorFilter] = useState([]);


    // this function will extract the price range from the products array
    const getPriceRange = (productsList) => {
        let prices = productsList.map(product => product.price)
        let min = Math.min(...prices)
        let max = Math.max(...prices)
        setPriceRange([min, max])
        setValuePriceFilter([min, max])
    }

    async function getFavourites() {
        try {
            await axios.get(API_URL + "user/" + userUID + "/favorites").then(response => {
                if (!response.data.products || response.data.products.length === 0) {
                    setLoading(false)
                    setNoResults(true)
                    return
                }
                let vendorsDict = response.data.vendors.reduce((obj, item) => {
                    obj[item.name] = item.logo_url
                    return obj
                }, {})
                setVendors(vendorsDict)
                getPriceRange(response.data.products)
                setProducts(response.data.products)
                setLoading(false)
                setError(false)
            }).catch(error => {
                console.log(error)
                setLoading(false)
                setError(true)
            })
        } catch (error) {
            console.log(error)
            setLoading(false)
            setError(true)
        }
    }

    useEffect(() => {
        getFavourites()
    }, [])



    return (
        <SafeAreaView
            className="h-screen"
            style={{ paddingTop: StatusBar.currentHeight + 5, paddingBottom: StatusBar.currentHeight * 5 }}
        >
            { /* Header */}
            <View className="flex-row items-center justify-center pt-8 pb-3">
                <MaterialCommunityIcons name="account-heart-outline" size={60} color="#666666" />
                <Text className="text-3xl font-bold ml-5">Favorites</Text>
            </View>
            <Text className="w-full h-px bg-gray-200 rounded-lg"></Text>

            { /* Drop Down Pickers */}
            <View className="flex-row items-center p-2 z-40 w-1/3 devide-x-0 divide-slate-500 divide-dashed ml-4">
                <DropDownPicker
                    className="border-0 rounded-sm bg-transparent justify-start items-center"
                    open={openView}
                    value={valueView}
                    items={itemsView}
                    setOpen={setOpenView}
                    setValue={setValueView}
                    setItems={setItemsView}
                    placeholder={"View type"}
                    disableBorderRadius={true}
                    style={loading || error || noResults ? { display: 'none' } : {}}
                    onPress={Keyboard.dismiss()}
                />
                <DropDownPicker
                    className="border-0 rounded-sm bg-transparent justify-start items-center"
                    open={openSort}
                    value={valueSort}
                    items={itemsSort}
                    setOpen={setOpenSort}
                    setValue={setValueSort}
                    setItems={setItemsSort}
                    placeholder={"Sort by"}
                    disableBorderRadius={true}
                    style={loading || error || noResults ? { display: 'none' } : {}}
                    onPress={Keyboard.dismiss()}
                />

                <View className={loading || error || noResults ? "hidden" : "justify-start"}>
                    <Modal
                        animationType="slide"
                        transparent={true}
                        visible={modalVisible}
                        onRequestClose={() => setModalVisible(false)}>
                        <View className="flex-1 justify-center items-center mt-2 z-30">
                            <View className="m-16 bg-[#FFFFFF] rounded-2xl p-2 pb-7 items-center shadow-xl shadow-black">
                                <Filters
                                    priceValues={valuePriceFilter}
                                    setPriceValues={setValuePriceFilter}
                                    minPrice={priceRange[0]}
                                    maxPrice={priceRange[1]}

                                    wordFilterRef={wordFilterRef}
                                    lastKeyword={lastKeyword}
                                    modalVisible={modalVisible}
                                    setModalVisible={setModalVisible}
                                    valueWordFilter={valueWordFilter}
                                    setValueWordFilter={setValueWordFilter}

                                    vendors={vendors}
                                    valueVendorFilter={valueVendorFilter}
                                    setValueVendorFilter={setValueVendorFilter}
                                />
                            </View>
                        </View>
                    </Modal>
                    <Pressable
                        onPress={() => {
                            setModalVisible(true)
                        }} >
                        <Text className="pl-4">Show Filters</Text>
                    </Pressable>
                </View>
            </View>

            {/* Favourites */}
            <Favourites
                view={valueView}
                sortBy={valueSort}
                filterPrice={valuePriceFilter}
                filterWord={valueWordFilter}
                filterVendor={valueVendorFilter}
                loading={loading}
                setLoading={setLoading}
                error={error}
                setError={setError}
                filtersOpen={modalVisible}
                getPriceRange={getPriceRange}
                noResults={noResults}
                setNoResults={setNoResults}
                products={products}
                vendors={vendors}
            />
        </SafeAreaView>
    );
}
