import React, { useEffect, useState } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';
import { StatusBar } from 'expo-status-bar';
import HomeScreen from './screens/HomeScreen';
import Loading from './screens/Loading';
import WelcomeScreen from './screens/WelcomeScreen';
import ResultsScreen from './screens/ResultsScreen';
import AuthContextProvider from './AuthContextProvider';
import 'expo-dev-client';
import { authUser } from './serverAPI';
import FavouritesScreen from './screens/FavouritesScreen';


const Stack = createNativeStackNavigator();

export default function App() {

  const [user, setUser] = useState(undefined);
  const [loading, setLoading] = useState(true);
  const [userUID, setUserUID] = useState(undefined);
  const [favourites, setFavourites] = useState([])

  useEffect(() => {
    if (userUID) {
      authUser(user.token, userUID);
    }
  }, [userUID]);

  return (
    <AuthContextProvider value={{
                                  user, setUser,
                                  loading, setLoading,
                                  userUID, setUserUID,
                                  favourites, setFavourites
                                }} >
      <NavigationContainer>
        <Stack.Navigator screenOptions={{ headerShown: false }} >
          {
            loading ? (
              <Stack.Screen name="Loading" component={Loading} />
            ) :
              user ? (
                <>
                  <Stack.Screen name="Home" component={HomeScreen} />
                  <Stack.Screen name="Results" component={ResultsScreen} />
                  <Stack.Screen name="Favourites" component={FavouritesScreen} />
                </>
              ) : (
                <>
                  <Stack.Screen name="Welcome" component={WelcomeScreen} />
                </>
              )
          }
        </Stack.Navigator>
        <StatusBar style='dark' />
      </NavigationContainer>
    </AuthContextProvider>
  );
}
