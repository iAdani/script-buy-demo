import { getAuth } from "firebase/auth";
import { initializeApp } from "firebase/app";
import { getFirestore } from "firebase/firestore";

const firebaseConfig = {
    apiKey: "",
    authDomain: "",
    projectId: "",
    storageBucket: "",
    messagingSenderId: "",
    appId: "",
    measurementId: ""
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
const db = getFirestore();


export { auth, db };

export function errorHandling(error) {
    switch (error.code) {
        case "auth/email-already-in-use":
            return "Email already in use."
        case "auth/invalid-email":
            return "Invalid email."
        case "auth/weak-password":
            return "Password is too weak."
        case "auth/wrong-password":
            return "Incorrect password."
        case "auth/user-not-found":
            return "Invalid username or password."
        case "auth/user-disabled":
            return "This account has been disabled."
        default:
            return "An error has occured.\nerror code: " + error.code
    }
}