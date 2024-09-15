package fi.isdc_helsinki.splitbit

import Item
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import kotlinx.coroutines.launch
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import Server


class ExpenseAddViewModel : ViewModel() {
    // Make the API call in a coroutine
    fun addItem(newItem: Item) {
        viewModelScope.launch {
            try {
                val response = Server.api.postItem(newItem)
                if (response.isSuccessful) {
                    response.body()?.let { _ ->
                        // Update local state if necessary
                    }
                }
            } catch (e: Exception) {
                // Handle error here
            }
        }
    }
}
