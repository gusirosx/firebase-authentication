# Firebase Authentication Backend REST API
This is a backend REST API that allows users to authenticate using Firebase Authentication. It is built using Golang and the Firebase Admin SDK, which provides functionality for use with backend code.

With Firebase Authentication, users are intended to sign in on the frontend and pass an ID token to the backend. The Firebase Admin SDK can be used to verify the ID token and perform additional authentication-related tasks on the backend.

Please note that this API is intended to be used with backend code, and is not designed for use with frontend applications written in Golang. If you need to implement Firebase Authentication in a frontend application, you can call the Firebase Auth REST APIs directly.

For more information on Firebase Authentication, please refer to the official [documentation](https://firebase.google.com/docs/auth?hl=pt-br).