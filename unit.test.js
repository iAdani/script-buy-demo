//in this file we will test our app , using jest and according to the test cases in jira
import axios from 'axios';
import { authUser } from './serverAPI'

/** @type {import('jest').Config} */
const config = {
    verbose: true,
};

module.exports = config;
const API_URL = "http://127.0.0.1:8080/";
let testProps = {
    validUID: "KNOU2lEdOzYoR2Jm41MKvhn8CfZ2",
    invalidUID: "KNOU2lEdOzY0R2Jm41MKvhn8CfZ2",
    expiredToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk5YmNiMDY5MzQwYTNmMTQ3NDYyNzk0ZGZlZmE3NWU3OTk2MTM2MzQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiI5MzI4OTA4OTQwMTMtcm41OTQ0YmN2aTQyam1uYWkwNnBjcGQyZzNyN2t1MjIuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI5MzI4OTA4OTQwMTMtajZtZzAxOXZrdTNlazNidGFmbDVpc212aDhmbWpsM28uYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDU3NTI1OTIyMjU4Mzg4OTc3MDYiLCJlbWFpbCI6ImZiYWNoYXRhcHBAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsIm5hbWUiOiJjaGF0IGFwcCIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS9BQWNIVHRjSE5fTXZGYXVCMHM4TnJTVXUtZnFHVGRoNDZnMWk3bk03cmIxRj1zOTYtYyIsImdpdmVuX25hbWUiOiJjaGF0IiwiZmFtaWx5X25hbWUiOiJhcHAiLCJsb2NhbGUiOiJoZSIsImlhdCI6MTY4Nzc5OTIyNCwiZXhwIjoxNjg3ODAyODI0fQ.SiNYZO5qQhqWrB-cTOrJoVHPd51DWQIU20k4n0XWl7_gUjFZMSE6XxDIwpI0JNpXJsGJVDa0JTdeZthHizhIoLY07WPlbQrfAh-y0BGesY7F-IwE4sz5tuJbNCCBLS5C2C_YVuqJVWQRPJvqo0Y9qVyMDIzwU6oWkIfiFqHFtfIYJ8PyKFDOFjcgGRmZpl9jOeinDmLfgdkGWohm7y9HcZ0taF-_2BmSzUZZeO2RNPREBZS1eOmmmWiL__A2mVd44rDFz8wuAHsX8bcR1XHDhC3ZDyP9TVi8-dkzEyZsNo9KOb7eVBUVEf_gsAw_Ssg6fE9V5_YSmsoi7HwQrfUS6Q",
    validFavorite: {
        "product": {
            "title": "test product",
            "vendor": "test vendor",
            "price": "999",
            "img": "test.jpg",
            "link": "testLink.com"
        },
        "vendor": {
            "name": "test vendor",
            "logo_url": "testLogo.gif"
        }
    },
    validCategory: "Electronics",
    invalidCategory: "Bottles",
    validQuery: { query: "mx master 3s" },
    invalidQuery: "",
}
let testUser = {
    "uid": "TEST_UID_STRING",
    "name": "new user",
    "email": "new_test_email@gmail.com",
    "searches": [],
    "favorites": [],
    "favorites_vendors": []
}

// Epic 1 : User Management & Logic Tests
describe('User Management & Logic Tests', () => {

    // Story 1 : User - Log In
    // As a user, I want to be able to log into my account so that I can manage my product favorites
    test('User - Log In    <valid uid>', async () => {
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + "token"
        await axios.post(API_URL + 'auth', {
            uid: testProps.validUID,
        }).then(response => {
            expect(response.status).toBe(200)
        })
    })
    test('User - Log In    <invalid uid>', async () => {
        axios.defaults.headers.common['Authorization'] = 'Bearer ' + testProps.expiredToken
        await axios.post(API_URL + 'auth', {
            uid: testProps.invalidUID,
        }).catch(error => {
            expect(error.response.status).toBe(400)
        })
    })
    test('User - Get User  <valid uid>', async () => {
        await axios.get(API_URL + 'user/' + testProps.validUID).then(response => {
            expect(response.status).toBe(200)
            expect(response.data.name).toBe("chat app")
            expect(response.data.email).toBe("fbachatapp@gmail.com")
        })
    })
    test('User - Get User  <invalid uid>', async () => {
        await axios.get(API_URL + 'user/' + testProps.invalidUID).catch(error => {
            expect(error.response.status).toBe(404)
        })
    })

    // Story 2 : User - Create Account
    // As a user, I want to be able to create an account so that I can track prices and save my search history
    test('User - Create User', async () => {
        await axios.post(API_URL + 'user', testUser).then(response => {
            expect(response.status).toBe(201)
            axios.get(API_URL + 'user/' + testUser.uid).then(response => {
                expect(response.status).toBe(200)
                expect(response.data).toMatchObject(testUser)
            })
        })
    })
    test('User - Remove User', async () => {
        await axios.delete(API_URL + 'user/' + testUser.uid).then(response => {
            expect(response.status).toBe(200)
            axios.get(API_URL + 'user/' + testUser.uid).catch(error => {
                expect(error.response.status).toBe(404)
            })
        })
    })
})

// Epic 2 : Favorites & Ternding Logic Tests
describe('Favorites & Ternding Logic Tests', () => {
    test('Favorites - Get  <valid uid>', async () => {
        await axios.get(API_URL + 'user/' + testProps.validUID + '/favorites').then(response => {
            expect(response.status).toBe(200)
        })
    })
    test('Favorites - Get  <invalid uid>', async () => {
        await axios.get(API_URL + 'user/' + testProps.invalidUID + '/favorites').catch(error => {
            expect(error.response.status).toBe(404)
        })
    })
    test('Favorites - Add', async () => {
        await axios.post(API_URL + 'user/' + testProps.validUID + '/favorites', testProps.validFavorite).then(response => {
            expect(response.status).toBe(201)
            axios.get(API_URL + 'user/' + testProps.validUID + '/favorites').then(response => {
                expect(response.status).toBe(200)
                expect(response.data.products).toContainEqual(testProps.validFavorite.product)
                expect(response.data.vendors).toContainEqual(testProps.validFavorite.vendor)
            })
        })
    })
    test('Favorites - Remove', async () => {
        await axios.delete(API_URL + 'user/' + testProps.validUID + '/favorites', { data: testProps.validFavorite }).then(response => {
            expect(response.status).toBe(200)
            axios.get(API_URL + 'user/' + testProps.validUID + '/favorites').then(response => {
                expect(response.status).toBe(200)
                expect(response.data.products).not.toContainEqual(testProps.validFavorite.product)
                expect(response.data.vendors).not.toContainEqual(testProps.validFavorite.vendor)
            })
        })
    })
    test('Trending - Get  <valid category>', async () => {
        await axios.get(API_URL + 'trending/' + testProps.validCategory).then(response => {
            expect(response.status).toBe(200)
            expect(response.data.products.length).toBeGreaterThan(0)
        })
    })
    test('Trending - Get  <invalid category>', async () => {
        await axios.get(API_URL + 'trending/' + testProps.invalidCategory).catch(error => {
            expect(error.response.status).toBe(404)
        })
    })
})

// Epic 3 : Search Logic Tests
describe('Search Logic Tests', () => {
    test('Search Product  <invalid query>', async () => {
        await axios.post(API_URL + 'search/', testProps.invalidQuery).catch(error => {
            expect(error.response.status).toBe(400)
        })
    })
})