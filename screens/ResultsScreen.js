import React, { useState, useRef, useEffect, useContext } from 'react'
import {
    Pressable,
    Text,
    Modal,
    View,
    SafeAreaView,
    TextInput,
    Keyboard,
    StatusBar
} from 'react-native'
import { AntDesign } from '@expo/vector-icons';
import Results from '../components/Results';
import DropDownPicker from 'react-native-dropdown-picker';
import { getAPIUrl } from '../serverAPI'
import axios from 'axios'
import { AuthContext } from '../AuthContextProvider'
import Filters from '../components/Filters';

const API_URL = getAPIUrl()

export default function ResultsScreen({ route, navigation }) {

    const { favourites } = useContext(AuthContext);

    // For loading and error
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

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

    // For getting results
    const searchRef = useRef(null)
    const [products, setProducts] = useState([])
    const [vendors, setVendors] = useState([])
    const [searchProduct, setSearchProduct] = useState(route.params.product)

    async function getResults() {
        try {
            let query = !searchRef.current || !searchRef.current.value || searchRef.current.value === '' ? searchProduct : searchRef.current.value
            await axios.post(API_URL + "search", {
                query: query
            }).then(response => {

                let vendorsDict = response.data.vendors.reduce((obj, item) => {
                    obj[item.name] = item.logo_url
                    return obj
                }, {})
                setProducts(response.data.products)
                getPriceRange(response.data.products)
                setVendors(vendorsDict)
                setError(false)
                setLoading(false)
                searchRef.current.value = ''
            }).catch(error => {
                setError(true)
                console.log(error)
            })
        } catch (error) {
            setError(true)
            console.log(error)
        }
    }

    // this function will extract the price range from the products array
    const getPriceRange = (productsList) => {
        let prices = productsList.map(product => product.price)
        let min = Math.min(...prices)
        let max = Math.max(...prices)
        setPriceRange([min, max])
        setValuePriceFilter([min, max])
    }

    useEffect(() => {
        getResults()
    }, [searchProduct])

    const searchSubmit = () => {
        Keyboard.dismiss()
        if (!searchRef.current.value || searchRef.current.value === '') return
        setSearchProduct(searchRef.current.value)
        setLoading(true)
        // getResults()
    }

    return (
        <SafeAreaView
            className="h-screen"
            style={{ paddingTop: StatusBar.currentHeight + 5, paddingBottom: StatusBar.currentHeight * 4 }}
        >

            { /* Search Bar */}
            <View className="flex-row justify-center p-2">
                <View className="flex-row bg-gray-200 rounded-lg ml-3 mr-3 p-3 pl-5 pr-5 flex-1">
                    <AntDesign name="search1" size={20} color="gray" />
                    <TextInput
                        ref={searchRef}
                        placeholder={searchProduct}
                        className="ml-2 font-l w-5/6"
                        editable={!loading}
                        caretHidden={true}
                        onChangeText={(value) => searchRef.current.value = value}
                        onSubmitEditing={searchSubmit}
                    />
                </View>
            </View>

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
                    style={loading || error ? { display: 'none' } : {}}
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
                    style={loading || error ? { display: 'none' } : {}}
                    onPress={Keyboard.dismiss()}
                />

                <View className={loading || error ? "hidden" : "justify-start"}>
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

            { /* Results */}
            <Results
                view={valueView}
                sortBy={valueSort}
                filterPrice={valuePriceFilter}
                filterWord={valueWordFilter}
                filterVendor={valueVendorFilter}
                loading={loading}
                setLoading={setLoading}
                error={error}
                setError={setError}
                products={products}
                setProducts={setProducts}
                vendors={vendors}
                setVendors={setVendors}
                getResults={getResults}
                searchProduct={searchProduct}
                filtersOpen={modalVisible}
                favourites={favourites}
            />

        </SafeAreaView>
    )
}

ResultsScreen.defaultProps = {
    product: 'Search...'
}
