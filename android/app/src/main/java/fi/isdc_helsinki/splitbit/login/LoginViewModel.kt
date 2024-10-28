package fi.isdc_helsinki.splitbit

import AuthRequest
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.launch
import Server


class LoginViewModel : ViewModel() {

    // Make the API call in a coroutine
    fun login(request: AuthRequest) {
        viewModelScope.launch {
            try {
//                val response = Server.api.login(request)
//                if (response.isSuccessful) {
//                    response.body()?.let { _ ->
//                        // Update local state if necessary
//                    }
//                }
            } catch (e: Exception) {
                // Handle error here
            }
        }
    }
}
