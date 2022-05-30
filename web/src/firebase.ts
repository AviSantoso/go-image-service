import { getAnalytics } from "firebase/analytics";
import { initializeApp } from "firebase/app";
import { getAuth, GoogleAuthProvider, signInWithPopup } from "firebase/auth";
import { createContext } from "solid-js";

export const FirebaseContext = createContext();

const firebaseConfig = {
  apiKey: "AIzaSyDVqrgaox7bYmd54LRi-Kas1Wh4FA0-i5E",
  authDomain: "go-image-service-eb058.firebaseapp.com",
  projectId: "go-image-service-eb058",
  storageBucket: "go-image-service-eb058.appspot.com",
  messagingSenderId: "468079863379",
  appId: "1:468079863379:web:23d1df9da0f30a0408c8a6",
  measurementId: "G-SP6G6PXTXY",
};

export const app = initializeApp(firebaseConfig);
export const analytics = getAnalytics(app);
export const auth = getAuth(app);

export const googleAuthProvider = new GoogleAuthProvider();

export const signInWithGoogle = () => signInWithPopup(auth, googleAuthProvider);
export const signOut = () => auth.signOut();
