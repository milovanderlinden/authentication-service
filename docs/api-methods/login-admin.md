**Login Admin**
----
  Check email and password for login in as an admin and return a token if credentials are correct. User has to be in the admin list to be able to sign in as an admin. The role 'admin' is written into the token.

* **URL**

  /v1/login/admin

* **Method:**

  `POST`

*  **URL Params**
  None

* **Data Params**
  **Required:**
  * **Type:** application/json
    **Content:** `{ "email": "<user email>", "password": "<user password>"}`

* **Success Response:**

  * **Code:** 200
    **Content:** `{ "token": "<token string>" }`

* **Error Response:**

  * **Code:** 400 Bad request
    **Content:** `{ "error" : "<error message>" }`
    **Typical reason:** Data format (json body of the Post request) wrong, e.g. missing key for email or password.

  OR

  * **Code:** 401 Unauthorized
    **Content:** `{ "error" : "<error message>" }`
    **Typical reason:** Password is wrong or email doesn't belong to any registered admin.

  OR

  * **Code:** 500 Internal server error
    **Content:** `{ "error" : "<error message>" }`
    **Typical reason:** Something went wrong during the token generation. User's credentials are ok, but method failed generating a valid token, e.g. because signing key is not available.

* **Sample Call:**
  TODO: add sample call for go

  ```javascript
    $.ajax({
      url: "/users/1",
      dataType: "json",
      type : "GET",
      success : function(r) {
        console.log(r);
      }
    });
  ```

* **Notes:**
  None