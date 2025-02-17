package fi.isdc_helsinki.splitbit

import android.content.SharedPreferences
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import fi.isdc_helsinki.splitbit.client.apis.DefaultApi
import fi.isdc_helsinki.splitbit.client.models.UserCredentials
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import androidx.security.crypto.MasterKey
import com.auth0.android.jwt.JWT


class LoginViewModel : ViewModel() {
    var basePath = ""
    // Make the API call in a coroutine
    fun login(request: UserCredentials) {
        viewModelScope.launch(Dispatchers.IO) {
            try {
                val api = DefaultApi("https://split-isdc.kuchta.dev")
                val response = api.loginPost(request)

                val jwt = JWT(response.token)
                val username = jwt.getClaim("username").asString() // Extract username
                val userID = jwt.getClaim("id").asInt()     // Extract userID

                val editor = g.encryptedSharedPreferences.edit()
                println(response)
                println(username)
                println(userID)
                //println(request.username)

                editor.putString("token", response.token) // should be working
                editor.putString("username", username)
                editor.putString("basePath", basePath)
                if (userID != null) {
                    editor.putInt("userID", userID)
                }
                editor.apply()

                val u = g.encryptedSharedPreferences.getString("username", null)
                val id = g.encryptedSharedPreferences.getInt("userID", -1)
                val t = g.encryptedSharedPreferences.getString("token", "")

                println("Username: $u, ID: $id, Is Logged In: $t")
//                println("kurwa")
//                val token = response.token
            } catch (e: Exception) {
                // Handle error here
                println(e)
            }
        }
    }

    fun checkServer(server: String): Int {
        viewModelScope.launch(Dispatchers.IO) {
            try {
                val response = DefaultApi(server).getPing()
                basePath = server
            } catch (e: Exception) {
                // Handle error here
                println(e)
            }
        }

        if (basePath == server) {
            return 1
        }
        else {
            return -1
        }
    }
}
