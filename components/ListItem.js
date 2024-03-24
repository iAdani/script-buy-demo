import React, { useState, useContext, useEffect } from 'react'
import { View, Text, Image, Linking, TouchableOpacity } from 'react-native'
import { AntDesign } from '@expo/vector-icons';
import { getAPIUrl } from '../serverAPI'
import { AuthContext } from '../AuthContextProvider'
import axios from 'axios'

const API_URL = getAPIUrl()
const FAV_ICON_SIZE = 21

export default function ListItem(props) {
  // For favourites
  const [isFavourite, setIsFavourite] = useState(false)
  const { userUID, favourites, setFavourites } = useContext(AuthContext);

  async function handleFavourite() {
    const product = {
      title: props.title,
      vendor: props.vendor,
      price: props.price,
      img: props.image,
      link: props.link,
    }
    const vendor = {
      name: props.vendor,
      logo_url: props.vendorImage,
    }
    if (isFavourite) {
      try {
        await axios.delete(API_URL + "user/" + userUID + "/favorites", { data: { product, vendor } }).then(response => {
          setIsFavourite(false)
          setFavourites(prevFavourites => prevFavourites.filter(favourite => (
            favourite.title !== product.title ||
            favourite.vendor !== product.vendor ||
            favourite.price !== product.price ||
            favourite.img !== product.img ||
            favourite.link !== product.link
          )));
        }).catch(error => {
          console.log(error)
        })
      } catch (error) {
        console.log(error)
      }
    } else {
      try {
        await axios.post(API_URL + "user/" + userUID + "/favorites", { product, vendor }).then(response => {
          setIsFavourite(true)
          setFavourites(prevFavourites => [...prevFavourites, product]);
        }).catch(error => {
          console.log(error)
        })
      } catch (error) {
        console.log(error)
      }
    }
  }

  useEffect(() => {
    const product = {
      title: props.title,
      vendor: props.vendor,
      price: props.price,
      img: props.image,
      link: props.link,
    }
    setIsFavourite(favourites?.some(favourite => JSON.stringify(favourite) === JSON.stringify(product)))
  }, [favourites])



  const [image, setImage] = useState(props.image) // For handling broken images
  let viewClass, imgSize, titleClass, vendorClass, vendorImageClass, textSize
  const baseViewClass = "border-0 items-center justify-between bg-white rounded-lg shadow-lg shadow-slate-700 "
  switch (props.view) {
    case 'big':
      viewClass = "w-full min-h-fit mb-4 mr-2 ml-2 p-2"
      vendorClass = "flex-row justify-between w-full pl-1 pt-2"
      vendorImageClass = "w-1/4"
      imgSize = 230
      textSize = 80
      titleClass = ""
      break;
    case 'list':
      viewClass = "flex-row-reverse w-full h-28 mb-4 mr-2 ml-2 p-2"
      vendorClass = "flex-col-reverse justify-between h-full pl-1 w-1/4 p-1"
      vendorImageClass = "h-1/2"
      imgSize = 80
      textSize = 70
      titleClass = " p-5"
      break;
    default:
      viewClass = "w-2/5 h-64 mb-4 mr-2 ml-2 pr-1 pl-1 pt-2"
      vendorClass = "flex-row justify-between w-full pl-1"
      vendorImageClass = "w-1/2"
      imgSize = 130
      textSize = 65
      titleClass = ""
  }

  const openLink = () => {
    Linking.openURL(props.link).catch((err) => console.error('Failed to open link:\n', err));
  }

  return (

    <TouchableOpacity onPress={openLink} activeOpacity={0.5} className={baseViewClass + viewClass}>
      <TouchableOpacity
        onPress={handleFavourite}
        className="absolute top-0 left-0 p-2 z-10 bg-gray-50 rounded-full"
        activeOpacity={0.6}
      >
        {isFavourite ?
          <AntDesign name="heart" size={FAV_ICON_SIZE} color="#cf1702" />
          :
          <AntDesign name="hearto" size={FAV_ICON_SIZE} color="black" />
        }
      </TouchableOpacity>
      <Image
        source={{ uri: image }}
        style={{ width: imgSize, height: imgSize }}
        className={props.view === "big" ? 'p-1 mb-5 rounded-md' : 'p-1 rounded-md'}
        resizeMode='contain'
        onError={(e) => { setImage("https://www.warnersstellian.com/Content/images/product_image_not_available.png") }}
      />
      <Text className={"shrink" + titleClass}>
        {props.title.length < 80 ? props.title : props.title.substring(0, textSize) + '...'}
      </Text>
      <View className={vendorClass}>
        <Text className={props.price.length < 7 ? "font-bold text-lg" : "font-bold text-md pb-1"}>
          {'₪' + props.price}
        </Text>
        <Image source={{ uri: props.vendorImage }} className={vendorImageClass} resizeMode='contain' />
      </View>
    </TouchableOpacity>
  )
}

ListItem.defaultProps = {
  title: 'Item title',
  price: '750',
  image: 'https://cdn.cashcow.co.il/images/e5e9a092-3d01-4f35-b0dc-2738347a0a71.png'
}