import React, { useContext, useEffect, useState } from "react";
import { View, Image, TouchableOpacity, Text, Dimensions } from "react-native";
import { AuthContext } from '../AuthContextProvider'
import { GoogleSignin } from '@react-native-google-signin/google-signin';
import Logo from '../components/Logo'
import { AntDesign } from '@expo/vector-icons';

const DEFAULT_IMG = 'https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_1280.png'
const IMG_SIZE = 45;
const SCREEN_WIDTH = Dimensions.get('window').width;

export default function UserLogo(props) {
  const { user, setUser, setUserUID } = useContext(AuthContext);
  const [img, setImg] = useState(DEFAULT_IMG)

  // For profile menu
  const [menuVisible, setMenuVisible] = useState(false)

  useEffect(() => {
    try {
      if (user && user.user) {
        if (!user.user.photo) {
          setImg(user.user.photoURL)
        } else {
          setImg(user.user.photo)
        }
      } else {
        setImg(user.photoURL)
      }
    }
    catch (error) {
      setImg(DEFAULT_IMG)
    }
  }, [user])

  const handleSignOut = async () => {
    setUser(undefined)
    setUserUID(undefined)
    try {
      await GoogleSignin.signOut()
    } catch (error) {
      console.error(error);
    }
  }

  const handleFavourites = () => {
    setMenuVisible(false)
    props.navigation.navigate('Favourites')
  }

  return (
    <View className="w-full p-5 pb-0 items-end">
      <TouchableOpacity
        className="rounded-full shadow-lg shadow-slate-900"
        activeOpacity={0.7}
        onPress={() => setMenuVisible(!menuVisible)}
      >
        <Image
          height={IMG_SIZE}
          width={IMG_SIZE}
          className="rounded-full"
          source={{ uri: img }}
        />
      </TouchableOpacity>
      {menuVisible && (
        <View
          className="absolute z-50 rounded-lg shadow-lg shadow-slate-700"
          style={{
            top: IMG_SIZE * 1.5,
            right: IMG_SIZE * 0.6,
            width: SCREEN_WIDTH * 0.31,
          }}
        >
          <View className="flex-1 bg-white p-1 rounded-md z-10 shadow-xl shadow-slate-700">
            <TouchableOpacity className="w-full p-1 flex-row items-center" onPress={handleFavourites}>
              <AntDesign name="hearto" size={16} color="black" />
              <Text className="pl-0.5 text-base font-medium pl-2">Favorites</Text>
            </TouchableOpacity>
            <Text className="w-full h-px bg-gray-200 rounded-lg"></Text>
            <TouchableOpacity className="w-full p-1 flex-row items-center" onPress={handleSignOut}>
              <AntDesign name="back" size={18} color="black" />
              <Text className="text-base font-medium pl-2">Sign Out</Text>
            </TouchableOpacity>
          </View>
        </View>
      )}
      <View className="w-full items-center pt-3">
        <Logo />
      </View>

    </View>
  );
}   