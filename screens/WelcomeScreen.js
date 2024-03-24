import React, { useEffect, useContext } from 'react'
import { Text, SafeAreaView, TouchableOpacity, Image } from 'react-native'
import Logo from '../components/Logo'
import auth from '@react-native-firebase/auth';
import 'expo-dev-client'
import { GoogleSignin } from '@react-native-google-signin/google-signin';
import { AuthContext } from '../AuthContextProvider'

let isFetching = false

export default function WelcomeScreen({ route, navigation }) {

  const { setUser, setUserUID } = useContext(AuthContext)

  GoogleSignin.configure({
    webClientId: '932890894013-j6mg019vku3ek3btafl5ismvh8fmjl3o.apps.googleusercontent.com',
  });

  const onAuthStateChanged = async (authUser) => {
    googleUser = await GoogleSignin.getCurrentUser()
    if (authUser && !isFetching && googleUser) {
      isFetching = true
      let newUser = JSON.parse(JSON.stringify(authUser))
      token = await GoogleSignin.getTokens()
      newUser['token'] = token.idToken
      setUser(newUser)
      setUserUID(authUser.uid)
      isFetching = false
    }
  }

  useEffect(() => {
    const subscriber = auth().onAuthStateChanged(onAuthStateChanged);
    return subscriber; // unsubscribe on unmount
  }, []);

  useEffect(() => {
    setUserUID(undefined)
  }, [])

  onGoogleButtonPress = async () => {
    try {// Check if your device supports Google Play
      await GoogleSignin.hasPlayServices({ showPlayServicesUpdateDialog: true });
      // Get the users ID token
      const { idToken } = await GoogleSignin.signIn();

      // Create a Google credential with the token
      const googleCredential = auth.GoogleAuthProvider.credential(idToken);

      // Sign-in the user with the credential
      const user_sign_in = auth().signInWithCredential(googleCredential);
      user_sign_in.then((user) => {
        onAuthStateChanged(user)
      }).catch((error) => {
        console.log(error)
      })
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <SafeAreaView className='justify-center items-center h-3/4'>
      <Logo />

      { /* Google */}
      <TouchableOpacity
        activeOpacity={0.8}
        className="bg-white rounded p-2 pl-5 pr-5 mt-5 w-5/6 rounded-lg flex-row items-center shadow"
        onPress={onGoogleButtonPress}
      >
        <Image source={require('../assets/google-logo.png')} className="h-6 w-6" resizeMode='contain' />
        <Text className="text-black text-center font-bold text-xl ml-10">Sign in with Google</Text>
      </TouchableOpacity>
      <Text className="text-center text-gray-500 mt-5">Please log in to use the app.</Text>
    </SafeAreaView>
  )
}

